// exercise-reader.go
package main

import (
	"code.google.com/p/go-tour/reader"
	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m *MyReader) Read(b []byte) (int, error) {
	n := cap(b)

	for i := 0; i < n; i++ {
		b[i] = 'A'
	}

	return n, nil
}

func main() {
	fmt.Println(MyReader{})
	//reader.Validate(MyReader{})
}
