package apis

import (
	"fmt"
	"gin_blog/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",gin.H{
		"title": "About website",
	})
}

func AddChaApi(c *gin.Context)  {
	yb := c.Request.FormValue("yb")
	zh := c.Request.FormValue("zh")
	a := models.Arts{Yb: yb,Zh: zh}
	 ra ,err := a.AddCha()
	 if err != nil{
	 	log.Fatalln(err)
	 }
	 msg := fmt.Sprintf("insert successful %d",ra)
	 c.HTML(http.StatusOK,"list.html",gin.H{
	 	"msg":msg,
	 	"yb":yb,
	 	"zh":zh,
	 })
}
func GetChaApi(c *gin.Context)  {
	a := new(models.Arts)
	ar,err:= a.GetCha()
	//ra ,err := a.AddCha()
	if err != nil{
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("%d",ar)
	//print(msg)
	c.HTML(http.StatusOK,"login.html",gin.H{
		"msg":msg,
	})
}

func GetName(c *gin.Context){
	name := c.Param("name")
	action := c.Param("action")
	c.String(http.StatusOK,"HELLO %s,%s",name,action)
}

func First(c *gin.Context){
	first := c.DefaultQuery("first","aaa")
	c.String(http.StatusOK,"HELLO %s",first)
}