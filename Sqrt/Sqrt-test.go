// Sqrt-test
package main

import (
	"fmt"
	"math"
	"time"
)

var z = float64(1)

func Sqrt(x float64) float64 {
	zq := z - (z*z-x)/(2*z)

	if math.Abs(z-zq) < 0.0000000001 {
		return zq
	} else {
		z = zq
	}

	return Sqrt(x)
}

func main() {
	var num float64 = 6

	defer fmt.Println(Sqrt(num))
	defer fmt.Println(math.Sqrt(num))

	defer fmt.Println(time.Now().Zone())

	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("two days")
	default:
		fmt.Println("today is ", today)
	}
}
