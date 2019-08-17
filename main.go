package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var pool *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	type Invertory struct {
		Material string
		Count    uint
	}
	sweaters := Invertory{"wool", 17}
	tmpl, err := template.ParseFiles("templates/index.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, sweaters)
	if err != nil {
		panic(err)
	}
}

func main() {
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	flag.Parse()
	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	var err error
	pool, err = sql.Open("postgres", *dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
