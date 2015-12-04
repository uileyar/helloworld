// gintest.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	responW interface {
		http.ResponseWriter
		http.Hijacker
	}
	responseWriter struct {
		http.ResponseWriter
		size   int
		status int
	}
	responseWriterA struct {
		http.ResponseWriter
		numA    int
		sizeA   int
		statusA int
	}
)

var _ responW = &responseWriter{}
var _ responW = &responseWriterA{}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	log.Printf("responseWriter Hijack")
	if w.size < 0 {
		w.size = 0
	}

	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *responseWriterA) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	log.Printf("responseWriterA Hijack")
	if w.sizeA < 0 {
		w.sizeA = 0
	}
	hjack, ok := w.ResponseWriter.(http.Hijacker)
	if ok {
		return hjack.Hijack()
	}
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func arraryTest() {
	x := [3]int{1, 2, 3}
	y := x[:]
	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(y)

	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])
}

func main() {
	//var Writer responW
	//log.Printf("%v\n", nameOfFunction(Writer))
	fmt.Println(strings.TrimSpace(" \t\n a lone  gopher \n\t\r\n")) // a lone gopher
	arraryTest()
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
