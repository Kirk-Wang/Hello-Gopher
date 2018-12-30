package user

import (
	. "github.com/Kirk-Wang/Hello-Gopher/api-14/handler"
	"github.com/Kirk-Wang/Hello-Gopher/api-14/model"
	"github.com/Kirk-Wang/Hello-Gopher/api-14/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
