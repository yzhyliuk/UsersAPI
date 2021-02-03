package dao

import (
	"fmt"
	"ms/usersAPI/utils/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//NewGeneralDataAccessObject : Returns new data access object connected to given datasource
func NewGeneralDataAccessObject(driver string, DSNstring string, model interface{}) (DataAccessObject, *errors.APIError) {
	//Decalring db
	var db *gorm.DB

	//Connect using driver
	switch driver {
	case "postgres":
		client, err := gorm.Open(postgres.Open(DSNstring), &gorm.Config{})
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		db = client
	default:
		return nil, errors.NewInternalServerError(fmt.Sprintf("Driver '%s' is not avaliable for this app", driver))
	}
	err := db.AutoMigrate(model)
	if err != nil {
		return nil, errors.NewInternalServerError("Can't update or create a schema for given model")
	}
	return &generalDataAccessObject{
		db: db,
	}, nil
}
