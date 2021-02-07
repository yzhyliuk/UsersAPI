package commonservice

import (
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/argon"
	"ms/usersAPI/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Login : handler that recives user login credentials and verify them returning access token
func (s *commonService) Login(c *gin.Context) {
	user := new(models.User)
	err := c.BindJSON(user)
	if err != nil {
		apierr := errors.NewInternalServerError("Unable to parse user credential from request body")
		c.JSON(apierr.Status, apierr)
		return
	}
	userCurrent := new(models.User)
	apierr := s.datasource.FindAll(&userCurrent, models.User{Email: user.Email})
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}
	if userCurrent.Email == "" {
		apierr := errors.NewBadRequestError("Incorrect user password or email")
		c.JSON(apierr.Status, apierr)
		return
	}

	ok, _ := argon.CompareStringToHash(user.Password, userCurrent.Password)
	if ok {
		c.JSON(http.StatusOK, userCurrent.UserData())
		return
	}
	apierr = errors.NewBadRequestError("Incorrect user password or email")
	c.JSON(apierr.Status, apierr)
}
