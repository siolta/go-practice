package main

import "fmt"

type rect struct {
	width, height int
}

// this area method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined with pointer or value recievers
// this perim() method has a value reciever type
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// call the 2 methods defined for the rect struct
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// go will handle conversion between values and pointers for method calls
	// may be useful to use a pointer reciever to avoid copying 
	// or allow the method to mutate the recieving struct
	rp := &r
	fmt.Println("area_p: ", rp.area())
	fmt.Println("perim_p: ", rp.perim())
}
