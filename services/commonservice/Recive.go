package commonservice

import (
	"ms/usersAPI/application/mdlwr"
	"ms/usersAPI/data/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Revice : recive user from datasource for given primary key
// ReciveUser key godoc
// @Summary returns user
// @Description recive user by primary key
// @ID get-user-by-pk
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserExported
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users/{id} [get]

func (s *commonService) Recive(c *gin.Context) {
	//geting primary key as path parameter
	id := s.parseID(c)
	if id == -1 {
		return
	}

	//Role 1-worker, 2-financial manager, 3-manager, 4-admin
	//Permission 1 - (R) read only, 2 - CR (create, read), 3 - 	CRU (create, read, update) 4 - CRUD - all actions
	allowed := mdlwr.RequiredPermission(c, worker, read)
	if !allowed {
		return
	}
	//creating new user
	user := new(models.User)
	//request to datasource
	_, apierror := s.datasource.Recive(user, id)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	borderViolated := mdlwr.CompanyBorder(c, user.CompanyID)
	if borderViolated {
		return
	}

	//return json
	c.JSON(http.StatusOK, user.Export())
}

//FindAll : returns JSON list of all objects that mathes given query parameter
// FindAll key godoc
// @Summary updates users
// @Description updates all users that matches given parameters
// @ID update-users
// @Accept    json
// @Produce  json
// @Success 200 {object} models.UserExportedList
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users [put]
func (s *commonService) FindAll(c *gin.Context) {
	allowed := mdlwr.RequiredPermission(c, worker, read)
	if !allowed {
		return
	}
	//if query parameters aren't set
	params := new(models.User)
	apierr := s.queryToStruct(c, params)
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}

	params.CompanyID = s.getCompanyID(c)

	usersList := new(models.UserList)
	apierror := s.datasource.FindAll(usersList, params)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.JSON(http.StatusOK, usersList.Export())
}
