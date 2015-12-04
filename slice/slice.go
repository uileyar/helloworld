// slice.go
package main

import (
	"fmt"
)

func sum1(a *[]int) int {
	var sum int

	for i := 0; i < len(*a); i++ {
		sum += (*a)[i]
	}

	for i, v := range *a {
		(*a)[i] += v
	}

	return sum
}

func sum2(a []int) int {
	var sum int

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	for i, v := range a {
		a[i] += v
	}

	return sum
}

func main() {
	fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n", 0x2&0x2+0x4)
	//prints: 0x2 & 0x2 + 0x4 -> 0x6
	//Go:    (0x2 & 0x2) + 0x4
	//C++:    0x2 & (0x2 + 0x4) -> 0x2

	fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n", 0x2+0x2<<0x1)
	//prints: 0x2 + 0x2 << 0x1 -> 0x6
	//Go:     0x2 + (0x2 << 0x1)
	//C++:   (0x2 + 0x2) << 0x1 -> 0x8

	fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n", 0xf|0x2^0x2)
	//prints: 0xf | 0x2 ^ 0x2 -> 0xd
	//Go:    (0xf | 0x2) ^ 0x2
	//C++:    0xf | (0x2 ^ 0x2) -> 0xf
}
func main() {
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)

	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)

	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}

func main2() {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main1() {
	x := "test"
	xbytes := []byte(x)
	xbytes[0] = 'd'
	fmt.Println(x, string(xbytes))

	a := []int{1, 2, 3, 4}
	s := sum1(&a)
	fmt.Println(a, s)

	s = sum2(a[0:])
	s = sum2([]int{1, 2, 3, 4})
	fmt.Println(a, s)
}
