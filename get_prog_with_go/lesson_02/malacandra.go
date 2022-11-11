package main

import (
	f "fmt"
)

const (
	malacandraDistance = 56000000
	hoursPerEarthDay = 24
)

var travelDays = 28

func main() {
	f.Printf("ship would need to travel at %v km/h to reach Malacandra in %v days.\n", malacandraDistance/travelDays/hoursPerEarthDay,travelDays)
}
