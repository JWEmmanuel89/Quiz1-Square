// Quiz 1 - CMPS 2242

package main

import (
	"fmt"
)

// #9. Create function Square, will return area and perimter
func square(side float64) (float64, float64) {
	area := side * side
	perimeter := 4 * side
	return area, perimeter
}

func main() {
	// #10. Test of square function
	result1, result2 := square(5.0)
	fmt.Println(result1, result2)
}
