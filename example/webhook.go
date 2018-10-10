package main

import (
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/GFG/seller-center-sdk-go/resource"
	"log"
	"os"
)

const (
	scApiBaseUrl = "https://sellerapi.sellercenter.net/"
	scApiUser    = "user@sellercenter.net"
	scApiKey     = "000000000000000000000000000000000000000"
)

func main() {
	logger := log.New(os.Stdout, "SC SDK", log.LstdFlags)

	clientConfig, err := client.NewClientConfig(
		scApiBaseUrl,
		scApiUser,
		scApiKey,
		logger,
	)

	if err != nil {
		logger.Panicf("%s\n", err)
	}

	scClient := client.NewClient(*clientConfig, logger)

	if nil == scClient {
		logger.Panicln("No client available")
	}

	webhookResource := resource.NewWebhook(scClient)

	getWebhookEntitiesExample(logger, webhookResource)

	createWebhookExample(logger, webhookResource)

	getWebhooksExample(logger, webhookResource)
}

func getWebhookEntitiesExample(logger *log.Logger, webhookResource resource.WebhookResource) {
	webhookEntities, err := webhookResource.GetWebhookEntities()

	if err != nil {
		logger.Panicln(err)
	}

	for _, webhookEntity := range webhookEntities.WebhookEntities {
		dumpWebhookEntity(webhookEntity, logger)
	}
}

func createWebhookExample(logger *log.Logger, webhookResource resource.WebhookResource) {
	callbackUrl := "https://sellerapi.sellercenter.net/"
	events := []string{"onOrderCreated", "onProductCreated"}

	webhookCreated, err := webhookResource.CreateWebhook(callbackUrl, events)
	if err != nil {
		logger.Printf("CreateWebhook failed: %s\n", err)
	} else {
		logger.Println("CreateWebhook created: %t", webhookCreated)
	}
}

func getWebhooksExample(logger *log.Logger, webhookResource resource.WebhookResource) {
	webhooks, err := webhookResource.GetWebhooks()

	if err != nil {
		logger.Panicln(err)
	}

	for _, webhook := range webhooks.Webhooks {
		dumpWebhook(webhook, logger)
	}
}

func dumpWebhookEntity(webhookEntity model.WebhookEntity, logger *log.Logger) {
	logger.Println("Name")
	logger.Printf("Name: %s\n", webhookEntity.Name)
	logger.Println("Events:")
	for _, event := range webhookEntity.Events.Events {
		logger.Printf("	EventName: %s\n", event.EventName)
		logger.Printf("	EventAlias: %s\n", event.EventAlias)
	}
}

func dumpWebhook(webhook model.Webhook, logger *log.Logger) {
	logger.Println("Webhook")
	logger.Printf("WebhookId: %s\n", webhook.WebhookId)
	logger.Printf("CallbackUrl: %s\n", webhook.CallbackUrl)
	logger.Printf("WebhookSource: %s\n", webhook.WebhookSource)
	logger.Println("Events:")
	for _, event := range webhook.Events.Events {
		logger.Printf("	%s\n", event)
	}
}
