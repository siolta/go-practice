package main 

import (
	f "fmt"
	u "unicode/utf8"
)

func main() {
	// zero value for strings is "" < empty string
	f.Println("9.1:")
	raw()
	f.Println("9.4:")
	runes()
	f.Println("9.5:")
	index()
	f.Println("9.6:")
	ceaser()
	f.Println("9.7:")
	rot13EN()
	f.Println("9.9:")
	spanish()
}

func raw() {
	// backticks can be used for raw strings
	f.Println("peace be upon you\nupon you be peace")
	f.Println(`strings can span multiple lines with the \n escape sequence`)

	// raw literals can span multiple lines of code
	f.Println(`
	  peace be upon you
	  upon you be peace`)

	// both are type string
	f.Printf("%v is a %[1]T\n", "literal string")
	f.Printf("%v is a %[1]T\n", `raw string literal`)
}

func runes() {
	// rune is an alias for int32
	// byte is an alias for uint8
	var (
		pi rune = 960
		alpha rune = 940
		omega rune = 969
		bang byte = 33
	)

	f.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	// use %c to print characters instead of values
	f.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	// go can also use char literals by enclosing with single quotes 'A'
	// go infers type rune if none specified
}

func index() {
	// go strings are immutable once declared, use indexes to access bytes of a string
	message := "shalom"
	c := message[5]
	f.Printf("%c\n", c)

	// qc 9.3
	for b := 0; b < 6; b++ {
		f.Printf("%c\n", message[b])
	}
}

func ceaser() {
	c := 'a'
	// need to account for x,y,z and wrap
	if c > 'z' {
		c = c - 26
	}
	c = c + 3
	f.Printf("%c\n", c)
	// shift to caps by subtracting 32 from a lowercase char
	f.Printf("%c : %c\n", 'g', 103 - 32)
}

func rot13EN() {
	// builtin len() func returns the length of various types, for strings it returns byte length
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		f.Printf("%c", c)
	}
}

func spanish() {
	// utf8 package has funcs to return len of a string in runes instead of bytes
	question := "¿Cómo estás?"
	f.Println(len(question), "bytes")
	f.Println(u.RuneCountInString(question), "runes")

	c, size := u.DecodeRuneInString(question)
	f.Printf("First rune: %c %v bytes\n", c, size)

	// 9.9 : range can decode utf8 encoded strings
	// can use _ if index isn't needed
	for i, c := range question {
		f.Printf("%v %c\n", i, c)
	}
}
