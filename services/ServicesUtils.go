package services

import (
	"encoding/json"
	"ms/usersAPI/utils"
	"ms/usersAPI/utils/errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

//queryToStruct : parse query parameter into a struct
func (s *userService) queryToStruct(c *gin.Context, obj interface{}) *errors.APIError {
	//Convert struct to map
	fields, err := utils.StructToMap(obj, "json")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//Iterate over fields of struct
	for key, value := range fields {
		//if password in query - stor request
		//Serching for params that match with fields name
		v, OK := c.GetQuery(key)
		if v == "password" {
			return errors.NewBadRequestError("Forbidden request parameter")
		}
		//if queryparameter exists add it's value to map
		if OK {
			switch reflect.TypeOf(value) {
			case reflect.TypeOf(0):
				fields[key], _ = strconv.Atoi(v)
				break
			case reflect.TypeOf(""):
				fields[key] = v
				break
			}
		}
	}
	//Converting map to json
	userString, err := json.Marshal(fields)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//Converting json back to struct
	err = json.Unmarshal(userString, obj)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
