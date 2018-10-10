package model

import (
	"encoding/json"
	"github.com/buger/jsonparser"
)

type WebhookEntities struct {
	WebhookEntities []WebhookEntity `json:"Entities"`
}

type WebhookEntity struct {
	Name   string              `json:"Name"`
	Events WebhookEntityEvents `json:"Events"`
}

func (w *WebhookEntities) UnmarshalJSON(b []byte) error {
	rawWebhooks, dataType, _, err := jsonparser.Get(b, "Entities", "Entity")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawWebhooks) == 0 || dataType == jsonparser.NotExist {
		*w = WebhookEntities{[]WebhookEntity{}}
		return nil
	}

	var webhookEntities []WebhookEntity
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawWebhooks, &webhookEntities); nil != err {
			return err
		}
	case jsonparser.Object:
		var webhookEntity WebhookEntity
		if err := json.Unmarshal(rawWebhooks, &webhookEntity); nil != err {
			return err
		}

		webhookEntities = []WebhookEntity{webhookEntity}
	}

	*w = WebhookEntities{webhookEntities}

	return nil
}

type WebhookEntityEvents struct {
	Events []WebhookEntityEvent `json:"Events"`
}

type WebhookEntityEvent struct {
	EventName  string `json:"EventName"`
	EventAlias string `json:"EventAlias"`
}

func (w *WebhookEntityEvents) UnmarshalJSON(b []byte) error {
	rawEvents, dataType, _, err := jsonparser.Get(b, "Event")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawEvents) == 0 || dataType == jsonparser.NotExist {
		*w = WebhookEntityEvents{[]WebhookEntityEvent{}}
		return nil
	}

	var webhookEvents []WebhookEntityEvent
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawEvents, &webhookEvents); nil != err {
			return err
		}
	case jsonparser.Object:
		var webhookEvent WebhookEntityEvent
		if err := json.Unmarshal(rawEvents, &webhookEvent); nil != err {
			return err
		}

		webhookEvents = []WebhookEntityEvent{webhookEvent}
	}

	*w = WebhookEntityEvents{webhookEvents}

	return nil
}

type Webhooks struct {
	Webhooks []Webhook `json:"Webhooks"`
}

type Webhook struct {
	WebhookId     string        `json:"WebhookId"`
	CallbackUrl   string        `json:"CallbackUrl"`
	WebhookSource string        `json:"WebhookSource"`
	Events        WebhookEvents `json:"Events"`
}

func (w *Webhooks) UnmarshalJSON(b []byte) error {
	rawWebhooks, dataType, _, err := jsonparser.Get(b, "Webhooks", "Webhook")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawWebhooks) == 0 || dataType == jsonparser.NotExist {
		*w = Webhooks{[]Webhook{}}
		return nil
	}

	var webhooks []Webhook
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawWebhooks, &webhooks); nil != err {
			return err
		}
	case jsonparser.Object:
		var webhook Webhook
		if err := json.Unmarshal(rawWebhooks, &webhook); nil != err {
			return err
		}

		webhooks = []Webhook{webhook}
	}

	*w = Webhooks{webhooks}

	return nil
}

type WebhookEvents struct {
	Events []WebhookEvent `json:"Events"`
}

type WebhookEvent string

func (w *WebhookEvents) UnmarshalJSON(b []byte) error {
	rawEvents, dataType, _, err := jsonparser.Get(b, "Event")

	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawEvents) == 0 || dataType == jsonparser.NotExist {
		*w = WebhookEvents{[]WebhookEvent{}}
		return nil
	}

	var webhookEvents []WebhookEvent
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawEvents, &webhookEvents); nil != err {
			return err
		}
	case jsonparser.String:
		webhookEvent := WebhookEvent(rawEvents)
		webhookEvents = []WebhookEvent{webhookEvent}
	}

	*w = WebhookEvents{webhookEvents}

	return nil
}
