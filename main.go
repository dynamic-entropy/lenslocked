package main

import (
	"lenslocked/views"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Name    string
	Surname string
	Domain  string
}

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := Data{Name: "Rahul"}
	if err := homeView.Template.Execute(w, data); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := Data{"Rahul", "Chauhan", "Golang"}
	if err := contactView.Template.Execute(w, data); err != nil {
		panic(err)
	}

}

func main() {
	homeView = views.NewView("views/home.html")
	contactView = views.NewView("views/contact.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
