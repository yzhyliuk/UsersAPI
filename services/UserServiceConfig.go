package services

import (
	"fmt"
	"ms/usersAPI/data/dao"
	"ms/usersAPI/data/sources/mysql"
	"ms/usersAPI/data/sources/postgres"
	"ms/usersAPI/utils/errors"
)

func (s *userService) configDataSource(dataSource, ConnectionString string) *errors.APIError {
	switch dataSource {
	case "mysql":
		err := mysql.InitConnection(ConnectionString)
		if err != nil {
			return err
		}
		s.datasource = dao.NewGeneralDataService(mysql.DataSource)
		return nil
	case "postgres":
		err := postgres.InitConnection(ConnectionString)
		if err != nil {
			return err
		}
		s.datasource = dao.NewGeneralDataService(postgres.DataSource)
		return nil
	default:
		return errors.NewInternalServerError(fmt.Sprintf("Data source '%s' is not avaliable for this app", dataSource))
	}
}
