// exercise-rot-reader.go
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyReader struct {
}

func (m MyReader) Read(b []byte) (int, error) {
	var i int
	for i = 0; i < cap(b); i++ {
		b[i] = 'A'
	}
	return i, nil
}

type rot13Reader struct {
	r io.Reader
}

func (ro *rot13Reader) Read(b []byte) (int, error) {
	p := make([]byte, cap(b))

	n, err := ro.r.Read(p)

	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
			b[i] = p[i] + 13
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') {
			b[i] = p[i] - 13
		} else {
			b[i] = p[i]
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	my := MyReader{}
	var test io.Reader
	test = my
	fmt.Println(test)

}
