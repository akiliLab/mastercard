package mastercard

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"

	mastercardpb "github.com/akiliLab/mastercard/proto"
)

// GetCurrencyConversion returns all the merchants that matches the merchantId pattern
// Default searchMethod equals to ExactMatch which returns either 1 or 0 merchant while
// FuzzyMatch returns from 0 to 20 merchants with a Matching confidence field that scores
// from 0 to 100
func (c *Client) GetCurrencyConversion(request *mastercardpb.MastercardCurrencyConversionRequest) (*mastercardpb.CurrencyResponse, error) {
	urlFull := &url.URL{
		Scheme: "https",
		Host:   c.BaseURL,
		//  Not very sure about this though
		Path: "/apiexplorer/mcapi/settlement/currencyrate/conversion-rate",
	}
	params := url.Values{}
	params.Set("Format", "XML")
	params.Set("fxDate", request.FxDate)
	params.Set("transCurr", request.TransCurr)
	params.Set("crdhldBillCurr", request.CrdhldBillCurr)
	params.Set("bankFee", strconv.FormatInt(request.BankFee, 10))
	params.Set("transAmt", fmt.Sprintf("%f", request.TransAmt))

	resp, err := c.oauthClient.Get(c.httpClient, &c.oauthClient.Credentials, urlFull.String(), params)
	if err != nil {
		log.Println("Error in mastercard currencyconversion: ", err.Error())
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println("Error in mastercard currencyconversion: status: ", resp.Status, "code: ", resp.StatusCode)
		return nil, errors.New("Failed request to mastercard api : " + resp.Status)
	}
	defer resp.Body.Close()

	dataConversionResponse := &mastercardpb.CurrencyResponse{}
	serializedBody, _ := ioutil.ReadAll(resp.Body)
	err = xml.Unmarshal(serializedBody, dataConversionResponse)
	if err != nil {
		return nil, errors.New("Error in mastercard currencyconversion: Couldn't parse XML response")
	}
	bodyString := string(serializedBody)
	fmt.Println("Body: ", bodyString)
	return dataConversionResponse, nil
}
