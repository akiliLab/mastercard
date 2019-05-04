package mastercard

import (
	"encoding/json"
	"fmt"
	"net/url"

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

	// params := url.Values{}

	res1B, _ := json.Marshal(request)

	fmt.Println(string(res1B))
	fmt.Println(urlFull)

	return nil, nil
}
