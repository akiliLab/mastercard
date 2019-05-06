package mastercard

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	OAUTH_VERSION          = "1.0"
	SIGNATURE_METHOD       = "RSA-SHA256"
	OAUTH_HEADER           = "OAuth "
	BODY_HASH_PARAM        = "oauth_body_hash"
	CONSUMER_KEY_PARAM     = "oauth_consumer_key"
	NONCE_PARAM            = "oauth_nonce"
	SIGNATURE_METHOD_PARAM = "oauth_signature_method"
	SIGNATURE_PARAM        = "oauth_signature"
	TIMESTAMP_PARAM        = "oauth_timestamp"
	TOKEN_PARAM            = "oauth_token"
	VERSION_PARAM          = "oauth_version"
)

var oauthKeys = []string{
	BODY_HASH_PARAM,
	CONSUMER_KEY_PARAM,
	NONCE_PARAM,
	SIGNATURE_PARAM,
	SIGNATURE_METHOD_PARAM,
	TIMESTAMP_PARAM,
	VERSION_PARAM,
}

var nonceCounter uint64

func getAuthorizationHeader(url *url.URL, params url.Values, method string, payload []byte, consumerKey string, privateKey *rsa.PrivateKey) string {

	oauthParams, _ := getOauthParams(consumerKey, payload)

	paramString := toOAuthParamString(params, oauthParams)

	h := sha256.New()

	writeSignatureBaseString(h, method, url, paramString)

	rawSignature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return ""
	}

	signature := base64.StdEncoding.EncodeToString(rawSignature)

	oauthParams[SIGNATURE_PARAM] = signature

	return getAuthorizationString(oauthParams)
}

//  Generates a hash based on request payload as per
//  https://tools.ietf.org/id/draft-eaton-oauth-bodyhash-00.html

func getBodyHash(payload []byte) string {
	hasher := sha256.New()
	hasher.Write(payload)
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}

// Generates a random string for replay protection as per
//  https://tools.ietf.org/html/rfc5849#section-3.3

func getNonce() string {
	return strconv.FormatUint(atomic.AddUint64(&nonceCounter, 1), 16)
}

// @param {String} consumerKey Consumer key set up in a Mastercard Developer Portal project
// @param {[]byte} payload Payload (nullable)
// @return {map[string]string}
func getOauthParams(consumerKey string, payload []byte) (map[string]string, error) {

	oauthParams := map[string]string{
		BODY_HASH_PARAM:        getBodyHash(payload),
		CONSUMER_KEY_PARAM:     consumerKey,
		NONCE_PARAM:            getNonce(),
		SIGNATURE_METHOD_PARAM: SIGNATURE_METHOD,
		TIMESTAMP_PARAM:        getTimeStamp(),
		VERSION_PARAM:          OAUTH_VERSION,
	}

	return oauthParams, nil

}

func getTimeStamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

type keyValue struct{ key, value []byte }

type byKeyValue []keyValue

func (p byKeyValue) Len() int      { return len(p) }
func (p byKeyValue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byKeyValue) Less(i, j int) bool {
	sgn := bytes.Compare(p[i].key, p[j].key)
	if sgn == 0 {
		sgn = bytes.Compare(p[i].value, p[j].value)
	}
	return sgn < 0
}

func (p byKeyValue) appendValues(values url.Values) byKeyValue {
	for k, vs := range values {
		k := encode(k, true)
		for _, v := range vs {
			v := encode(v, true)
			p = append(p, keyValue{k, v})
		}
	}
	return p
}

// noscape[b] is true if b should not be escaped per section 3.6 of the RFC.
var noEscape = [256]bool{
	'A': true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	'a': true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	'0': true, true, true, true, true, true, true, true, true, true,
	'-': true,
	'.': true,
	'_': true,
	'~': true,
}

// encode encodes string per section 3.6 of the RFC. If double is true, then
// the encoding is applied twice.
func encode(s string, double bool) []byte {
	// Compute size of result.
	m := 3
	if double {
		m = 5
	}
	n := 0
	for i := 0; i < len(s); i++ {
		if noEscape[s[i]] {
			n++
		} else {
			n += m
		}
	}

	p := make([]byte, n)

	// Encode it.
	j := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if noEscape[b] {
			p[j] = b
			j++
		} else if double {
			p[j] = '%'
			p[j+1] = '2'
			p[j+2] = '5'
			p[j+3] = "0123456789ABCDEF"[b>>4]
			p[j+4] = "0123456789ABCDEF"[b&15]
			j += 5
		} else {
			p[j] = '%'
			p[j+1] = "0123456789ABCDEF"[b>>4]
			p[j+2] = "0123456789ABCDEF"[b&15]
			j += 3
		}
	}
	return p
}

/**
 * Lexicographically sort all parameters and concatenate them into a string as per
 * https://tools.ietf.org/html/rfc5849#section-3.4.1.3.2
 */

func toOAuthParamString(queryParams url.Values, oauthParams map[string]string) byKeyValue {

	p := make(byKeyValue, 0, len(queryParams)+len(oauthParams))
	p = p.appendValues(queryParams)

	for k, v := range oauthParams {
		p = append(p, keyValue{encode(k, true), encode(v, true)})
	}
	sort.Sort(p)

	return p
}

func writeSignatureBaseString(w io.Writer, httpMethod string, u *url.URL, queryParams byKeyValue) {

	w.Write(encode(strings.ToUpper(httpMethod), false))
	w.Write([]byte{'&'})

	scheme := strings.ToLower(u.Scheme)
	host := strings.ToLower(u.Host)

	uNoQuery := *u
	uNoQuery.RawQuery = ""
	path := uNoQuery.RequestURI()

	switch {
	case scheme == "http" && strings.HasSuffix(host, ":80"):
		host = host[:len(host)-len(":80")]
	case scheme == "https" && strings.HasSuffix(host, ":443"):
		host = host[:len(host)-len(":443")]
	}

	w.Write(encode(scheme, false))
	w.Write(encode("://", false))
	w.Write(encode(host, false))
	w.Write(encode(path, false))
	w.Write([]byte{'&'})

	// Write the parameters.
	encodedAmp := encode("&", false)
	encodedEqual := encode("=", false)
	sep := false
	for _, kv := range queryParams {
		if sep {
			w.Write(encodedAmp)
		} else {
			sep = true
		}
		w.Write(kv.key)
		w.Write(encodedEqual)
		w.Write(kv.value)
	}
}

func getAuthorizationString(p map[string]string) string {
	var h []byte
	// Append parameters in a fixed order to support testing.
	for _, k := range oauthKeys {
		if v, ok := p[k]; ok {
			if h == nil {
				h = []byte(OAUTH_HEADER)
			} else {
				h = append(h, ","...)
			}
			h = append(h, k...)
			h = append(h, `="`...)
			h = append(h, encode(v, false)...)
			h = append(h, '"')
		}
	}
	return string(h)
}
