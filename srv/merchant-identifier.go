package mastercard

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	mastercardpb "github.com/akiliLab/mastercard/pb"
)

// SearchMethod is the type of search to be performed
type SearchMethod int64

// String returns a string that corresponds to the search method
func (sm *SearchMethod) String() string {
	if *sm == FuzzyMatch {
		return "FuzzyMatch"
	}
	return "ExactMatch"
}

// Search types that equals to the type of search that will be performed
const (
	FuzzyMatch = 0
	ExactMatch = 1
)

// GetMerchantIdentifiers returns all the merchants that matches the merchantId pattern
// Default searchMethod equals to ExactMatch which returns either 1 or 0 merchant while
// FuzzyMatch returns from 0 to 20 merchants with a Matching confidence field that scores
// from 0 to 100
func (c *Client) GetMerchantIdentifiers(merchantID string, search SearchMethod) (*mastercardpb.MerchantIDs, error) {
	urlFull := &url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   "/merchantid/v1/merchantid",
	}
	params := url.Values{}
	params.Set("Format", "XML")
	params.Set("MerchantId", merchantID)
	params.Set("Type", search.String())

	resp, err := c.oauthClient.Get(c.httpClient, &c.oauthClient.Credentials, urlFull.String(), params)
	if err != nil {
		log.Println("Error in mastercard merchantId: ", err.Error())
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println("Error in mastercard merchantId: status: ", resp.Status, "code: ", resp.StatusCode)
		return nil, errors.New("Failed request to mastercard api : " + resp.Status)
	}
	defer resp.Body.Close()

	merchantIds := &mastercardpb.MerchantIDs{}
	serializedBody, _ := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(serializedBody, merchantIds)
	if err != nil {
		return nil, errors.New("Error in mastercard merchant id: Couldn't parse XML response")
	}
	bodyString := string(serializedBody)
	fmt.Println("Body: ", bodyString)
	return merchantIds, nil
}
