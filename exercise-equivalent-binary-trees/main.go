// 1 project main.go
package main

import (
	"fmt"
	"runtime"
	"tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	WalkLoop(t, ch)
	close(ch)
}

func WalkLoop(t *tree.Tree, ch chan int) {

	//runtime.Gosched() // 显式地让出CPU时间给其他goroutine

	if t == nil {
		return
	}

	if t.Left != nil {
		WalkLoop(t.Left, ch)
	}

	ch <- t.Value
	fmt.Printf("ch = %v,Value = %v \n", ch, t.Value)

	if t.Right != nil {
		WalkLoop(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	var flag bool
	var s1 string
	var s2 string

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		s1 += fmt.Sprint(i) + " "
	}

	for i := range ch2 {
		s2 += fmt.Sprint(i) + " "
	}

	if s1 == s2 {
		flag = true
	}

	fmt.Println(flag)
	return flag
}

func main() {
	runtime.GOMAXPROCS(2) // 最多使用2个核

	fmt.Println(runtime.NumCPU())
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(tree.New(1), ch1)
	go Walk(tree.New(2), ch2)

	for i := range ch1 {
		fmt.Sprint(i)
	}

	for i := range ch2 {
		fmt.Sprint(i)
	}

	//Same(tree.New(2), tree.New(2))
	//Same(tree.New(2), tree.New(3))
}
