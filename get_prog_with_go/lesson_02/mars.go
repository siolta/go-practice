// my weight loss program
package main

import "fmt"

// main is the function when it all begins 
func main() {
	fmt.Printf("My weight on the surface of %v is %v lbs, and I would be %v years old.\n", 
	"Mars", 
	163.0 * 0.3873, 
	41 * 365 / 687)

	// numbers can be used as pad values, positive for left negative for right
	fmt.Printf("%-15v $%4v\n", "Spacex", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
