package main

import (
	f "fmt"
)

func main() {
	str := "true"
	
	var toBool bool
	
	switch str {
	case "true", "yes", "1":
		toBool = true
		f.Println("Converted value:", toBool)
	case "false", "no", "0":
		toBool = false
		f.Println("Converted value:", toBool)
	default:
		f.Println("ERROR:", str, "is not truthy")
	}
}
