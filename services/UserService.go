package services

import (
	"fmt"
	"ms/usersAPI/data/dao"
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//userService serves requsts for UsersAPI
//
//contains methods as a request handlers for server
type userService struct {
	datasource dao.DataAccessObject
}

//NewUserService : returns userservice with datasource with given credentials
func NewUserService(dataSource, ConnectionString string) (Service, *errors.APIError) {
	service := new(userService)
	err := service.configDataSource(dataSource, ConnectionString)
	if err != nil {
		return nil, err
	}
	return service, nil
}

//Revice : recive user from datasource for given primary key
func (s *userService) Recive(c *gin.Context) {
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
func (s *userService) Create(c *gin.Context) {
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
	//request to datasource
	apierror = s.datasource.Create(user)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}

	c.String(http.StatusOK, "User successfuly created")
}

func (s *userService) Update(c *gin.Context) {
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

func (s *userService) Delete(c *gin.Context) {
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

func (s *userService) List(c *gin.Context) {
	usersList := new(models.UserList)
	apierror := s.datasource.List(usersList)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.JSON(http.StatusOK, usersList.Export())
}

//FindsAll : returns JSON list of all objects that mathes given query parameter
func (s *userService) FindAll(c *gin.Context) {
	userParams := new(models.User)
	apierr := s.queryToStruct(c, userParams)
	if apierr != nil {
		c.JSON(apierr.Status, apierr)
		return
	}
	usersList := new(models.UserList)
	apierror := s.datasource.FindAll(usersList, userParams)
	if apierror != nil {
		c.JSON(apierror.Status, apierror)
		return
	}
	c.JSON(http.StatusOK, usersList.Export())
}
