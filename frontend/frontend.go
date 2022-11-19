package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"os"
)

type Page struct {
	Out string
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	out := "test"
	p := &Page{Out: out}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", testHandler);
	log.Fatal(http.ListenAndServe(":8080", nil));
}