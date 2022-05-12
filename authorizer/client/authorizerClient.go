package main

import (
	pb "authorizer/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func addAuthorization(clientAuthorizer pb.AuthorizerClient, ar *pb.AuthRequest) {
	reply, err := clientAuthorizer.AddAuthorization(context.Background(), ar)
	if err != nil {
		fmt.Println("Unable to add authorization:", err)
		return
	}
	fmt.Println("Reply: ", reply)
}

func getAuthorization(clientAuthorizer pb.AuthorizerClient, ar *pb.AuthRequest) {
	reply, err := clientAuthorizer.IsAuthorized(context.Background(), ar)
	if err != nil {
		fmt.Println("Unable to get authorization: ", err)
		return
	}
	fmt.Println("Reply: ", reply)
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
	//Authenticator needs to be run from within the same directory, otherwise the connection does not work
	token, err := clientAuthenticator.Login(context.Background(), user)
	if err != nil {
		fmt.Println("Unable to log in: ", err)
	}

	ar := &pb.AuthRequest{Token: token.Token, Access: []string{"home", "home/pictures", "home/videos"}}
	arUpdate := &pb.AuthRequest{Token: token.Token, Access: []string{"home/private", "home/school", "home/school/prog"}}
	arAll := &pb.AuthRequest{Token: token.Token, Access: []string{"home", "home/pictures", "home/videos", "home/private", "home/school", "home/school/prog"}}

	arNoAccess := &pb.AuthRequest{Token: token.Token, Access: []string{"home/private/super_private"}}
	arNoToken := &pb.AuthRequest{Token: "", Access: []string{"home/private"}}

	fmt.Println("Adding ar")
	addAuthorization(clientAuthorizer, ar)
	fmt.Println("Authorizing ar")
	getAuthorization(clientAuthorizer, ar)

	fmt.Println("Adding arUpdate")
	addAuthorization(clientAuthorizer, arUpdate)
	fmt.Println("Authorizing arAll")
	getAuthorization(clientAuthorizer, arAll)

	fmt.Println("Authorizing arNoAccess")
	getAuthorization(clientAuthorizer, arNoAccess)
	fmt.Println("Authorizing arNoTokens")
	getAuthorization(clientAuthorizer, arNoToken)
}
