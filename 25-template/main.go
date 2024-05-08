package main

import (
	"os"
	"text/template"
)

var tpl = `
	{{ "\"Constant output\""}}
	{{ printf "%s" "Constant output" }}
	{{ printf "%s" "Length of" }}
	{{ len "Length of" }}
	{{if .Name}}
		Hello {{.Name}}!
	{{else}}
		Hello World!
	{{end}}
`

// {{range .array}}
//
//	{{.}}
//
// {{end}}
func main() {
	t := template.New("template")
	t1, e := t.Parse(tpl)
	if e != nil {
		println(e.Error())
		return
	}
	err := t1.Execute(os.Stdout, map[string]string{"Name": "John"})
	if err != nil {
		return
	}

	ranges()
}

type Person struct {
	Name string
	Age  int
}

var people = []Person{
	{"John", 20},
	{"Jane", 30},
	{"Jack", 40},
}

var tpl2 = `
{{range .}}
	{{.Name}} is {{.Age}} years old.
{{end}}
`

func ranges() {
	t := template.New("template")
	t1, e := t.Parse(tpl2)
	if e != nil {
		println(e.Error())
		return
	}
	err := t1.Execute(os.Stdout, people)
	if err != nil {
		return
	}
}
