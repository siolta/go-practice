package main

import (
	f "fmt"
	s "strings"
)

func main() {
	f.Println("You find yourself in a dimly lit cavern.")

	var (
		command = "walk outside"
		wearShades = true
		exit = s.Contains(command, "outside")
		read = s.Contains(command, "read")
	)

	f.Println("You leave the cave:", exit)
}
