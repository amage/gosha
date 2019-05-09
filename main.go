package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	type Invertory struct {
		Material string
		Count    uint
	}
	sweaters := Invertory{"wool", 17}
	tmpl, err := template.ParseFiles("index.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, sweaters)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
