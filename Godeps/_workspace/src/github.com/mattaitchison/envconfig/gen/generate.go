package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"
)

var genTemplate = `
// generated code -- DO NOT EDIT

package envconfig

import "net"

{{range .}}
// {{.Upper}} retrieves a environment variable by name and parses it to a {{.Lower}}
// defaultVal will be returned if the variable is not found.
func {{.Upper}}(name string, defaultVal {{.Lower}}, description string) {{.Lower}} {
	v := NewVar(new{{.Upper}}Value(defaultVal), name, description)
	return v.Value.Get().({{.Lower}})
}
{{if .Options}}
// {{.Upper}}Option like {{.Upper}} except the value must be in options.
// defaultVal will be returned if the variable is not found or is not a valid option.
func {{.Upper}}Option(name string, defaultVal {{.Lower}}, options []{{.Lower}}, description string) {{.Lower}} {
	v := NewVar(new{{.Upper}}Value(defaultVal), name, description)
	for _, option := range options {
		if option == v.Value.Get().({{.Lower}}) {
			return v.Value.Get().({{.Lower}})
		}
	}
	return defaultVal
}
{{end}}
{{end}}`

// // {{.Upper}}s is like {{.Upper}} but allows a list of {{.Lower}}.
// // defaultVal will be returned if the variable is not found.
// func {{.Upper}}s(name string, defaultVal []{{.Lower}}, separator string, description string) {{.Lower}} {
// 	v := NewVar(new{{.Upper}}Value(defaultVal), name, description)
// 	return v.Value.Get().({{.Lower}})
// }

func main() {

	// Prepare some data to insert into the template.
	type Type struct {
		Upper, Lower string
		Options      bool
	}
	var types = []Type{
		{"String", "string", true},
		{"Bool", "bool", false},
		{"Float64", "float64", true},
		{"Int", "int", true},
		{"Int64", "int64", true},
		{"Uint", "uint", true},
		{"Uint64", "uint64", true},
		{"IP", "net.IP", false},
	}

	tmpl := template.Must(template.New("envconfig").Parse(genTemplate))

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, types)
	if err != nil {
		log.Println("executing template:", err)
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
	}

	err = ioutil.WriteFile("helpers.go", src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

}
