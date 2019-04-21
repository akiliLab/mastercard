package mastercard

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gomodule/oauth1/oauth"
	"golang.org/x/crypto/pkcs12"
)

const (
	productionBaseURL = "api.mastercard.com/"
	stagingBaseURL    = "stage.api.mastercard.com"
	devBaseURL        = "dev.api.mastercard.com"
	sandboxBaseURL    = "sandbox.api.mastercard.com"
	userAgent         = "go-mastercard"
)

// A Client manages communication with the Mastercard API.
type Client struct {
	httpClient  *http.Client  // HTTP client used to communicate with the API.
	oauthClient *oauth.Client // Oauth 1 client to set authentication up

	// Base URL for API requests. Defaults to the public Mastercard API.
	BaseURL string

	// User agent used when communicating with the Mastercard API.
	UserAgent string
}

// EnvType equals the type of environment to work in
type EnvType int

// Environment types
const (
	PRODUCTION = 0
	STAGING    = 1
	DEV        = 2
	SANDBOX    = 3
)

// BaseURL Returns a string that corresponds to the environment type
func (e EnvType) BaseURL() string {
	switch e {
	case PRODUCTION:
		return productionBaseURL
	case STAGING:
		return stagingBaseURL
	case DEV:
		return devBaseURL
	case SANDBOX:
		return sandboxBaseURL
	}
	return sandboxBaseURL
}

// NewClient returns an instance of a Mastercard api client that allows
// accessing endpoints with more ease. Oauth1 authentification is managed
// internally, but you need to pass your mastercard consumerKey, the path
// to the .p12 file and the keystore password for the client to retrieve
// your RSA private key and sign requests correctly
func NewClient(consumerKey string, keystorePath string, keystorePassword string, env EnvType) (*Client, error) {
	client := &Client{
		httpClient:  http.DefaultClient,
		BaseURL:     env.BaseURL(),
		UserAgent:   userAgent,
		oauthClient: &oauth.Client{},
	}

	privateKey, err := extractRSAPrivateKey(keystorePath, keystorePassword)
	if err != nil {
		return nil, err
	}

	client.oauthClient.PrivateKey = privateKey
	client.oauthClient.SignatureMethod = oauth.RSASHA1
	client.oauthClient.Credentials = oauth.Credentials{
		Token: consumerKey,
	}
	return client, nil
}

func extractRSAPrivateKey(keystorePath string, keystorePassword string) (*rsa.PrivateKey, error) {
	p12data, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		log.Println("Error in mastercard client initialization: ", err.Error())
		return nil, err
	}

	data, _, err := pkcs12.Decode(p12data, keystorePassword)
	if err != nil {
		log.Println("Error in mastercard client initialization: ", err.Error())
		return nil, err
	}

	privateKey := data.(*rsa.PrivateKey)
	return privateKey, nil
}
