package user

import (
	. "github.com/Kirk-Wang/Hello-Gopher/api-16/handler"
	"github.com/Kirk-Wang/Hello-Gopher/api-16/pkg/errno"
	"github.com/Kirk-Wang/Hello-Gopher/api-16/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// List list the users in the database.
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
