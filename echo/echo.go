// echo.go
package main

import (
	"flag" // command line option parser
	"fmt"
	"os"
)

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	var s string = ""
	var omitNewline = flag.Bool("n", false, "don't print final newline")

	flag.Parse() // Scans the arg list and sets up flags

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline
	}

	os.Stdout.WriteString(s)
	fmt.Println(*omitNewline)
}
