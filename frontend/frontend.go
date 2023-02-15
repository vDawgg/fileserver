package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	pb "main/proto"
	"main/helpers"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		Out string
	}

	clientAuthenticator := helpers.GetAuthenticatorClient()
	
	empty := &pb.Empty{}
	keys, err := clientAuthenticator.GetKeys(context.Background(), empty)
	if err != nil {
		fmt.Println("Failed to get keys: ", err)
		return
	}

	p := &Page{Out: keys.Keys}
	t, _ := template.ParseFiles("templates/index.html")
	err = t.Execute(w, p)
	if err != nil {
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	clientAuthenticator := helpers.GetAuthenticatorClient()

	empty := &pb.User{}
	token, err := clientAuthenticator.Login(context.Background(), empty)
	if err != nil {
		fmt.Println("Failed to get keys: ", err)
		return
	}

	fmt.Println(token)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", testHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
