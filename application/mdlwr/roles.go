package mdlwr

import (
	"fmt"
	"ms/usersAPI/utils/errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//RequiredPermission : Chech if this user allowed to perform current action
func RequiredPermission(c *gin.Context, roleRequred int, permRquired int) bool {
	userrole, _ := c.Get("role")
	str := strings.Split(fmt.Sprintf("%s", userrole), "-")
	role, err := strconv.Atoi(str[0])
	perm, err := strconv.Atoi(str[1])
	if err != nil {
		apierr := errors.NewUnauthorizedError("Role is not correct")
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return false
	}
	if role < roleRequred {
		apierr := errors.NewUnauthorizedError("Action not allowed")
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return false
	}
	if perm < permRquired {
		apierr := errors.NewUnauthorizedError("Action not allowed")
		c.AbortWithStatusJSON(apierr.Status, apierr)
		return false
	}
	return true
}
