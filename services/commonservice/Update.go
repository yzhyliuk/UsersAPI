package commonservice

import (
	"ms/usersAPI/application/mdlwr"
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Update : updates row with given primary key in datasource
// UpdateUser key godoc
// @Summary updates user
// @Description update user with given primary key
// @ID update-user
// @Accept    json
// @Param id path int true "User ID"
// @Success 200 {object} string
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users/{id} [put]
func (s *commonService) Update(c *gin.Context) {
	//Parsing Id
	id := s.parseID(c)
	if id == -1 {
		return
	}
	//Parsing Body
	user := new(models.User)
	err := s.parseBody(c, user)
	if err != nil {
		return
	}
	//Reciving user with current id
	currentUser := new(models.User)
	_, apierror := s.datasource.Recive(currentUser, id)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	//if user updating itself - don't ask permission
	if currentUser.ID != s.getUserID(c) {
		allowed := mdlwr.RequiredPermission(c, manager, update)
		if !allowed {
			return
		}
	} else {
		//Othervise don't give them rights to change the role
		user.Role = currentUser.Role
	}
	user.ID = id
	user.Merge(currentUser)

	if user.CompanyID != s.getCompanyID(c) {
		apierror = errors.NewUnauthorizedError("Company border is violated")
		c.AbortWithStatusJSON(apierror.Status, apierror)
		return
	}

	apierror = user.Validate()
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	apierror = s.datasource.Update(user)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.String(http.StatusOK, "User successfully updated")
}

//UpdateWhere : updates all rows that matches given parameters
// UpdateUsersWhere key godoc
// @Summary updates users
// @Description updates all users that matches given parameters
// @ID update-users
// @Accept    json
// @Success 200 {object} string
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users [put]
func (s *commonService) UpdateWhere(c *gin.Context) {
	allowed := mdlwr.RequiredPermission(c, manager, update)
	if !allowed {
		return
	}
	//Pasing body
	user := new(models.User)
	err := s.parseBody(c, user)
	if err != nil {
		return
	}
	//Validating input
	apierr := user.Validate()
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}
	//Parse params
	userParams := new(models.User)
	apierr = s.queryToStruct(c, userParams)
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}
	//Setting company border
	userParams.CompanyID = s.getCompanyID(c)

	//request to datasource
	apierror := s.datasource.UpdateWhere(user, userParams)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	c.String(http.StatusOK, "resources successfully updated")
}
