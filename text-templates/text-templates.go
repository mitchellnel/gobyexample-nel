package main

import (
	"os"
	"text/template"
)

// Go offers built-in support for creating dynamic content or showing customised output to the user
// with the text/template package

// A sibling package named html/tempalte provides the same API but has additional security features
// and should be used for generating HTML

func main() {
	// we can create a new template and parse its body from a string
	// templates are a mix of static text and "actions" enclosed in {{...}} that are used to
	// dynamically insert content
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// alternatively, we can use the template.Must function to panic in case Parse returns and error
	// this is especially useful for tempaltes initalised in the global scope
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// by "executing" the template, we generate its text with specific values for its actions
	// the {{.}} action is replaced bt the value passed as a parameter to Execute
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
		"Python",
	})

	// helper function we'll use below
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// if the data is a struct we can use the {{.FieldName}} action to access its fields
	// the fields should be xported to be accessible when a template is executing
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"George Russell"})

	// the same applied to maps
	// with maps there is no restriction on the case of key names
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Lewis Hamilton",
	})

	// if/else provides conditional execution for templates
	// a value is condiered false if it's the default value of type, such as 0 or the empty string
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	// this sample demosntrates antoher feature of templates: using - in actions to trim whitespace

	// range blocks let us loop through slices, arrays, maps or channels
	// inside the range block, {{.}} is set to the current item of the iteration
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
		"Python",
	})
}
