package client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"net/mail"
	"net/url"
	"time"
)

const fieldUserId = "UserID"
const fieldSignature = "Signature"
const fieldTimestamp = "Timestamp"
const fieldAction = "Action"
const fieldVersion = "Version"
const fieldFormat = "Format"

// can be JSON` or `XML` currently only ResponseBuilder for JSON implemented
const defaultResponseFormat = "JSON"

const MethodGET = "GET"
const MethodPOST = "POST"

const V1 = "1.0"
const V2 = "2.0"

const timeoutInSeconds = 10

const maxRetries = 5

// Errors
var (
	NotSupportedMethod = errors.New("unsupported method")
)

type Client interface {
	Call(request Request) (Response, error)
	GetLogger() *log.Logger
}

type Request interface {
	GetMethod() string
	GetRequestParams() url.Values
	GeneratePostXml() ([]byte, error)
	SetVersion(v string)
	SetRequestParam(key string, value string)
	SetPostData(data interface{})
}

type postEnvelope struct {
	XMLName xml.Name `xml:"Request"`
	Data    interface{}
}

func NewGenericRequest(action string, method string) *genericRequest {
	return &genericRequest{
		params:  url.Values{},
		action:  action,
		method:  method,
		version: V1,
	}
}

type genericRequest struct {
	params   url.Values
	action   string
	method   string
	version  string
	postData interface{}
}

func (gr genericRequest) GetMethod() string {
	return gr.method
}

func (gr genericRequest) GetRequestParams() url.Values {
	params := url.Values{}
	params.Add(fieldAction, gr.action)
	params.Add(fieldVersion, gr.version)
	params.Add(fieldFormat, defaultResponseFormat)

	for k, v := range gr.params {
		params.Add(k, v[0])
	}

	return params
}

func (gr genericRequest) GeneratePostXml() ([]byte, error) {

	if gr.postData == nil {
		var emptyPost []byte
		return emptyPost, nil
	}

	envelope := postEnvelope{Data: gr.postData}

	postDataXml, err := xml.MarshalIndent(envelope, "", "    ")
	if err != nil {
		return nil, err
	}
	postDataXml = []byte(xml.Header + string(postDataXml))

	return postDataXml, nil
}

func (gr *genericRequest) SetVersion(v string) {
	gr.version = v
}

func (gr genericRequest) SetRequestParam(key string, value string) {
	gr.params.Set(key, value)
}

func (gr *genericRequest) SetPostData(data interface{}) {
	gr.postData = data
}

type clientConfig struct {
	Url  string
	User string
	Key  string
}

func NewClientConfig(url, user, key string, l *log.Logger) (*clientConfig, error) {
	err := validateClientConfigEmail(user, l)
	if err != nil {
		return nil, err
	}

	return &clientConfig{
		Url:  url,
		User: user,
		Key:  key,
	}, nil
}

func validateClientConfigEmail(user string, l *log.Logger) error {
	_, err := mail.ParseAddress(user)
	if err != nil {
		l.Fatalln(err.Error())

		return err
	}
	return nil
}

type client struct {
	httpClient       http.Client
	clientUrlBuilder ClientUrlBuilder
	responseBuilder  ResponseBuilder
	logger           *log.Logger
}

func NewClient(clientConfig clientConfig, l *log.Logger) Client {
	timeout := time.Duration(int64(timeoutInSeconds) * int64(time.Second))

	return &client{
		httpClient:       http.Client{Timeout: timeout},
		clientUrlBuilder: NewClientUrlBuilder(clientConfig),
		responseBuilder:  NewResponseBuilder(),
		logger:           l,
	}
}

func (c client) GetLogger() *log.Logger {
	return c.logger
}

func (c client) Call(request Request) (Response, error) {
	switch request.GetMethod() {
	case MethodGET:
		return c.Get(request)
	case MethodPOST:
		return c.Post(request)
	}

	return nil, NotSupportedMethod
}

func (c client) Get(request Request) (Response, error) {
	getUrl, err := c.clientUrlBuilder.BuildUrl(request.GetRequestParams())
	if err != nil {
		return nil, err
	}

	for i := 1; i <= maxRetries; i++ {
		time.Sleep(time.Duration(i-1) * 200 * time.Millisecond)

		response, err := c.httpClient.Get(getUrl)

		if response == nil || response.StatusCode == 503 {
			c.logger.Printf("Sellercenter Client Get call. empty response or http 503, url: %s, try: %d \n", getUrl, i)
		} else {
			c.logger.Printf("Sellercenter Client Get call. httpResponseCode: %d, url: %s, try: %d \n", response.StatusCode, getUrl, i)

			if err != nil {
				return nil, err
			}

			return c.responseBuilder.BuildResponse(*response)
		}
	}

	if err == nil {
		err = errors.New("empty response")
	}

	return nil, err
}

func (c client) Post(request Request) (Response, error) {
	postUrl, err := c.clientUrlBuilder.BuildUrl(request.GetRequestParams())
	if err != nil {
		return nil, err
	}

	postDataXml, err := request.GeneratePostXml()
	if err != nil {
		return nil, err
	}

	for i := 1; i <= maxRetries; i++ {
		time.Sleep(time.Duration(i-1) * 200 * time.Millisecond)

		response, err := c.httpClient.Post(
			postUrl,
			"text/xml",
			bytes.NewBuffer(postDataXml),
		)

		if response == nil || response.StatusCode == 503 {
			c.logger.Printf("Sellercenter Client Post call. empty response or http 503, url: %s, data: %s, try: %d \n", postUrl, string(postDataXml), i)
		} else {
			c.logger.Printf("Sellercenter Client Post call. httpResponseCode: %d, url: %s, data: %s, try: %d \n", response.StatusCode, postUrl, string(postDataXml), i)

			if err != nil {
				return nil, err
			}

			return c.responseBuilder.BuildResponse(*response)
		}
	}

	if err == nil {
		err = errors.New("empty response")
	}

	return nil, err
}
