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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

//TODO: Look for proper logging

type authorizerServer struct {
	pb.UnimplementedAuthorizerServer
}

type Doc struct {
	Id     string
	Access []string
}

func (s *authorizerServer) AddAuthorization(ctx context.Context, request *pb.AuthRequest) (*pb.Added, error) {
	iV, uid, err := isValid(request)

	fmt.Println("********Adding authorization info")
	fmt.Println("********User id: ", uid)

	if !iV {
		return &pb.Added{WasAdded: false}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Unable to connect to mongodb: ", err)
	}
	defer client.Disconnect(ctx)

	authCollection := client.Database("authorization").Collection("authDict")

	_, err = authCollection.InsertOne(context.Background(), bson.D{
		{"_id", uid},
		{"access", request.Access},
	})
	if err != nil {
		fmt.Println("Unable to insert document: ", err)

		var access Doc
		err = authCollection.FindOne(ctx, bson.D{{"_id", uid}}).Decode(&access)
		if err != nil {
			fmt.Println("Unable to retrieve access old access info: ", err)
			return &pb.Added{WasAdded: false}, err
		}

		_, err = authCollection.UpdateOne(ctx,
			bson.D{{"_id", uid}},
			bson.D{{"$set", bson.D{{"access", append(access.Access, request.Access...)}}}})
		if err != nil {
			fmt.Println("Unable to update document: ", err)
			return &pb.Added{WasAdded: false}, err
		}
	}

	return &pb.Added{WasAdded: true}, err
}

//TODO: Reformat code to make it more readable!
func (s *authorizerServer) IsAuthorized(ctx context.Context, request *pb.AuthRequest) (*pb.AuthReply, error) {
	iV, uid, err := isValid(request)
	fmt.Println("Received authorization request for: ", request.Access)

	if !iV {
		return &pb.AuthReply{IsAuthorized: false}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	authCollection := client.Database("authorization").Collection("authDict")

	cursor, err := authCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Unable to retrieve any documents: ", err)
	}
	var authInfo []bson.M
	if err = cursor.All(ctx, &authInfo); err != nil {
		fmt.Println("Cannot decode cursor: ", err)
	}
	//fmt.Println("Info currently in db: ", authInfo)

	var result Doc
	fmt.Println("Trying to retrieve info with uid: ", uid)
	err = authCollection.FindOne(context.Background(), bson.D{{"_id", uid}}).Decode(&result)
	if err != nil {
		fmt.Println("Unable to retrieve authorization information for user!", err)
		return &pb.AuthReply{IsAuthorized: false}, err
	}
	fmt.Println("++++++++Access contains: ", result.Access)

	//TODO: Verify whether or not the user actually has access after retrieving info on him

	if accessMatches(request.Access, result.Access) {
		fmt.Println("Granted access!")
		return &pb.AuthReply{IsAuthorized: true}, err
	}
	fmt.Println("User does not have access for: ", request.Access)

	//False flag for IsAuthorized returns empty reply -> Why???
	return &pb.AuthReply{IsAuthorized: false}, status.Error(codes.Internal, "User is not authorized")
}

func isValid(request *pb.AuthRequest) (bool, string, error) {
	serverAddr := "localhost:50051" //Needs to be changed to environment variable
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		fmt.Println("Unable to connect to authenticator: ", err)
		//Should this always return an error -> the Authreply doesnt seem to be sent
		return false, "", err
	}
	defer conn.Close()
	authenticator := pb.NewAuthenticatorClient(conn)
	public, err := authenticator.GetKeys(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Println("Call for GetKeys failed: ", err)
		return false, "", err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(public.Keys))
	if err != nil {
		fmt.Println("Unable to parse public key: ", err)
		return false, "", err
	}

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		fmt.Println("Could not parse the token: ", err)
		return false, "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	user := claims["user"].(map[string]interface{})
	fmt.Println("******Uid taken from claims: ", user["id"].(string))

	fmt.Println("Is token valid: ", token.Valid) //Token is not valid, although it was created seconds ago!

	if token.Valid {
		return true, user["id"].(string), err
	}
	fmt.Println("Could not validate token")
	return false, "", err
}

func contains(a []string, b string) bool {
	for _, c := range a {
		if c == b {
			return true
		}
	}
	return false
}

func accessMatches(request []string, db []string) bool {
	for _, s := range request {
		if !contains(db, s) {
			return false
		}
	}
	return true
}

func main() {
	//Change this port?
	fmt.Println("Starting authorizer")
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5555))
	if err != nil {
		//TODO: Change output to logs
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthorizerServer(server, &authorizerServer{})
	server.Serve(listen)
}
