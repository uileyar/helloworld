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

	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "false")
	//flag.Set("logtostderr", "true")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	glog.Infof("NumCPU = %d", runtime.GOMAXPROCS(0))
}

func logPrint(c chan int, i int) {
	//time.Sleep(time.Duration(1) * time.Millisecond)
	max := 1000
	for j := 0; j < max; j++ {
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		glog.Infof("%d, %d ", i, j)
	}

	c <- 1
}

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

func main() {

	StartCPUProfile()
	defer StopCPUProfile()

	fmt.Println("Hello World!")
	n := 10000
	c := make(chan int, n)
	for i := 0; i < n; i++ {
		go logPrint(c, i)
	}

	for {
		select {
		case <-c:
			if n--; n <= 0 {
				return
			}
		}
	}
}
