// number parsing in go with strconv
package main

import (
	"fmt"
	"strconv"
)

 func main() {
	 // 64 is the bits of precision to parse
	 f, _ := strconv.ParseFloat("1.234", 64)
	 fmt.Println(f)

	 // 0 means infer the base from the string, 64 requires result fit in 64 bits
	 i, _ := strconv.ParseInt("123", 0, 64)
	 fmt.Println(i)

	 // ParseInt recognizes hex formatted numbers
	 d, _ := strconv.ParseInt("0x1c8", 0, 64)
	 fmt.Println(d)

	 // uints also available
	 u, _ := strconv.ParseUint("789", 0, 64)
	 fmt.Println(u)

	 // atoi for base-10 parsing
	 k, _ := strconv.Atoi("135")
	 fmt.Println(k)

	 // parse returns an error on bad input
	 _, e := strconv.Atoi("wat")
	 fmt.Println(e)
 }
