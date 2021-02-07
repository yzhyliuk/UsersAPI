package commonservice

import (
	"encoding/json"
	"fmt"
	"ms/usersAPI/utils"
	"ms/usersAPI/utils/errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

//queryToStruct : parse query parameter into a struct
func (s *commonService) queryToStruct(c *gin.Context, obj interface{}) *errors.APIError {
	//Convert struct to map
	fields, err := utils.StructToMap(obj, "json")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	emptyQuery := true
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
			emptyQuery = false
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
	if emptyQuery {
		return errors.NewBadRequestError("Empty query, or wrong parameter name")
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

//parseID : parses id from the context and returns it as integer. Returning -1 if can't converte
func (s *commonService) parseID(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		apierror := errors.NewBadRequestError("Invalid path parameter of 'id', should be grater than 0")
		c.JSON(apierror.Status, apierror)
		return -1
	}
	return id
}

//parseBody : parses body from request
func (s *commonService) parseBody(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		apierror := errors.NewBadRequestError("Can't parse request body into struct")
		c.JSON(apierror.Status, apierror)
		return fmt.Errorf("Parsing error")
	}
	return nil
}

func (s *commonService) getCompanyID(c *gin.Context) int {
	//Setting company borders
	cid, _ := c.Get("companyid")
	return int(cid.(float64))
}

func (s *commonService) getUserID(c *gin.Context) int {
	//Setting company borders
	cid, _ := c.Get("userid")
	return int(cid.(float64))
}
