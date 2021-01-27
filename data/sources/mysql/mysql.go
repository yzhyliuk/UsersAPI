package mysql

import (
	"ms/usersAPI/data/models"
	"ms/usersAPI/utils/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     string = "root"
	password string = "s0meG0od@ndStr0ngP@ssWird"
)

//DataSource provides an connection to existing database
var DataSource *gorm.DB

//InitConnection : initialize connection to database
func InitConnection(dataSourceName string) *errors.APIError {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	DataSource = db
	err = DataSource.AutoMigrate(&models.User{})
	if err != nil {
		return errors.NewInternalServerError("Can't migrate SQL schema")
	}
	return nil
}
