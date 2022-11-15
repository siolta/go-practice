package main
import f "fmt"

func main() {
	command := "go east"

	if command == "go east" {
		f.Println("You head further up the mountain.")
	} else if command == "go inside" {
		f.Println("You enter the cave where you live out the rest of your life.")
	} else {
		f.Println("Didn't quite get that.")
	}	
}
