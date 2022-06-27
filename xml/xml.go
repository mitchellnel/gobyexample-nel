package main

import (
	"encoding/xml"
	"fmt"
)

// Go offers built-in support for XML and XML-like formats with the encoding/xml package

// Plant will be mapepd to XML
// Similarly to the JSON examples, field tags contain directives for the encoder and decoder
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

// Here, we use some special features of the XML package: the XMLName field name dictates the name
// of the XML element representing this struct; id,attr means that the Id field is an XML attribute
// rather than a nested element

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// emit XML representing our plant
	// using MarshalIndent to produce a more human-readable output
	out, _ := xml.MarshalIndent(coffee, " ", "\t")
	fmt.Println(string(out))

	// to add a generic XML header to the output, append it explicitly
	fmt.Println(xml.Header + string(out))

	// use Unmarshal to parse a string of bytes with XML into a data structure
	// if the XML is malformed or cannot be mapped onto plant, a descriptive error will be returned
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "United States of America"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	// the parent>child>plant tag tells the encode to nest all plants under <parent><child>

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "\t")
	fmt.Println(string(out))
}
