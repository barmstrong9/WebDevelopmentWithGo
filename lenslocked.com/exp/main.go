package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}	

	data := struct {
		Name string
		Age int
	}{"Brandon Armstrong", 20}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}