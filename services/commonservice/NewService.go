package commonservice

import (
	"ms/usersAPI/data/dao"
	"ms/usersAPI/data/models"
	"ms/usersAPI/services"
	"ms/usersAPI/utils/errors"
)

const (
	read   = 1 //read
	create = 2 //read, create
	update = 3 //read, create, update
	delete = 4 //read, create, update, delete

	worker    = 1
	financial = 2
	manager   = 3
	admin     = 4
)

//commonService serves requsts for UsersAPI
//
//contains methods as a request handlers for server
type commonService struct {
	datasource dao.DataAccessObject
}

//NewUserService : returns commonService with datasource with given credentials
func NewUserService(driver, DSNstring string) (services.Service, *errors.APIError) {
	dataAccessObject, err := dao.NewGeneralDataAccessObject(driver, DSNstring, models.User{})
	if err != nil {
		return nil, err
	}
	return &commonService{
		datasource: dataAccessObject,
	}, nil
}
