package helpers

import (
	"fmt"
	pb "main/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetAuthenticatorClient() pb.AuthenticatorClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connAuthenticator, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		fmt.Println("Unable to connect to Authenticator: ", err)
	}

	clientAuthenticator := pb.NewAuthenticatorClient(connAuthenticator)

	return clientAuthenticator
}

func GetRetrieverClient() pb.RetrieverClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connRetriever, err := grpc.Dial("localhost:9390", opts...)
	if err != nil {
		fmt.Println("Unable to connect to Retriever: ", err)
	}

	retrieverClient := pb.NewRetrieverClient(connRetriever)

	return retrieverClient
}
