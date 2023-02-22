package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/home", handlers.FileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
