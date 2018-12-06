package middlewares

import (
	"github.com/gin-gonic/gin"
)

func  CookieMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.SetCookie("name", "sasasa", 10, "/", c.Request.Host, false, true)
		c.Writer.Header().Set("Request-Ip",c.ClientIP())
		c.Writer.Header().Set("ss","sss")
		c.Next()
	}
}
