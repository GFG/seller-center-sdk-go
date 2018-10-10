package resource

import (
	"errors"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"reflect"
	"testing"
)

const callbackUrl = "https://sellerapi.sellercenter.net/"

var webhookEvents = []string{"aWebhookEvent"}

func Test_Get_WebhookEntities_Returns_Client_Error(t *testing.T) {
	clientError := errors.New("Timeout")

	fakeClient := client.FakeClient{
		FakeResponse: nil,
		FakeError:    clientError,
	}

	resource := NewWebhook(fakeClient)

	webhookEntities, err := resource.GetWebhookEntities()

	if !reflect.DeepEqual(clientError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", clientError, err)
	}

	empty := model.WebhookEntities{}

	if !reflect.DeepEqual(empty, webhookEntities) {
		t.Fatalf("did not return empty GetWebhookEntities. received: `%v`.", webhookEntities)
	}
}

func Test_Get_WebhookEntities_Can_Handle_ErrorResponse(t *testing.T) {
	clientErrorResponse := client.ErrorResponse{
		HeadObject: client.HeadErrorResponse{
			ErrorCode:    "E27",
			ErrorMessage: "Wrong signature",
		},
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientErrorResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	webhookEntities, err := resource.GetWebhookEntities()

	expectedError := &ApiResponseError{
		"E27",
		"Wrong signature",
	}

	if !reflect.DeepEqual(expectedError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", expectedError, err)
	}

	empty := model.WebhookEntities{}

	if !reflect.DeepEqual(empty, webhookEntities) {
		t.Fatalf("did not return empty GetWebhookEntities. received: `%v`.", webhookEntities)
	}
}

func Test_Get_WebhookEntities_Can_Handle_EmptyResponseBody(t *testing.T) {
	clientResponse := client.SuccessResponse{
		Body: []byte(``),
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	webhookEntities, err := resource.GetWebhookEntities()

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	empty := model.WebhookEntities{}

	if !reflect.DeepEqual(empty, webhookEntities) {
		t.Fatalf("did not return empty GetWebhookEntities. received: `%v`.", webhookEntities)
	}
}

func Test_Get_WebhookEntities_Can_Handle_Response(t *testing.T) {
	payloadBody := []byte(`{"Entities":{"Entity":[ {"Name":"Feed","Events":{"Event":{"EventName":"Completed", "EventAlias":"onFeedCompleted"} } }, {"Name":"Order","Events":{"Event": [ {"EventName":"Created", "EventAlias":"onOrderCreated"}, {"EventName":"ItemsStatusChanged", "EventAlias":"onOrderItemsStatusChanged"} ] } } ] } }`)

	clientResponse := client.SuccessResponse{
		Body: payloadBody,
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	webhookEntities, err := resource.GetWebhookEntities()

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	if len(webhookEntities.WebhookEntities) != 2 {
		t.Fatalf("did not receive GetWebhookEntities with 2 items. received: `%d`.", len(webhookEntities.WebhookEntities))
	}
}

func Test_Get_Webhooks_Returns_Client_Error(t *testing.T) {
	clientError := errors.New("Timeout")

	fakeClient := client.FakeClient{
		FakeResponse: nil,
		FakeError:    clientError,
	}

	resource := NewWebhook(fakeClient)

	webhooks, err := resource.GetWebhooks()

	if !reflect.DeepEqual(clientError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", clientError, err)
	}

	empty := model.Webhooks{}

	if !reflect.DeepEqual(empty, webhooks) {
		t.Fatalf("did not return empty GetWebhooks. received: `%v`.", webhooks)
	}
}

func Test_Get_Webhooks_Can_Handle_ErrorResponse(t *testing.T) {
	clientErrorResponse := client.ErrorResponse{
		HeadObject: client.HeadErrorResponse{
			ErrorCode:    "E27",
			ErrorMessage: "Wrong signature",
		},
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientErrorResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	webhooks, err := resource.GetWebhooks()

	expectedError := &ApiResponseError{
		"E27",
		"Wrong signature",
	}

	if !reflect.DeepEqual(expectedError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", expectedError, err)
	}

	empty := model.Webhooks{}

	if !reflect.DeepEqual(empty, webhooks) {
		t.Fatalf("did not return empty GetWebhooks. received: `%v`.", webhooks)
	}
}

func Test_Get_Webhooks_Can_Handle_EmptyResponseBody(t *testing.T) {
	clientResponse := client.SuccessResponse{
		Body: []byte(``),
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	webhooks, err := resource.GetWebhooks()

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	empty := model.Webhooks{}

	if !reflect.DeepEqual(empty, webhooks) {
		t.Fatalf("did not return empty GetWebhooks. received: `%v`.", webhooks)
	}
}

func Test_Get_Webhooks_Can_Handle_Response(t *testing.T) {
	payloadBody := []byte(`{"Webhooks":{"Webhook":[{"WebhookId":"ae45284c-a882-4c04-a4c9-f7099166d9fd","CallbackUrl":"https://www.shop.com/webhook/1","WebhookSource":"web","Events":{"Event":"onFeedCompleted" } }, {"WebhookId":"691957b2-a9da-4c08-9a53-269fd1c39b15","CallbackUrl":"https://www.shop.com/webhook/2","WebhookSource":"api","Events":{"Event":["onOrderCreated", "onProductCreated"] } }] } }`)

	clientResponse := client.SuccessResponse{
		Body: payloadBody,
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	webhooks, err := resource.GetWebhooks()

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	if len(webhooks.Webhooks) != 2 {
		t.Fatalf("did not receive GetWebhooks with 2 items. received: `%d`.", len(webhooks.Webhooks))
	}
}

func Test_Post_CreateWebhook_Returns_Client_Error(t *testing.T) {
	clientError := errors.New("Timeout")

	fakeClient := client.FakeClient{
		FakeResponse: nil,
		FakeError:    clientError,
	}

	resource := NewWebhook(fakeClient)

	success, err := resource.CreateWebhook(callbackUrl, webhookEvents)

	if !reflect.DeepEqual(clientError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", clientError, err)
	}

	if success != false {
		t.Fatalf("did not return false CreateWebhook. received: `%t`.", success)
	}
}

func Test_Post_CreateWebhook_Can_Handle_ErrorResponse(t *testing.T) {
	clientErrorResponse := client.ErrorResponse{
		HeadObject: client.HeadErrorResponse{
			ErrorCode:    "E27",
			ErrorMessage: "Wrong signature",
		},
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientErrorResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	success, err := resource.CreateWebhook(callbackUrl, webhookEvents)

	expectedError := &ApiResponseError{
		"E27",
		"Wrong signature",
	}

	if !reflect.DeepEqual(expectedError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", expectedError, err)
	}

	if success != false {
		t.Fatalf("did not return false CreateWebhook. received: `%t`.", success)
	}
}

func Test_Post_CreateWebhook_Can_Handle_EmptyResponseBody(t *testing.T) {
	clientResponse := client.SuccessResponse{
		Body: []byte(``),
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	success, err := resource.CreateWebhook(callbackUrl, webhookEvents)

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	if success != true {
		t.Fatalf("did not return true CreateWebhook. received: `%t`.", success)
	}
}

func Test_Post_CreateWebhook_Can_Handle_Response(t *testing.T) {
	payloadBody := []byte(`"Webhook": { "WebhookId": "60ecb631-d237-4092-a336-fa9d8aeddde5", "CreatedAt": "2018-10-10T11:22:10+0200" }`)

	clientResponse := client.SuccessResponse{
		Body: payloadBody,
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewWebhook(fakeClient)

	success, err := resource.CreateWebhook(callbackUrl, webhookEvents)

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	if success != true {
		t.Fatalf("did not receive CreateWebhook successful. received: `%t`.", success)
	}
}
