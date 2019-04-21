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
