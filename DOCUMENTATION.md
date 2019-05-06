## getMerchantIdentifiers

### Models

```go

// CountrySubdivision contains few informations about the merchant's country
type CountrySubdivision struct {
	Name string
	Code string
}

// Country contains country informations
type Country struct {
	Name string
	Code string
}

// Address contains the full address of a merchant
type Address struct {
	Line1              string
	Line2              string
	City               string
	PostalCode         string
	CountrySubdivision CountrySubdivision
	Country            Country
}

// Merchant is a structure containing extra informations about a merchant
type Merchant struct {
	Address              Address
	PhoneNumber          string
	BrandName            string
	MerchantCategory     string
	MerchantDbaName      string
	DescriptorText       string
	LegalCorporateName   string
	Comment              string
	LocationID           int
	SoleProprietorName   string
	MatchConfidenceScore int
}

// MerchantIDs is the a structure containing from 0 to multiple MerchantId structures
// depending on how many results are returned by Mastercard API
type MerchantIDs struct {
	Message           string
	ReturnedMerchants []Merchant
}

```

> Test Data

```

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
		TransferReference:             "4117751524693558719224774292598091015092",
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
		AdditionalMessage:             "Lunch",
		InterchangeRateDesignator:     "QR",
		MastercardAssignedId:          "123456",
	}

	req := &mastercardpb.MerchantTransferFundingAndPaymentRequest{
		PartnerId:        "ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo",
		MerchantTransfer: merchantTransfer,
	}

	resp, _ := client.MerchantTransferFundingAndPayment(req)

	log.Println(resp)
```
