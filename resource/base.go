package resource

import (
	"fmt"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/buger/jsonparser"
)

const (
	saleDateTimeFormat = "2006-01-02 15:04:05"
)

func newApiResponseError(errorHead client.HeadErrorResponse) error {
	return &ApiResponseError{
		errorHead.ErrorCode,
		errorHead.ErrorMessage,
	}
}

type ApiResponseError struct {
	Code    string
	Message string
}

func (e *ApiResponseError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func extractRequestId(response client.Response) (string, error) {
	rawHead := response.GetHead()
	requestId, err := jsonparser.GetString(rawHead, "RequestId")

	if err != nil {
		return "", err
	}

	return requestId, nil
}
