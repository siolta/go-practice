package main

import (
	f "fmt"
	t "time"
	r "math/rand"
)

func main() {
	count := 10

	for count > 0 {
		f.Println(count)
		t.Sleep(t.Second)
		count--
	}

	f.Println("LiftOff!!!")

	// QC 3.6
	for count > 0 {
		f.Println(count)
		t.Sleep(t.Second)
		if r.Intn(100) == 0 {
			break
		}
		count--
	}

	if count == 0 {
		f.Println("Liftoff!")
	} else {
		f.Println("Launch Failed.")
	}

}
