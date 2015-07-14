// reader.go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %s \n", n, err, b)
		fmt.Printf("b[:%d] = %q\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}
