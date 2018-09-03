package client

import (
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
)

type Response interface {
	IsError() bool
	GetHeadObject() interface{}
	GetHead() []byte
	GetBody() []byte
}

type SuccessResponse struct {
	HeadObject headSuccessResponse
	Head       []byte
	Body       []byte
}

func (sr SuccessResponse) IsError() bool {
	return false
}

func (sr SuccessResponse) GetHeadObject() interface{} {
	return sr.HeadObject
}

func (sr SuccessResponse) GetHead() []byte {
	return sr.Head
}

func (sr SuccessResponse) GetBody() []byte {
	return sr.Body
}

type headSuccessResponse struct {
	RequestId     string `json:"RequestId"`
	RequestAction string `json:"RequestAction"`
	ResponseType  string `json:"ResponseType"`
	Timestamp     string `json:"Timestamp"`
}

type ErrorResponse struct {
	HeadObject HeadErrorResponse `json:"Head"`
	Head       []byte
}

type HeadErrorResponse struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

func (er ErrorResponse) IsError() bool {
	return true
}

func (er ErrorResponse) GetHeadObject() interface{} {
	return er.HeadObject
}

func (er ErrorResponse) GetHead() []byte {
	return er.Head
}

func (er ErrorResponse) GetBody() []byte {
	return nil
}

// Errors
var (
	NoHttp200ResponseError = errors.New("unexpected response")
)

type ResponseBuilder interface {
	BuildResponse(response http.Response) (Response, error)
}

func NewResponseBuilder() responseBuilder {
	return responseBuilder{}
}

type responseBuilder struct {
}

func (rb responseBuilder) BuildResponse(response http.Response) (Response, error) {
	// ... check if http code 200
	if response.StatusCode != http.StatusOK {
		return nil, NoHttp200ResponseError
	}

	// ... read http response body
	defer response.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// ... handle error response
	errorResponse, err := rb.handleErrorResponse(responseBodyBytes)
	if err == nil {
		return errorResponse, nil
	}

	// ... http response seems to be really broken
	if err != jsonparser.KeyPathNotFoundError {
		return nil, err
	}

	// ... handle success response
	return rb.handleSuccessResponse(responseBodyBytes)
}

func (rb responseBuilder) handleErrorResponse(responseBodyBytes []byte) (Response, error) {
	errorResponseData, _, _, err := jsonparser.Get(responseBodyBytes, "ErrorResponse")

	if err != nil {
		return nil, err
	}

	var errorResponse ErrorResponse
	err = json.Unmarshal(errorResponseData, &errorResponse)
	if err != nil {
		return nil, err
	}

	errorResponseDataHead, _, _, err := jsonparser.Get(responseBodyBytes, "ErrorResponse", "Head")

	if err != nil {
		return nil, err
	}

	errorResponse.Head = errorResponseDataHead

	return errorResponse, nil
}

func (rb responseBuilder) handleSuccessResponse(responseBodyBytes []byte) (Response, error) {
	_, _, _, err := jsonparser.Get(responseBodyBytes, "SuccessResponse")

	// ... broken body of HTTP response
	if err != nil {
		return nil, err
	}

	// ... Head
	successResponseHeadData, _, _, err := jsonparser.Get(responseBodyBytes, "SuccessResponse", "Head")
	if err != nil {
		return nil, err
	}

	var successResponseHead headSuccessResponse
	err = json.Unmarshal(successResponseHeadData, &successResponseHead)
	if err != nil {
		return nil, err
	}

	// ... Body
	successResponseBodyData, _, _, err := jsonparser.Get(responseBodyBytes, "SuccessResponse", "Body")
	if err != nil {
		return nil, err
	}

	successResponse := SuccessResponse{
		HeadObject: successResponseHead,
		Head:       successResponseHeadData,
		Body:       successResponseBodyData,
	}

	return successResponse, nil
}
