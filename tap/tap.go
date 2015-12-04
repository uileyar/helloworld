package main

import (
	"fmt"

	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
)

const BUFFERSIZE = 1522

func main() {
	ifce, err := water.NewTAP("tap0")
	fmt.Printf("%v, %v\n\n", err, ifce)
	buffer := make([]byte, BUFFERSIZE)
	for {
		_, err = ifce.Read(buffer)
		if err != nil {
			fmt.Printf("err = %v\n", err)
			break
		}
		ethertype := waterutil.MACEthertype(buffer)
		fmt.Printf("ethertype = %#v\n", ethertype)
		if ethertype == waterutil.IPv4 {
			packet := waterutil.MACPayload(buffer)
			if waterutil.IsIPv4(packet) {
				fmt.Printf("IPv4 Source:      %v [%v]\n", waterutil.MACSource(buffer), waterutil.IPv4Source(packet))
				fmt.Printf("IPv4 Destination: %v [%v]\n", waterutil.MACDestination(buffer), waterutil.IPv4Destination(packet))
				fmt.Printf("IPv4 Protocol:    %v\n\n", waterutil.IPv4Protocol(packet))
			} else if waterutil.IsIPv6(packet) {
				fmt.Printf("IPv6 Source:      %v\n", waterutil.MACSource(buffer))
				fmt.Printf("IPv6 Destination: %v\n\n", waterutil.MACDestination(buffer))
			} else {
				fmt.Printf("unknow Source:      %v\n", waterutil.MACSource(buffer))
				fmt.Printf("unknow Destination: %v\n\n", waterutil.MACDestination(buffer))
			}
		}
	}
}
