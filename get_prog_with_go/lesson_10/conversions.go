package main

// go types are static once declared

import (
	f "fmt"
	m "math"
	"strconv"
)

func main() {
	f.Println("10.1: ")
	marsAge()
	f.Println("10.2: ")
	ariane()
	f.Println("10.3: ")
	runeConvert()
	f.Println("10.4: ")
	itoa()
	f.Println("10.5: ")
	launch()
	f.Println("10.6: ")
	toBool()
}

func marsAge() {
	age := 41
	marsDays := 687.0
	earthDays := 365.2425

	// errors with mismatched types
	// f.Println("I am", age*earthDays/marsDays, "years old on Mars.")

	// instead wrap the variable with the required type
	marsAge := float64(age)
	marsAge = marsAge * earthDays / marsDays
	f.Println("I am", marsAge, "years old on Mars.")
	// can convert from float to int as well, digits after decimal are truncated
	f.Println(int(earthDays))
	// conversions required between uint and int types, and differing sizes
	// always safe to convert to types with larger sizes, but smaller can hazard large risk
}

func ariane() {
	// if bh is 32768 which is too big for int16 it will wrap around to -32768
	var bh float64 = 32767
	var h = int16(bh)
	f.Println(h)
	// 'math' package has min/max constants that can detect if type conversion will be invalid
	if bh < m.MinInt16 || bh > m.MaxInt16 {
		// handle out of range value
		f.Println("Bh out of range")
	}

	// qc 10.3
	v := 42
	if v >= 0 && v <= m.MaxUint8 {
		v8 := uint8(v)
		f.Println("converted:", v8)
	}
}

func runeConvert() {
	// runes and bytes can be converted with the same syntax as numerics
	// gives same result as using '%c' format verb
	var (
		pi rune = 960
		alpha rune = 940
		omega rune = 969
		bang byte = 33		
	)
	f.Print(string(pi), string(alpha), string(omega), string(bang))
}

func itoa() {
	// converting digits to strings requires conversion to codepoints first
	// Itoa func from strconv can manage this
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	f.Println(str)
	// another method is to use the 'Sprintf' func, which returns a 'string' instead of displaying it
	countdown = 9
	str = f.Sprintf("Launch in T minus %v seconds.", countdown)
	f.Println(str)
	// strconv also has the inverse function 'Atoi' (ASCII to int), strings may return errors if too big
	countdown, err := strconv.Atoi("10")
	if err != nil {
		// oops, something went wrong
	}
	f.Println(countdown)
}

func launch() {
	// use Sprintf to convert bools to text or nums via if blocks
	launch := false

	launchText := f.Sprintf("%v", launch)
	f.Println("Ready for launch:", launchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	f.Println("Ready for launch:", yesNo)
}

func toBool() {
	// compiler errors if attempting to convert bool with string(false), int(false)
	// likewise for bool(1), bool("yes")
	yesNo := "no"	
	launch := (yesNo == "yes")
	f.Println("Ready for launch:", launch)
	// qc 10.5
	qcBool := true
	var boolToI int
	if qcBool {
		boolToI = 1
	} else {
		boolToI = 0
	}
	f.Printf("QC 10.5: %v | Type: %T\n", boolToI, boolToI)
}
