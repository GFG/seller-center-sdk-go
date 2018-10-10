package resource

import (
	"encoding/json"
	"encoding/xml"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
)

type WebhookResource struct {
	client client.Client
}

func NewWebhook(client client.Client) WebhookResource {
	return WebhookResource{client: client}
}

func (wr WebhookResource) CreateWebhook(callbackUrl string, events []string) (bool, error) {
	r := client.NewGenericRequest("CreateWebhook", client.MethodPOST)
	r.SetVersion(client.V1)

	postData := webhookXmlBody{
		CallbackUrl: callbackUrl,
		Events:      webhookEventsEntries{Event: events},
	}

	r.SetPostData(postData)

	response, err := wr.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return false, newApiResponseError(errorResponse.HeadObject)
	}

	return true, nil
}

func (wr WebhookResource) GetWebhookEntities() (model.WebhookEntities, error) {
	r := client.NewGenericRequest("GetWebhookEntities", client.MethodGET)
	r.SetVersion(client.V1)

	response, err := wr.client.Call(r)

	if err != nil {
		return model.WebhookEntities{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.WebhookEntities{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()

	if len(rawBody) == 0 {
		return model.WebhookEntities{}, nil
	}

	var webhookEntities model.WebhookEntities
	if err := json.Unmarshal(rawBody, &webhookEntities); nil != err {
		return model.WebhookEntities{}, err
	}

	return webhookEntities, nil
}

func (wr WebhookResource) GetWebhooks() (model.Webhooks, error) {
	r := client.NewGenericRequest("GetWebhooks", client.MethodGET)
	r.SetVersion(client.V1)

	response, err := wr.client.Call(r)

	if err != nil {
		return model.Webhooks{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Webhooks{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()

	if len(rawBody) == 0 {
		return model.Webhooks{}, nil
	}

	var webhooks model.Webhooks
	if err := json.Unmarshal(rawBody, &webhooks); nil != err {
		return model.Webhooks{}, err
	}

	return webhooks, nil
}

type webhookXmlBody struct {
	XMLName     xml.Name `xml:"Webhook"`
	CallbackUrl string   `xml:"CallbackUrl`
	Events      webhookEventsEntries
}

type webhookEventsEntries struct {
	XMLName xml.Name `xml:"Events"`
	Event   []string `xml:"Events>Event`
}
