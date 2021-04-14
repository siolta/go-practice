// go has built-in support for JSON

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// these two types will be used to demonstrate encoding and decoding
type response1 struct {
	Page int
	Fruits []string
}

// only exported fields will be encoded, fields must start with capital letters
type response2 struct {
	Page 	int	 		`json:"page"`
	Fruits	[]string 	`json:"fruits"`
}

func main() {
	// first we encode basic data types to JSON strings
	// here are atomic value examples
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// here are examples for slices and maps, which encode to JSON arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map [string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// the JSON package can auto encode custom data types
	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// tags can be used on struct field declarations to customize the JSON key names
	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// generic data structure example
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// need a var where JSON can put the decoded data
	// this will hold a map of strings to arbitrary data types
	var dat map[string]interface{}

	// actual decoding, and a check for errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// to use the values we need to convert them to their type
	// here we convert the val in num to float64
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires multiple conversions
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// can decode JSON into custom data types
	// this adds additional type-safety to our programs
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// can also stream json encodings directly to os.Writers like os.Stdout
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
