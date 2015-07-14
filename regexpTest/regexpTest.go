package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	const filename = "./tmp/1.html"
	//const pattern = "<\\s?title\\s?>.?</title>"
	const pattern = "(?s)<(?i)head\\s*>(.*?)<(?i)title\\s*>"

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s", b)

	regPattern, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Errorf("regexp.Compile(%#v) error: %s", pattern, err)
	}
	findres := regPattern.FindAllIndex(b, -1)

	for i, v := range findres {
		s := []string{string(b[:v[1]]), "127.0.0.1;", string(b[v[1]:])}
		fmt.Println(i)
		fmt.Println(strings.Join(s, ""))
	}

}
