package commonservice

import (
	"ms/usersAPI/data/dao"
	"ms/usersAPI/data/models"
	"ms/usersAPI/services"
	"ms/usersAPI/utils/errors"
)

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
