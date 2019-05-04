package mastercard

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"

	mastercardpb "github.com/akiliLab/mastercard/proto"
)

// MerchantTransferFundingAndPayment secures the funds from a person and transfer it to the merchant
// It is one of the p2m api
func (c *Client) MerchantTransferFundingAndPayment(request *mastercardpb.MerchantTransferFundingAndPaymentRequest) (*mastercardpb.MerchantTransferFundingAndPaymentResponse, error) {

	urlFull := &url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		Path:   "/send/v1/partners/" + request.PartnerId + "/merchant/transfer",
	}

	params := structToMap(request)

	params.Set("Format", "XML")

	resp, err := c.oauthClient.Get(c.httpClient, &c.oauthClient.Credentials, urlFull.String(), params)

	if err != nil {
		log.Println("Error in mastercard merchantPayment: ", err.Error())
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Println("Error in mastercard merchantPayment: status: ", resp.Status, "code: ", resp.StatusCode)
		return nil, errors.New("Failed request to mastercard api : " + resp.Status)
	}

	defer resp.Body.Close()

	merchantPaymentResponse := &mastercardpb.MerchantTransferFundingAndPaymentResponse{}
	serializedBody, _ := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(serializedBody, merchantPaymentResponse)
	if err != nil {
		return nil, errors.New("Error in mastercard merchantPayment: Couldn't parse XML response")
	}
	bodyString := string(serializedBody)
	fmt.Println("Body: ", bodyString)
	return merchantPaymentResponse, nil
}

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		values.Set(typ.Field(i).Name, fmt.Sprint(iVal.Field(i)))
	}
	return
}
