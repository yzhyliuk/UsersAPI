package mdlwr

import (
	"encoding/base64"
	"encoding/json"
	"ms/usersAPI/utils/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const authServer = "http://localhost:1000"

//AuthenticationRequired : authentication middleware
func AuthenticationRequired(c *gin.Context) {
	client := resty.New()
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		apierr := errors.NewUnauthorizedError("Session has expired. Cookie dosen't set")
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return
	}
	response, err := client.R().SetHeader("Content-Type", "application/json").SetCookie(cookie).Post(authServer + "/verify")
	if err != nil {
		apierr := errors.NewInternalServerError(err.Error())
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return
	}
	if response.StatusCode() != http.StatusOK {
		apierr := new(errors.APIError)
		err := json.Unmarshal(response.Body(), apierr)
		if err != nil {
			apierr = errors.NewInternalServerError(err.Error())
		}
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return
	}
	TokenParse(c)
	c.Next()
}

//TokenParse : parses toke values into context
func TokenParse(c *gin.Context) {
	jwtString, _ := c.Cookie("token")
	payload := strings.Split(jwtString, ".")[1]
	data, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		apierr := errors.NewInternalServerError(err.Error())
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return
	}
	keyValue := make(map[string]interface{})
	json.Unmarshal(data, &keyValue)
	for key, val := range keyValue {
		c.Set(key, val)
	}
}
