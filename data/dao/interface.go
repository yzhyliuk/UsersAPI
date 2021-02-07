package dao

import "ms/usersAPI/utils/errors"

//DataAccessObject : defines an interface of Services
type DataAccessObject interface {
	Create(obj interface{}) *errors.APIError
	Recive(obj interface{}, primaryKey int) (interface{}, *errors.APIError)
	Update(obj interface{}) *errors.APIError
	UpdateWhere(obj interface{}, params interface{}) *errors.APIError
	Delete(obj interface{}, primaryKey int) *errors.APIError
	DeleteWhere(obj interface{}, params interface{}) *errors.APIError
	FindAll(obj interface{}, params interface{}) *errors.APIError
}
