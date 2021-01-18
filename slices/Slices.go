package slices

import "fmt"

var mySlices = make([]string, 3, 6)

/*AppendToSlices name to slice*/
func AppendToSlices(name string) {
	mySlices = append(mySlices, name)
	fmt.Printf("Capacity: %d ", cap(mySlices))
}
