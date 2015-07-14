// exercise-errors.go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

/*
 在 Error 方法内调用 fmt.Sprint(e) 将会让程序陷入死循环。可以通过先转换 e 来避免这个问题：`fmt.Sprint(float64(e))`。
*/
func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	} else {
		return fmt.Sprintln(float64(e))
	}
}

func Sqrt(x float64) (float64, error) {
	return math.Sqrt(x), ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
