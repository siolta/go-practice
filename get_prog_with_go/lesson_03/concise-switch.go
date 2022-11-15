package main

import f "fmt"

func main() {
	f.Println("There is a cavern entrance here and a path to the east.")
	command := "go inside"

	switch command {
	case "go east":
		f.Println("You head further up the mountain.")
	case "enter cave", "go inside": // comma delimited list accepted in case statements
		f.Println("You find yourself in a dimly lit cavern")
	case "read sign":
		f.Println("The sign reads 'No Minors'.")
	default:
		f.Println("Didn't quite get that")
	}

	// 3.7

	room := "lake"

	switch room {
	case "cave":
		f.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		f.Println("The ice seems solid enough.")
		fallthrough // perform the next case as well
	case "underwater":
		f.Println("The Water is freezing cold.")
	}
}
