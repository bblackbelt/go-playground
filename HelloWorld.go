package main

import (
	s "HelloWorld/slices"
	"fmt"
)

func init() {
	fmt.Println("Hello Init")
}

func main() {
	fmt.Println("Hello World")
	s.AppendToSlices("test")
}
