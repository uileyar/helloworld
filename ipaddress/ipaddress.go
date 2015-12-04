// ipaddress.go
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")
	//interfaces, err := net.Interfaces()
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("err = %v", err)
	}
	for i, v := range interfaces {
		fmt.Println(i, v)
		fmt.Printf("%v, %#v\n", i, v)
	}
}
