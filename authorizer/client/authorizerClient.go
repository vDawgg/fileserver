package main

import (
	pb "authorizer/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func getAuthorization(clientAuthorizer pb.AuthorizerClient, token *pb.AuthRequest) {
	reply, err := clientAuthorizer.IsAuthorized(context.Background(), token)
	if err != nil {
		fmt.Println("*******Reply: ", reply)
		panic(err)
	}
	fmt.Println(reply)
}

func main() {
	//Change this to secure connectiongr
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connAuthorizer, err := grpc.Dial("localhost:5555", opts...)
	if err != nil {
		fmt.Println("Unable to connect to Authorizer: ", err)
	}
	defer connAuthorizer.Close()

	clientAuthorizer := pb.NewAuthorizerClient(connAuthorizer)

	connAuthenticator, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		fmt.Println("Unable to connect to Authenticator: ", err)
	}
	defer connAuthenticator.Close()

	clientAuthenticator := pb.NewAuthenticatorClient(connAuthenticator)

	name, err := os.ReadFile("client/nameEncrypted")
	if err != nil {
		fmt.Println("Unable to open nameEncrypted: ", err)
	}
	pass, err := os.ReadFile("client/passEncrypted")
	if err != nil {
		fmt.Println("Unable to open passEncrypted: ", err)
	}

	user := &pb.User{Name: name, Password: pass}
	token, err := clientAuthenticator.Login(context.Background(), user)

	t := &pb.AuthRequest{Token: token.Token}

	getAuthorization(clientAuthorizer, t)
}
