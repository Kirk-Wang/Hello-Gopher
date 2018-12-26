package user

import (
	"strconv"

	. "github.com/Kirk-Wang/Hello-Gopher/api-11/handler"
	"github.com/Kirk-Wang/Hello-Gopher/api-11/model"
	"github.com/Kirk-Wang/Hello-Gopher/api-11/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
