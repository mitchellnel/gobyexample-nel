package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go offers built-in support for JSON encoding and decoding, including to and from built-in and
// custom data types

// We'll use these two structs to demonstrate encoding and decoding of custom types below
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON
// Fields must start with capital letters to be exported
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// first we'll look at encoding basic data types to JSON strings

	// here are some examples for atomic values
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// and here are some for slices and maps, which enode to JSON arrays and objects as you'd expect
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// the JSON package can automatically encode your custom data types
	// it will only include exported fields in the encoded output, and will by default use those
	// names as the JSON keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// you can use tags on struct field declarations to customise the encoded JSON key names
	// check the definition of response2 to see these tags
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// now let's look at decoding JSON data into Go values

	// here's an example for a generic data structure
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// we need to provide a variable where the JSON package can put the decoded data
	// this map[string]interface{} will hold a map of strings to arbitrary data types
	var dat map[string]interface{}

	// here's the actual decoding, and a check for associated errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// in order to use the values in the decoded map, we'll need to convert them to their
	// appropriate type
	// for example, here we convert the value in num to the expected float64 type
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of conversions
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// we can also decode JSON into custom data types
	// this has the advantages of adding additional type-safety to our programs and eliminates the
	// need for type assertions when accessing the decoded data
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// in the examples above we always used bytes and strings as intermediates between the data and
	// JSON representation on stdout
	// we can also stream JSON encodings directly to os.Writers or even HTTP response bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// https://go.dev/blog/json
	// https://pkg.go.dev/encoding/json
}
