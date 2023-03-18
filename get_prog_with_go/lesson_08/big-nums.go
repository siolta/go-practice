package main

import (
	f "fmt"
	b "math/big"
)

const (
	lightSpeed = 299792 // km/s
	secondsPerDay = 86400
)

func main() {
	f.Println("8.1:")
	alpha()
	f.Println("8.2:")
	andromeda()
	f.Println("8.3:")
	constant()
}

func alpha() {
	var distance int64 = 41.3e12
	f.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	f.Println("That is", days, "days of travel at light speed.")
	// overflow uint64 with a number beyond 18quintillion
	// var distance uint64 = 24e18
	// exponent 'e' syntax right pads zeros, eg
	// var distance int = 56e6 // 56,000,000
}

func andromeda() {
	lightSpeed := b.NewInt(299792)	
	secondsPerDay := b.NewInt(86400)

	distance := new(b.Int)
	distance.SetString("24000000000000000000", 10)
	f.Println("Adromeda Galaxy is", distance, "km away.")

	seconds := new(b.Int)
	seconds.Div(distance, lightSpeed)

	days := new(b.Int)
	days.Div(seconds, secondsPerDay)

	f.Println("That is", days, "days of travel at light speed.")
}

func constant() {
	// constants and literal values can be untyped:
	f.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light days away")
}

// func constantTwo() {
// 	// constant and literal calculations are done at compile time
// 	// untyped numeric constants use the 'big' package, enabling usual operations with nums > 18quint
// 	const distance = 24000000000000000000
// 	const lightSpeed = 299792
// 	const secondsPerDay = 86400
	
// 	const days = distance / lightSpeed / secondsPerDay
// 	f.Println("Andromeda Galaxy is", days, "light days away.")
// 	// constants can be assigned to vars so long as they fit, 'int' is too small for 24quint, but 926k is fine
// 	km := distance // 24quint overflows int
// 	days := distance / lightSpeed / secondsPerDay // 926k
// }
