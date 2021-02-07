package commonservice

import (
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/argon"
	"ms/usersAPI/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create : creates new entry of user in datasource
// CreateUser key godoc
// @Summary creates user
// @Description create new user
// @ID create-new-user
// @Accept    json
// @Success 200 {object} string
// @Header 200 {string} string
// @Failure 404 {object} errors.APIError
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Failure default {object} errors.APIError
// @Router /users [Post]
func (s *commonService) Create(c *gin.Context) {

	//retrieving user from request body
	user := new(models.User)
	err := s.parseBody(c, user)
	if err != nil {
		return
	}
	//Validation of user fields
	apierror := user.Validate()
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	hash, err := argon.StringEncode(user.Password)
	if err != nil {
		apierror := errors.NewInternalServerError(err.Error())
		c.JSON(apierror.Status, apierror)
		return
	}
	user.Password = hash
	//request to datasource
	apierror = s.datasource.Create(user)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	c.String(http.StatusOK, "User successfuly created")
}
