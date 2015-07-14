package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
type error interface {
	Error() string
}
*/
type MyError struct {
	When time.Time
	What string
}

/*
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
*/

func run() *MyError {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("wwww")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err.Error())
	}
	fmt.Println("Converted integer:", i)
}
