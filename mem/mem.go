// mem.go
package main

import (
	"fmt"
	"sync"
)

var c = make(chan int)
var l sync.Mutex
var once sync.Once
var c2 = make(chan int, 2)
var a string

func f_channel() {
	a = "hello, world channel" //3
	<-c                        //4
}

func f_lock() {
	a = "hello, world lock" //4
	l.Unlock()              //5
}

func setup() {
	a = "hello, world"
	fmt.Println("once")
	c2 <- 0
}

func doprint() {
	once.Do(setup) //只会被调用一次！！！
	fmt.Println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	/*
		go f_channel() //1
		c <- 0         //2
		fmt.Println(a) //5

		l.Lock()       //1
		go f_lock()    //2
		l.Lock()       //3
		fmt.Println(a) //6
	*/
	twoprint()
	<-c2

}
