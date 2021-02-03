package dao

import (
	"ms/usersAPI/utils/errors"

	"gorm.io/gorm"
)

//Postservice :
type generalDataAccessObject struct {
	db *gorm.DB
}

//AddDataSource : adds data source for given service

//Create : creates a new entry of row in db
func (s *generalDataAccessObject) Create(obj interface{}) *errors.APIError {
	result := s.db.Create(obj)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

//Recive : gets data from db by primary key
func (s *generalDataAccessObject) Recive(obj interface{}, primaryKey int) (interface{}, *errors.APIError) {
	result := s.db.First(obj, primaryKey)
	if result.Error != nil {
		return nil, errors.NewBadRequestError(result.Error.Error())
	}
	return obj, nil
}

//Update : updates object
func (s *generalDataAccessObject) Update(obj interface{}) *errors.APIError {
	result := s.db.Save(obj)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

//UpdateWhere : updates all objects matchig parameters
func (s *generalDataAccessObject) UpdateWhere(obj interface{}, params interface{}) *errors.APIError {
	result := s.db.Model(obj).Where(params).Updates(obj)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

//Delete : deletes first entry of object with given primary key
func (s *generalDataAccessObject) Delete(obj interface{}, primaryKey int) *errors.APIError {
	result := s.db.Delete(obj, primaryKey)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

//DeleteWhere : deletes first entry of object with given primary key
func (s *generalDataAccessObject) DeleteWhere(obj interface{}, params interface{}) *errors.APIError {
	result := s.db.Model(obj).Where(params).Delete(obj)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

//List : returns list of all Users
func (s *generalDataAccessObject) List(obj interface{}) *errors.APIError {
	result := s.db.Find(obj)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}

//FindAll : finds all objects that match given parameters
func (s *generalDataAccessObject) FindAll(obj interface{}, params interface{}) *errors.APIError {
	result := s.db.Find(obj, params)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}
