package mdlwr

import (
	"ms/usersAPI/utils/errors"

	"github.com/gin-gonic/gin"
)

//CompanyBorder : check if company border is violated
func CompanyBorder(c *gin.Context, companyBorder int) bool {
	cid, _ := c.Get("companyid")
	companyid := int(cid.(float64))
	if companyid != companyBorder {
		apierr := errors.NewUnauthorizedError("Company border is violated")
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return true
	}
	return false
}
