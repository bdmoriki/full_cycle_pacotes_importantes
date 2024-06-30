package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"header.html", "content.html", "footer.html"}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Go", CargaHoraria: 40},
		{Nome: "Java", CargaHoraria: 80},
		{Nome: "Python", CargaHoraria: 120},
	})
	if err != nil {
		panic(err)
	}
}
