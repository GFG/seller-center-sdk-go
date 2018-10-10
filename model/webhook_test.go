package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_WebhookEntitiesEmpty(t *testing.T) {
	j := []byte(`{}`)

	expected := WebhookEntities{[]WebhookEntity{}}

	var c WebhookEntities
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_WebhookEntitiesSingle(t *testing.T) {
	j := []byte(`{"Entities":{"Entity":{"Name":"Feed","Events":{"Event":{"EventName":"Completed", "EventAlias":"onFeedCompleted"} } } } }`)

	expected := WebhookEntities{[]WebhookEntity{
		{
			"Feed",
			WebhookEntityEvents{
				[]WebhookEntityEvent{
					{
						"Completed",
						"onFeedCompleted",
					},
				},
			},
		},
	},
	}

	var c WebhookEntities
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_WebhookEntitiesMultiple(t *testing.T) {
	j := []byte(`{"Entities":{"Entity":[ {"Name":"Feed","Events":{"Event":{"EventName":"Completed", "EventAlias":"onFeedCompleted"} } }, {"Name":"Order","Events":{"Event": [ {"EventName":"Created", "EventAlias":"onOrderCreated"}, {"EventName":"ItemsStatusChanged", "EventAlias":"onOrderItemsStatusChanged"} ] } } ] } }`)

	expected := WebhookEntities{[]WebhookEntity{
		{
			"Feed",
			WebhookEntityEvents{
				[]WebhookEntityEvent{
					{
						"Completed",
						"onFeedCompleted",
					},
				},
			},
		},
		{
			"Order",
			WebhookEntityEvents{
				[]WebhookEntityEvent{
					{
						"Created",
						"onOrderCreated",
					},
					{
						"ItemsStatusChanged",
						"onOrderItemsStatusChanged",
					},
				},
			},
		},
	},
	}

	var c WebhookEntities
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_WebhooksEmpty(t *testing.T) {
	j := []byte(`{}`)

	expected := Webhooks{[]Webhook{}}

	var c Webhooks
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_WebhooksSingle(t *testing.T) {
	j := []byte(`{"Webhooks":{"Webhook":{"WebhookId":"ae45284c-a882-4c04-a4c9-f7099166d9fd","CallbackUrl":"https://www.shop.com/webhook","WebhookSource":"web","Events":{"Event":"onFeedCompleted" } } } }`)

	expected := Webhooks{[]Webhook{
		{
			"ae45284c-a882-4c04-a4c9-f7099166d9fd",
			"https://www.shop.com/webhook",
			"web",
			WebhookEvents{[]WebhookEvent{"onFeedCompleted"}},
		},
	},
	}

	var c Webhooks
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_WebhooksMultiple(t *testing.T) {
	j := []byte(`{"Webhooks":{"Webhook":[{"WebhookId":"ae45284c-a882-4c04-a4c9-f7099166d9fd","CallbackUrl":"https://www.shop.com/webhook/1","WebhookSource":"web","Events":{"Event":"onFeedCompleted" } }, {"WebhookId":"691957b2-a9da-4c08-9a53-269fd1c39b15","CallbackUrl":"https://www.shop.com/webhook/2","WebhookSource":"api","Events":{"Event":["onOrderCreated", "onProductCreated"] } }] } }`)

	expected := Webhooks{[]Webhook{
		{
			"ae45284c-a882-4c04-a4c9-f7099166d9fd",
			"https://www.shop.com/webhook/1",
			"web",
			WebhookEvents{[]WebhookEvent{"onFeedCompleted"}},
		},
		{
			"691957b2-a9da-4c08-9a53-269fd1c39b15",
			"https://www.shop.com/webhook/2",
			"api",
			WebhookEvents{[]WebhookEvent{"onOrderCreated", "onProductCreated"}},
		},
	},
	}

	var c Webhooks
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}
