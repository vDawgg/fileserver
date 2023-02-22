package handlers

import (
	"context"
	"fmt"
	"html/template"
	"main/helpers"
	pb "main/proto"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	clientAuthenticator := helpers.GetAuthenticatorClient()
	
	empty := &pb.Empty{}
	keys, err := clientAuthenticator.GetKeys(context.Background(), empty)
	if err != nil {
		fmt.Println("Failed to get keys: ", err)
		return
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	//TODO: check this input for injections etc.!
	user := &pb.User{Name: []byte(r.FormValue("name")), Password: []byte(r.FormValue("password"))}
	user = helpers.EncryptUser(user, keys)

	token, err := clientAuthenticator.Login(context.Background(), user)
	if err != nil || token.Status == pb.AuthStatus_FAILED {
		fmt.Println("Failed to login: ", err)
		tmpl.Execute(w, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "jwt-token", Value: token.Token})

	tmpl.Execute(w, struct{ Success bool }{true})
	fmt.Println(token)
}