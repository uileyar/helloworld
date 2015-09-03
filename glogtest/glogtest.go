// glogtest.go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"./glog"
)

const (
	//DEFLOGPATH  string = "/var/log/glogtest1"
	//DEFMOVEPATH string = "/var/log/glogtest2"

	DEFLOGPATH  string = "/tmp/glogtest1"
	DEFMOVEPATH string = "/home/john/glogtest2"
)

func StartCPUProfile() {
	filename := "cpu-" + strconv.Itoa(os.Getegid()) + ".pprof"
	f, err := os.Create(filename)
	if err != nil {
		glog.Fatal("record cpu profile failed: ", err)
	}
	pprof.StartCPUProfile(f)
	//time.Sleep(time.Duration(sec) * time.Second)

	fmt.Printf("create cpu profile %s \n", filename)
}

func StopCPUProfile() {
	pprof.StopCPUProfile()
}

func init() {
	//flag.Set("logtostderr", "false")
	flag.Set("logtostderr", "true")
	flag.Set("log_dir", DEFLOGPATH)
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	glog.Infof("NumCPU = %d", runtime.GOMAXPROCS(0))
}

func logPrint(c chan int, i int) {
	//time.Sleep(time.Duration(1) * time.Millisecond)
	//max := 1000
	j := 0
	for ; j < 100; j++ {
		j++
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		glog.Infof("%d, %d ", i, j)
		glog.Statisf("%d, %d ", i, j)
		//glog.Errorf("%d, %d ", i, j)
		glog.Warningf("%d, %d ", i, j)
		//glog.Fatalf("%d, %d ", i, j)
	}

	c <- 1
}

func main() {

	//StartCPUProfile()
	//defer StopCPUProfile()

	fmt.Println("Hello World!")
	n := 10
	c := make(chan int, n)
	for i := 0; i < n; i++ {
		go logPrint(c, i)
	}
	for i := 0; i < 1; i++ {
		go MoveAndCreateTest()
	}

	for {
		select {
		case <-c:
			if n--; n <= 0 {
				fmt.Println("go log fin!!!\n")
				return
			}
		}
	}
}

func MoveAndCreateTest() {
	for {
		select {
		case <-time.After(time.Duration(3) * time.Second):
			if err := glog.MoveAndCreateNewFiles(); err != nil {
				fmt.Printf("MoveAndCreateNewFiles err = %v\n", err)
				return
			}

			fmt.Printf("MoveAndCreateNewFiles fin\n")
		}
	}
	return
}
