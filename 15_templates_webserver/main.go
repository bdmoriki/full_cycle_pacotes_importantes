package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{Nome: "Go", CargaHoraria: 40},
			{Nome: "Java", CargaHoraria: 80},
			{Nome: "Python", CargaHoraria: 120},
		})
		if err != nil {
			panic(err)
		}
	}))
}
