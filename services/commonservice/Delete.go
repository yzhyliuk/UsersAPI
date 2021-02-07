package commonservice

import (
	"fmt"
	"ms/usersAPI/application/mdlwr"
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Delete : delete user from datasource for given primary key
// DeleteUser key godoc
// @Summary deletes user
// @Description deletes user by primary key
// @ID delete-user-by-pk
// @Param id path int true "User ID"
// @Success 200 {object} string
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users/{id} [Delete]
func (s *commonService) Delete(c *gin.Context) {
	allowed := mdlwr.RequiredPermission(c, manager, delete)
	if !allowed {
		return
	}
	id := s.parseID(c)
	if id == -1 {
		return
	}
	userToDelete := new(models.User)
	s.datasource.Recive(userToDelete, id)

	if userToDelete.CompanyID != s.getCompanyID(c) {
		apierr := errors.NewUnauthorizedError("Company border violated")
		c.JSON(apierr.Status, apierr)
		return
	}

	apierror := s.datasource.Delete(userToDelete, id)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("User with id %d succesfully deleted", id))
}

//DeleteWhere : deletes all rows that matches with given parameters
// DeleteUsersWhere key godoc
// @Summary deletes users
// @Description deletes all users that matches given parameters
// @ID delete-users
// @Accept    json
// @Success 200 {object} string
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users [Delete]
func (s *commonService) DeleteWhere(c *gin.Context) {
	allowed := mdlwr.RequiredPermission(c, manager, delete)
	if !allowed {
		return
	}
	//Parse params
	userParams := new(models.User)
	apierr := s.queryToStruct(c, userParams)
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}
	//Setting company borders
	userParams.CompanyID = s.getCompanyID(c)
	//request to datasource
	apierror := s.datasource.DeleteWhere(userParams, userParams)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	c.String(http.StatusOK, "users successfully deleted")
}
