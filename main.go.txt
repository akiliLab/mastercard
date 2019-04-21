package main

import (
	"fmt"
	"io/ioutil"
	"mastercard"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("mastercard_consumer.key")
	if err != nil {
		fmt.Println("Couldn't read consumer.key file")
		os.Exit(1)
	}

	consumerKey := string(data)

	c, err := mastercard.NewClient(consumerKey, "credentials.p12", "keystorepassword", mastercard.SANDBOX)

	if err != nil {
		os.Exit(1)
	}

	_, err = c.GetMerchantIdentifiers("STILLWATERSGENERALSTBRANSONMO", mastercard.FuzzyMatch)
	if err != nil {
		os.Exit(1)
	}
}
