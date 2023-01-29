// Quiz 1 - CMPS 2242

package main

// #9. Create function Square, will return area and perimter
func square(side float64) (float64, float64) {
	area := side * side
	perimeter := 4 * side
	return area, perimeter
}
