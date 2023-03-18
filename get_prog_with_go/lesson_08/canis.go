package main

import (
	f "fmt"
)

const (
	lightSpeed = 299792 // km/s
	secondsPerDay = 86400
	daysPerYear = 365
	distance = 236000000000000000
	lightYears = distance / lightSpeed / secondsPerDay / daysPerYear
)

func main() {
	f.Println("Canis is", lightYears, "light years away.")
}
