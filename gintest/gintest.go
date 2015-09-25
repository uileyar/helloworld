// gintest.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.GET("/welcome", func(c *gin.Context) {
		log.Printf("GET welcome\n")
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.GET("/sina", func(c *gin.Context) {
		log.Printf("GET sina\n")
		c.Redirect(http.StatusMovedPermanently, "http://www.sina.cn/")

	})
	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
