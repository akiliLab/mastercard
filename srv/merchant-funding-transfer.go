package mastercard

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	mastercardpb "github.com/akiliLab/mastercard/proto"
)

// MerchantTransferFundingAndPayment secures the funds from a person and transfer it to the merchant
// It is one of the p2m api
func (c *Client) MerchantTransferFundingAndPayment(request *mastercardpb.MerchantTransferFundingAndPaymentRequest) (*mastercardpb.MerchantTransferFundingAndPaymentResponse, error) {

	urlFull := &url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   "/send/static/v1/partners/" + request.PartnerId + "/merchant/transfer",
	}

	params := url.Values{}
	params.Add("Format", "JSON")

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	authorization := getAuthorizationHeader(urlFull, params, "POST", body, c.ConsumerKey, c.PrivateKey)

	urlFull.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", urlFull.String(), bytes.NewBuffer(body))
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Println("Error in mastercard merchantPayment: ", err.Error())
		return nil, err
	}

	if resp.StatusCode != 200 {
		serializedBody, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(serializedBody)
		fmt.Println("Body: ", bodyString)
		log.Println("Error in mastercard merchantPayment: status: ", resp.Status, "code: ", resp.StatusCode)
		return nil, errors.New("Failed request to mastercard api : " + resp.Status)
	}

	defer resp.Body.Close()

	merchantPaymentResponse := &mastercardpb.MerchantTransferFundingAndPaymentResponse{}

	serializedBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(serializedBody, merchantPaymentResponse)

	if err != nil {
		return nil, errors.New("Error in mastercard merchantPayment: Couldn't parse XML response")
	}

	return merchantPaymentResponse, nil

}
