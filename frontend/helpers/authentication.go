package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	pb "main/proto"
)

func EncryptUser(user *pb.User, keys *pb.Keys) *pb.User {
	pubKeyBlock, _ := pem.Decode([]byte(keys.Keys))
	hash := sha1.New()
	random := rand.Reader

	pubInterface, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Failed to parse key", err)
		return nil
	}
	pub := pubInterface.(*rsa.PublicKey)
	
	encryptedName, err := rsa.EncryptOAEP(hash, random, pub, user.Name, nil)
	if err != nil {
		fmt.Println("Failed to encrypt name")
	}
	encryptedPassword, err := rsa.EncryptOAEP(hash, random, pub, user.Password, nil)
	if err != nil {
		fmt.Println("Failed to encrypt password")
	}

	return &pb.User{Name: encryptedName, Password: encryptedPassword}
}	
