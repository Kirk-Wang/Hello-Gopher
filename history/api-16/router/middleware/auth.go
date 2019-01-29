package middleware

import (
	"github.com/Kirk-Wang/Hello-Gopher/api-16/handler"
	"github.com/Kirk-Wang/Hello-Gopher/api-16/pkg/errno"
	"github.com/Kirk-Wang/Hello-Gopher/api-16/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
