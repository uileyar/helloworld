// stringer.go
package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var b Stringer
	a := &Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	b = &z
	fmt.Println(b)
	fmt.Println(a, z)
}
