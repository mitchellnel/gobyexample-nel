package main

import "fmt"

// Go's structs are typed collections of fields
// They're useful for grouping data together to form records

// This person struct type has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// We can safely return a pointer to a local variable, as a local variable will survive the
	// scope of the function
	return &p
}

func main() {
	// Creating a new struct
	bob := person{"Bob", 20}
	fmt.Println(bob)

	// We can specifically name the fields when initialising a struct
	alice := person{name: "Alice", age: 30}
	fmt.Println(alice)

	// Omitted fields will be zero-valued
	fred := person{name: "Fred"}
	fmt.Println(fred)

	// Pointer to a struct
	annPtr := &person{name: "Ann", age: 40}
	fmt.Println(annPtr)

	// It's idomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// We can also use dots with struct pointers -- the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
