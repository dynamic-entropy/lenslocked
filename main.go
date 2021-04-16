package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Name    string
	Surname string
	Domain  string
}

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := Data{Name: "Rahul"}
	if err := homeTemplate.Execute(w, data); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := Data{"Rahul", "Chauhan", "Golang"}
	if err := contactTemplate.Execute(w, data); err != nil {
		panic(err)
	}

}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.html")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
