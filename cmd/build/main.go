package main

import (
	_ "embed"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed main.tpl
var mainTmpl string

func main() {
	file, err := os.Create(filepath.Join("dist", "main.go"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t := template.Must(template.New("main").Parse(mainTmpl))
	err = t.Execute(file, nil)
	if err != nil {
		panic(err)
	}
}
