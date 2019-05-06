package main

import (
	"context"
	"log"
	"net"
	"os"

	mastercardpb "github.com/akiliLab/mastercard/proto"
	mastercard "github.com/akiliLab/mastercard/srv"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var client *mastercard.Client

type server struct{}

func (s *server) GetMerchantIdentifiers(ctx context.Context, request *mastercardpb.MastercardRequest) (*mastercardpb.MastercardReply, error) {
	// MerchantId = "STILLWATERSGENERALSTBRANSONMO"
	// Search = FuzzyMatch
	response, err := client.GetMerchantIdentifiers(request.MerchantID, mastercard.FuzzyMatch)

	rep := &mastercardpb.MastercardReply{MerchantIDs: response}
	return rep, err
}

func (s *server) MerchantTransferFundingAndPayment(ctx context.Context, request *mastercardpb.MerchantTransferFundingAndPaymentRequest) (*mastercardpb.MerchantTransferFundingAndPaymentResponse, error) {

	response, err := client.MerchantTransferFundingAndPayment(request)

	return response, err
}

func (s *server) GetCurrencyConversion(ctx context.Context, request *mastercardpb.MastercardCurrencyConversionRequest) (*mastercardpb.MastercardCurrencyConversionReply, error) {

	response, err := client.GetCurrencyConversion(request)

	rep := &mastercardpb.MastercardCurrencyConversionReply{Response: response}

	return rep, err
}

func main() {

	log.Println("Service has started")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	keyStorePassword := os.Getenv("KEY_STORE_PASSWORD")
	credentialP2 := os.Getenv("P2_PATH")

	data, _ := mastercard.ExtractRSAPrivateKey(credentialP2, keyStorePassword)

	client, _ = mastercard.NewClient(consumerKey, data, keyStorePassword, mastercard.SANDBOX)

	lis, err := net.Listen("tcp", ":5005")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	mastercardpb.RegisterMastercardServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
