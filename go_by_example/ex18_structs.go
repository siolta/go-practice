package main

import "fmt"

// structs are typed collections of fields

type person struct {
	// name and age fields
	name string
	age int
}

func newPerson(name string) *person {
	
	p := person{name: name}
	p.age = 42

	// safe to return pointer to local var 
	// as a local var will survive the scope
	return &p
}

func main() {
	// syntax to create a new struct
	fmt.Println(person{"Bob", 20})

	// possible to name the fields during initialization
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields are zero-valued
	fmt.Println(person{name: "Fred"})

	// & prefix yields pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// idiomatic to encapsulate creation of structs in constructor functions
	fmt.Println(newPerson("Jon"))

	// access fields with dot syntax
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// can also use dots with pointers, pointers are auto dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
