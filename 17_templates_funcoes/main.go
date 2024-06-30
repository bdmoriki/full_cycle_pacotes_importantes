package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"header.html", "content.html", "footer.html"}

	t := template.New("content.html")
	t.Funcs(map[string]interface{}{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Go", CargaHoraria: 40},
		{Nome: "Java", CargaHoraria: 80},
		{Nome: "Python", CargaHoraria: 120},
	})
	if err != nil {
		panic(err)
	}
}
