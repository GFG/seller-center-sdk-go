package client

import (
	"github.com/buger/jsonparser"
	"net/http"
	"strings"
	"testing"
)

type MyReadCloser struct {
	*strings.Reader
}

func (rc MyReadCloser) Close() error {
	return nil
}

func Test_Can_Build_Success_Response(t *testing.T) {

	body := `{
  "SuccessResponse": {
    "Head": {
      "RequestId": "789",
      "RequestAction": "CreateWebhook",
      "ResponseType": "Webhook",
      "Timestamp": "2018-07-06T15:37:57+0200"
    },
    "Body": {
      "Webhook": {
        "WebhookId": "a241742a-9ea8-44e1-a7e1-5d758fd7d019",
        "CreatedAt": "2018-07-06T15:37:57+0200"
      }
    }
  }
}`

	bodyReader := &MyReadCloser{strings.NewReader(body)}

	httpResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyReader,
	}

	responseBuilder := NewResponseBuilder()

	response, err := responseBuilder.BuildResponse(httpResponse)

	if err != nil {
		t.Fatal("can not build response. expected error to be nil.")
	}

	if response.IsError() == true {
		t.Fatal("can not build response. expected get SuccessResponse.")
	}

	successResponse := SuccessResponse{
		HeadObject: response.GetHeadObject().(headSuccessResponse),
		Head:       response.GetHead(),
		Body:       response.GetBody(),
	}

	if successResponse.HeadObject.ResponseType != "Webhook" ||
		successResponse.HeadObject.RequestAction != "CreateWebhook" ||
		successResponse.HeadObject.Timestamp != "2018-07-06T15:37:57+0200" ||
		successResponse.HeadObject.RequestId != "789" {
		t.Fatal("can not build response. failed to build successHeader.")
	}

	if nil == successResponse.GetHead() {
		t.Fatal("can not build response. failed to set Head property.")
	}

	responseType, err := jsonparser.GetString(response.GetHead(), "ResponseType")
	if err != nil || responseType != "Webhook" {
		t.Fatalf("can not build response. failed to parse Head property ResponseType. actual: `%s`, expected: `Webhook`", responseType)
	}

	if nil == successResponse.GetBody() {
		t.Fatal("can not build response. failed to set Body property.")
	}

	id, err := jsonparser.GetString(response.GetBody(), "Webhook", "WebhookId")
	if err != nil || id != "a241742a-9ea8-44e1-a7e1-5d758fd7d019" {
		t.Fatalf("can not build response. failed to parse Bodu property WebhookId. actual: `%s`, expected: `a241742a-9ea8-44e1-a7e1-5d758fd7d019`", id)
	}

}

func Test_Can_Build_Error_Response(t *testing.T) {

	body := `{
  "ErrorResponse": {
    "Head": {
      "RequestAction": "CreateWebhook",
      "ErrorType": "Sender",
      "ErrorCode": "98",
      "ErrorMessage": "E098: Invalid Webhook Callback Url, \"The callback url (https://www.shop.com/my-super-webhook-new) is already in use.\""
    },
    "Body": ""
  }
}`

	bodyReader := &MyReadCloser{strings.NewReader(body)}

	httpResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyReader,
	}

	responseBuilder := NewResponseBuilder()

	response, err := responseBuilder.BuildResponse(httpResponse)

	if err != nil {
		t.Fatal("can not build error response. expected error to be nil.")
	}

	if response.IsError() == false {
		t.Fatal("can not build response. expected get ErrorResponse.")
	}

	errorResponse := ErrorResponse{
		HeadObject: response.GetHeadObject().(HeadErrorResponse),
		Head:       response.GetHead(),
	}

	if errorResponse.HeadObject.ErrorCode != "98" ||
		errorResponse.HeadObject.ErrorMessage != "E098: Invalid Webhook Callback Url, \"The callback url (https://www.shop.com/my-super-webhook-new) is already in use.\"" {
		t.Fatal("can not build response. failed to build errorHeader.")
	}

	if nil == errorResponse.GetHead() {
		t.Fatal("can not build error response. failed to set Head property.")
	}

	a, err := jsonparser.GetString(response.GetHead(), "RequestAction")
	if err != nil || a != "CreateWebhook" {
		t.Fatalf("can not build error response. failed to parse Head property RequestAction. actual: `%s`, expected: `CreateWebhook`", a)
	}

	if nil != errorResponse.GetBody() {
		t.Fatal("can not build error response. Body is expected to be empty.")
	}
}

func Test_Build_Response_With_Http_400(t *testing.T) {

	body := `{}`

	bodyReader := &MyReadCloser{strings.NewReader(body)}

	httpResponse := http.Response{
		StatusCode: 400,
		Body:       bodyReader,
	}

	responseBuilder := NewResponseBuilder()

	response, err := responseBuilder.BuildResponse(httpResponse)

	if err != NoHttp200ResponseError {
		t.Fatalf("expected to fail to build response, but actual `%s`.", err)
	}

	if response != nil {
		t.Fatalf("expected response to be nil, actual `%s`.", response)
	}
}

func Test_Build_Response_With_Broken_Body(t *testing.T) {

	brokenBody := `{`

	bodyReader := &MyReadCloser{strings.NewReader(brokenBody)}

	httpResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyReader,
	}

	responseBuilder := NewResponseBuilder()

	response, err := responseBuilder.BuildResponse(httpResponse)

	if err != jsonparser.KeyPathNotFoundError {
		t.Fatalf("expected to fail to build response, but actual `%s`.", err)
	}

	if response != nil {
		t.Fatalf("expected response to be nil, actual `%s`.", response)
	}
}

func Test_Build_Response_With_Incomplete_Body(t *testing.T) {

	// Head is missing
	body := `{
  "SuccessResponse": {
    
    "Body": {
      "Webhook": {
        "WebhookId": "a241742a-9ea8-44e1-a7e1-5d758fd7d019",
        "CreatedAt": "2018-07-06T15:37:57+0200"
      }
    }
  }
}`

	bodyReader := &MyReadCloser{strings.NewReader(body)}

	httpResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyReader,
	}

	responseBuilder := NewResponseBuilder()

	response, err := responseBuilder.BuildResponse(httpResponse)

	if err != jsonparser.KeyPathNotFoundError {
		t.Fatalf("expected to fail to build response, but actual `%s`.", err)
	}

	if response != nil {
		t.Fatalf("expected response to be nil, actual `%s`.", response)
	}
}

func Test_Build_Response_With_Incomplete_Body2(t *testing.T) {

	// Body is missing
	body := `{
  "SuccessResponse": {
    "Head": {
      "RequestId": "789",
      "RequestAction": "CreateWebhook",
      "ResponseType": "Webhook",
      "Timestamp": "2018-07-06T15:37:57+0200"
    },
  }
}`

	bodyReader := &MyReadCloser{strings.NewReader(body)}

	httpResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyReader,
	}

	responseBuilder := NewResponseBuilder()

	response, err := responseBuilder.BuildResponse(httpResponse)

	if err != jsonparser.KeyPathNotFoundError {
		t.Fatalf("expected to fail to build response, but actual `%s`.", err)
	}

	if response != nil {
		t.Fatalf("expected response to be nil, actual `%s`.", response)
	}
}
