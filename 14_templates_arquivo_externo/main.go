package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Go", CargaHoraria: 40},
		{Nome: "Java", CargaHoraria: 80},
		{Nome: "Python", CargaHoraria: 120},
	})

	if err != nil {
		panic(err)
	}
}
