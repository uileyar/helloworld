package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/golang/glog"
)

var num int

func curltest() {
	out, err := exec.Command("curl", "-x", "127.0.0.1:8000", "http://m.baidu.com/").Output()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		num++
		glog.Infof("%d:%v", num, string(out))
	}

}

func nslookup() {
	command := "we-oppo.urlhunter.cn"
	//command := "we-dns.urlhunter.cn"

	out, err := exec.Command("nslookup", command).Output()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	glog.Infof("%v\n", string(out))
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	wtchan := time.Tick(time.Duration(1) * time.Microsecond)

	for {
		<-wtchan
		//nslookup()
		curltest()
	}
}
