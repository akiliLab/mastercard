package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

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

	data, err := ioutil.ReadFile("mastercard_consumer.key")
	if err != nil {
		fmt.Println("Couldn't read consumer.key file")
		os.Exit(1)
	}

	consumerKey := string(data)

	client, errs = NewClient(consumerKey, "credentials.p12", "keystorepassword", SANDBOX)

	if errs != nil {
		os.Exit(1)
	}
}
