package main

import (
	"fmt"
)

func init() {
	fmt.Println("Hello Init")
}

func main() {
	fmt.Println("Hello World")
	slices.appendToSlices("test")
}
