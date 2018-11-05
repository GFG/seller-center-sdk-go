package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"
	"time"
)

type ClientUrlBuilder interface {
	BuildUrl(requestParams url.Values) (string, error)
}

func NewClientUrlBuilder(clientConfig clientConfig) ClientUrlBuilder {
	return clientUrlBuilder{
		config: clientConfig,
		hashHmacRequestSignature: hashHmacRequestSignature{
			key: clientConfig.Key,
		},
		datetimeProvider: datetimeProvider{
			now: time.Now(),
		},
	}
}

type clientUrlBuilder struct {
	config                   clientConfig
	hashHmacRequestSignature HashHmacRequestSignature
	datetimeProvider         DatetimeProvider
}

func (urlBuilder clientUrlBuilder) BuildUrl(requestParams url.Values) (string, error) {
	currentUrl, err := urlBuilder.baseUrl()
	if err != nil {
		return "", err
	}

	requestParams.Add(fieldUserId, urlBuilder.config.User)
	requestParams.Add(fieldTimestamp, urlBuilder.datetimeProvider.getFormatted())

	signature := urlBuilder.createSignatureForRequest(requestParams)

	requestParams.Add(fieldSignature, signature)

	currentUrl.RawQuery = strings.Replace(requestParams.Encode(), "+", "%20", -1)

	return currentUrl.String(), nil
}

func (urlBuilder clientUrlBuilder) createSignatureForRequest(params url.Values) string {
	queryString := strings.Replace(params.Encode(), "+", "%20", -1)
	return urlBuilder.hashHmacRequestSignature.sign(queryString)
}

func (urlBuilder clientUrlBuilder) baseUrl() (*url.URL, error) {
	baseUrl, err := url.ParseRequestURI(urlBuilder.config.Url)
	if err != nil {
		return nil, err
	}

	return baseUrl, nil
}

type DatetimeProvider interface {
	getFormatted() string
}

type datetimeProvider struct {
	now time.Time
}

func (d datetimeProvider) getFormatted() string {
	return d.now.Format(time.RFC3339)
}

type HashHmacRequestSignature interface {
	sign(requestAsString string) string
}

type hashHmacRequestSignature struct {
	key string
}

func (hashH hashHmacRequestSignature) sign(requestAsString string) string {
	key := []byte(hashH.key)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(requestAsString))
	return hex.EncodeToString(h.Sum(nil))
}
