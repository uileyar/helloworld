// channels.go
package main

import (
	"fmt"
)

/*
默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。
*/
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
		fmt.Printf("c = %v, sum = %v, v = %v \n", c, sum, v)
		c <- sum // 将和送入 c
	}
	//c <- sum // 将和送入 c
	close(c)
}

func channelLoop() {
	c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {

	//channelLoop()
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 1) //提供第二个参数作为缓冲长度来初始化一个缓冲
	//向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。
	fmt.Println(len(c), cap(c))

	go sum(a[:len(a)/2], c)
	//go sum(a[len(a)/2:], c)

	for i := range c {
		fmt.Println(i)
	}

}
