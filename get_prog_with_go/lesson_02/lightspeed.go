// How long does it take to get to mars?
package main

import f "fmt"

func main() {
	const (
		lightSpeed = 299792 // km/s
		transportSpeed = 100800 // km/h
		hoursPerEarthDay = 24 // hours per day
	)

	var distance = 56000000   // km

	f.Println(distance/lightSpeed, "seconds")

	distance = 40100000
	f.Println(distance/lightSpeed, "seconds")

	distance = 96300000 // transport launch distance
	f.Println("transport would take", distance/transportSpeed/hoursPerEarthDay, "days")
}
