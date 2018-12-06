package main

import (
	"fmt"
	"gin_blog/apis"
	"gin_blog/middlewares"

	//"gin_blog/middlewares"
	"github.com/gin-gonic/gin"
	//. "aze.orgd/apis"
)

func intiRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.CookieMiddleWare())
	//router := gin.Default()
	router.GET("/readcookie", func(context *gin.Context) {
		val,_ := context.Cookie("name")
		fmt.Println(val)
		context.String(200, "Cookie:%s", val)
	})
	router.GET("/write_cookie", func(context *gin.Context) {
		context.SetCookie("name", "sasasa", 10, "/", context.Request.Host, false, true)
	})
	router.GET("/",apis.IndexApi)
	router.POST("/art",apis.AddChaApi)
	router.GET("/arts",apis.GetChaApi)
	router.GET("/ass/:name",apis.GetName)
	router.GET("/ass/:name/*action",apis.GetName)
	router.GET("/first",apis.First)
	//router.GET("arts",apis.)
	router.LoadHTMLGlob("templates/*")
	return router
}

