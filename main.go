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

	transferAmount := &mastercardpb.TransferAmount{
		Value:    "18",
		Currency: "USD",
	}

	address := &mastercardpb.Address{
		Line1:              "21 Broadway",
		Line2:              "Apartment A-6",
		City:               "OFallon",
		CountrySubdivision: "MO",
		PostalCode:         "63368",
		Country:            "USA",
	}

	additionalMerchantData := &mastercardpb.AdditionalMerchantData{
		PaymentFacilitatorId: "00000123456",
		SubMerchantId:        "A1234D",
	}

	sender := &mastercardpb.Sender{
		FirstName:              "John",
		MiddleName:             "Tyler",
		LastName:               "Jones",
		Address:                address,
		Phone:                  "11234565555",
		Email:                  "John.Jones123@abcmail.com",
		AdditionalMerchantData: additionalMerchantData,
	}

	customFields := []*mastercardpb.CustomField{{
		Name:  "ABC",
		Value: "456",
	}, {
		Name:  "DEF",
		Value: "789",
	}, {
		Name:  "GHI",
		Value: "123",
	}}

	reconciliationData := &mastercardpb.ReconciliationData{
		CustomField: customFields,
	}

	participant := &mastercardpb.Participant{
		CardAcceptorName: "WELLS FARGO BANK NA",
	}

	recipient := &mastercardpb.Recipient{
		FirstName:  "Jane",
		MiddleName: "Tyler",
		LastName:   "Smith",
		Address: &mastercardpb.Address{
			Line1:              "1 Main St",
			Line2:              "Apartment 9",
			City:               "OFallon",
			CountrySubdivision: "MO",
			PostalCode:         "63368",
			Country:            "USA",
		},
		Phone: "11234567890",
		Email: "Jane.Smith123@abcmail.com",
		AdditionalMerchantData: &mastercardpb.AdditionalMerchantData{
			PaymentFacilitatorId: "00000123456",
			SubMerchantId:        "A1234D",
		},
		MerchantCategoryCode: "3000",
	}

	merchantTransfer := &mastercardpb.MerchantTransfer{
		TransferReference:             "4007751524693558719224774292598091015092",
		PaymentType:                   "P2M",
		TransferAmount:                transferAmount,
		PaymentOriginationCountry:     "USA",
		SenderAccountUri:              "pan:5299920210000277;exp=2077-08;cvc=123",
		DigitalAccountReferenceNumber: "pan:5299920210000277",
		Sender:                        sender,
		RecipientAccountUri:           "pan:5299920210000277;exp=2077-08",
		Recipient:                     recipient,
		ReconciliationData:            reconciliationData,
		TransactionLocalDateTime:      "2016-09-22T13:22:11-05:30",
		Participant:                   participant,
		ParticipationId:               "TERMINAL34728",
		AdditionalMessage:             "Lunch @",
		InterchangeRateDesignator:     "QR",
		MastercardAssignedId:          "123456",
	}

	req := &mastercardpb.MerchantTransferFundingAndPaymentRequest{
		PartnerId:        "ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo",
		MerchantTransfer: merchantTransfer,
	}

	resp, _ := client.MerchantTransferFundingAndPayment(req)

	log.Println(resp)

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
