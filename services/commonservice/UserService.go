package commonservice

import (
	"fmt"
	"ms/usersAPI/data/dao"
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/argon"
	"ms/usersAPI/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//commonService serves requsts for UsersAPI
//
//contains methods as a request handlers for server
type commonService struct {
	datasource dao.DataAccessObject
}

//Revice : recive user from datasource for given primary key
func (s *commonService) Recive(c *gin.Context) {
	//geting primary key as path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		apierror := errors.NewBadRequestError("Invalid path parameter of 'id', should be grater than 0")
		c.JSON(apierror.Status, apierror)
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
	//return json
	c.JSON(http.StatusOK, user.Export())
}

//Create : creates new entry of user in datasource
func (s *commonService) Create(c *gin.Context) {
	//retrieving user from request body
	user := new(models.User)
	err := c.ShouldBindJSON(user)
	if err != nil {
		apierror := errors.NewBadRequestError("Can't parse request body into User struct")
		c.JSON(apierror.Status, apierror)
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

func (s *commonService) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		apierror := errors.NewBadRequestError("Invalid path parameter of 'id', should be grater than 0")
		c.JSON(apierror.Status, apierror)
		return
	}
	user := new(models.User)
	err = c.ShouldBindJSON(user)
	if err != nil {
		apierror := errors.NewBadRequestError("Can't parse request body into User struct")
		c.JSON(apierror.Status, apierror)
		return
	}
	currentUser := new(models.User)
	_, apierror := s.datasource.Recive(currentUser, id)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	user.ID = id
	user.Merge(currentUser)
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

func (s *commonService) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.String(http.StatusBadRequest, "Invalid path parameter of 'id', should be grater than 0")
		return
	}
	emptyUser := new(models.User)
	apierror := s.datasource.Delete(emptyUser, id)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("User with id %d succesfully deleted", id))
}

//UpdateWhere : updates all rows that matches with given parameters
func (s *commonService) UpdateWhere(c *gin.Context) {
	user := new(models.User)
	err := s.parseBody(c, user)
	if err != nil {
		return
	}
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
	//request to datasource
	apierror := s.datasource.UpdateWhere(user, userParams)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	c.String(http.StatusOK, "resources successfully updated")
}

//FindsAll : returns JSON list of all objects that mathes given query parameter
func (s *commonService) FindAll(c *gin.Context) {
	//if query parameters aren't set
	params := new(models.User)
	apierr := s.queryToStruct(c, params)
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}
	usersList := new(models.UserList)
	apierror := s.datasource.FindAll(usersList, params)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.JSON(http.StatusOK, usersList.Export())
}

//Login : handler that recives user login credentials and verify them returning access token
func (s *commonService) Login(c *gin.Context) {
	user := new(models.User)
	err := c.BindJSON(user)
	if err != nil {
		apierr := errors.NewInternalServerError("Unable to parse user credential from request body")
		c.JSON(apierr.Status, apierr)
		return
	}
	fmt.Println(user)
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
		c.String(http.StatusOK, "Password Ok!")
		return
	}
	apierr = errors.NewBadRequestError("Incorrect user password or email")
	c.JSON(apierr.Status, apierr)
}
