package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"main/helpers"
	pb "main/proto"
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

	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	clientAuthenticator := helpers.GetAuthenticatorClient()
	
	empty := &pb.Empty{}
	keys, err := clientAuthenticator.GetKeys(context.Background(), empty)
	if err != nil {
		fmt.Println("Failed to get keys: ", err)
		return
	}
	fmt.Println(keys.Keys)

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	//TODO: check this input for injections etc.!
	user := &pb.User{Name: []byte(r.FormValue("name")), Password: []byte(r.FormValue("password"))}

	token, err := clientAuthenticator.Login(context.Background(), user)
	if err != nil {
		fmt.Println("Failed to login: ", err)
		tmpl.Execute(w, nil)
		return
	}

	fmt.Println(token.Token)

	tmpl.Execute(w, struct{ Success bool }{true})

	fmt.Println(token)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", testHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
