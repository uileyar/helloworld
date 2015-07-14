// web-servers.go
package main

//package http

import (
	"fmt"
	"log"
	"net/http"
)

type String string

/*
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
*/

type Hello struct {
	Greeting string
	Punct    string
	Who      string
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "String!")
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/Hello", &Hello{"Hello", ":", "Gophers!"})

	//var h Hello
	err := http.ListenAndServe("localhost:4000", nil)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

}
