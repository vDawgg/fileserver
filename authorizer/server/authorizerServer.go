package main

import (
	pb "authorizer/proto"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)

type authorizerServer struct {
	pb.UnimplementedAuthorizerServer
}

//TODO: Implement Authorize() -> Method for filling the db with authorization information

//TODO: Reformat code to make it more readable!
func (s *authorizerServer) IsAuthorized(ctx context.Context, request *pb.AuthRequest) (*pb.AuthReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	serverAddr := "localhost:50051" //Needs to be changed to environment variable
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		fmt.Println("Unable to connect to authenticator: ", err)
		//Should this always return an error -> the Authreply doesnt seem to be sent
		return &pb.AuthReply{IsAuthorized: false}, err
	}
	defer conn.Close()

	authenticator := pb.NewAuthenticatorClient(conn)
	public, err := authenticator.GetKeys(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Println("Call for GetKeys failed: ", err)
		//Should this always return an error -> the Authreply doesnt seem to be sent
		return &pb.AuthReply{IsAuthorized: false}, err
	}

	fmt.Println("********Key: ", public) //TODO: Delete this once done with debugging
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(public.Keys))

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	fmt.Println("Is token valid: ", token.Valid) //Token is not valid although it was created seconds ago!
	fmt.Println("Token Claims: ", token.Claims)  //How can the specific claims be indexed?

	authCollection := client.Database("authorization").Collection("authDict")

	err = authCollection.Drop(context.Background())
	var result bson.D

	err = authCollection.FindOne(context.Background(), bson.D{{"uid", request.Token}}).
		Decode(&result)
	if err != nil {
		fmt.Println("Unable to retrieve authorization information for user!")
		return &pb.AuthReply{IsAuthorized: false}, err
	}

	fmt.Println("Result: ", result)

	//TODO: Edit proto to find out what values to check for in the token

	reply := &pb.AuthReply{IsAuthorized: true}
	return reply, err
}

func main() {
	//Change this port?
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5555))
	if err != nil {
		//TODO: Change output to logs
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthorizerServer(server, &authorizerServer{})
	server.Serve(listen)
}
