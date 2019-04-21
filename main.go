package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	mastercard "github.com/ubunifupay/mastercard/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var client *Client

type server struct{}

func (s *server) GetMerchantIdentifiers(ctx context.Context, request *mastercard.MastercardRequest) (*mastercard.MastercardReply, error) {
	// MerchantId = "STILLWATERSGENERALSTBRANSONMO"
	// Search = FuzzyMatch
	response, err := client.GetMerchantIdentifiers(request.MerchantID, FuzzyMatch)

	rep := &mastercard.MastercardReply{MerchantIDs: response}
	return rep, err
}

func main() {
	var errs error

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lis, err := net.Listen("tcp", ":5005")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	mastercard.RegisterMastercardServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	keyStorePassword := os.Getenv("KEY_STORE_PASSWORD")
	credentialP2 := os.Getenv("P2_PATH")

	client, errs = NewClient(consumerKey, credentialP2, keyStorePassword, SANDBOX)

	if errs != nil {
		os.Exit(1)
	}
}
