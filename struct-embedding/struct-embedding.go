package main

import "fmt"

// Go supports embedding of structs and interfaces to express a more seamless composition of types

// This is implemented in contrast to the classical sense of inheritance

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base
// An embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {
	// When creating structs with literals, we have to initialise the embedding explicity
	// Here, the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base's fields directly on co
	// The base's fields in co are known as **promoted fields**
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively we can spell out the full path using the embedded type name
	fmt.Println("also num:", co.base.num)

	// Since container embeds base, the methods of base also become methods of a container
	// Here, we invoke a method that was embedded from base directly on co
	fmt.Println("describe:", co.describe())
	// Note that, despite describe() being called on a container, it is passed a base receiver
	// This differs from classical inheritances in languages like C++

	// Embedding structs wit methods may be used to bestow interface implementations onto other
	// structs
	// Here we see that a container now implements the describer interface because it embeds base
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	// If the embedding struct has a field x, and the embedded struct also has a field x, then the
	// embedded struct's field is **shadowed**, and when calling embedder.x, we get the embedding
	// struct's x field
}
