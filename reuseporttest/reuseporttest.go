// reuseport.go
package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"

	//"github.com/kavu/go_reuseport"
)

var num int

func test1() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	listener, err := NewReusablePortListener("tcp4", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(os.Getpid())
		num++
		fmt.Fprintf(w, "Hello, %v, %v, %q\n", num, os.Getpid(), html.EscapeString(r.URL.Path))
	})

	panic(server.Serve(listener))
}

func test2() {
	//net.USE_SO_REUSEPORT = true
	//ln, err := NewReusablePortListener("tcp4", ":8000")

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	s := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(os.Getpid())
		num++
		fmt.Fprintf(w, "Hello, %v, %v, %q\n", num, os.Getpid(), html.EscapeString(r.URL.Path))
	})
	panic(s.Serve(ln))
}

func main() {
	test2()
}
