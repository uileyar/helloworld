package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Icon struct {
	Title string `xml:"title,attr"`
	Type  string `xml:"type,attr"`
	Name  string `xml:"name,attr"`
}
type Icons struct {
	Icon []Icon `xml:"icon"`
}

func main() {
	var v Icons
	xmlFile, err := ioutil.ReadFile("iconNames.xml")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	err1 := xml.Unmarshal(xmlFile, &v)
	fmt.Printf("%#v", v)
	if err1 != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
