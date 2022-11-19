package main

import (
	pb "main/proto"
	"fmt"
	"context"
	"html/template"
	"log"
	"net/http"
	//"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testHandler (w http.ResponseWriter, r *http.Request) {
	type Page struct {
		Out string
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connAuthenticator, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		fmt.Println("Unable to connect to Authenticator: ", err)
	}
	defer connAuthenticator.Close()

	clientAuthenticator := pb.NewAuthenticatorClient(connAuthenticator)

	empty := &pb.Empty{}
	keys, err := clientAuthenticator.GetKeys(context.Background(), empty)
	if err != nil {
		fmt.Println("Failed to get keys: ", err)
		return
	}

	p := &Page{Out: keys.Keys}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", testHandler);
	log.Fatal(http.ListenAndServe(":8080", nil));
}