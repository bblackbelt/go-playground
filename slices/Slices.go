package slices

import "fmt"

mySlices := []string{3, 6}

func appendToSlices(name string) {
	mySlices.append(name)
	fmt.Println("Capacity: %d ", mySlices.cap)
}