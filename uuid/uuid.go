package main

import (
	"encoding/hex"
	"fmt"

	"github.com/bsm/go-guid"
	"github.com/satori/go.uuid"
)

func main() {
	//uuidtest()
	guidtest()
}

func uuidtest() {
	i := 0
	// Creating UUID Version 4
	for i < 1000 {
		u1 := uuid.NewV4()
		fmt.Printf("UUIDv4: %s\n", u1)
		i++
	}

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

func guidtest() {
	// Create a new 12-byte guid
	g1 := guid.New96()
	fmt.Println(hex.EncodeToString(g1.Bytes()))

	// Create a new 16-byte guid
	g2 := guid.New128()
	fmt.Println(hex.EncodeToString(g2.Bytes()))
}
