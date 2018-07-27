package main

import (
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/GFG/seller-center-sdk-go/resource"
	"log"
	"os"
	"time"
)

const (
	scApiBaseUrl = "https://sellerapi.sellercenter.net/"
	scApiUser    = "user@sellercenter.net"
	scApiKey     = "000000000000000000000000000000000000000"
)

func main() {
	logger := log.New(os.Stdout, "SC SDK", log.LstdFlags)

	feedResource := getFeedResource(logger)

	feedListExample(logger, feedResource)
	feedStatusExample(logger, feedResource)
}

func feedListExample(logger *log.Logger, feedResource resource.FeedResource) {
	feedList, err := feedResource.FeedList()

	if err != nil {
		logger.Panicln(err)
	}

	for _, feed := range feedList.Feeds {
		dumpFeed(feed, logger)
	}
}

func feedStatusExample(logger *log.Logger, feedResource resource.FeedResource) {
	feedStatus, err := feedResource.FeedStatus("89e767bc-bf18-4f92-88a9-24368bd6a08c")

	if err != nil {
		logger.Panicln(err)
	}

	dumpFeedStatus(feedStatus, logger)
}

func getFeedResource(logger *log.Logger) resource.FeedResource {
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

	return resource.NewFeed(scClient)
}

func dumpFeed(feed model.Feed, logger *log.Logger) {
	logger.Println("---------")
	logger.Printf("Feed: %s\n", feed.Feed)
	logger.Printf("Status: %s\n", feed.Status)
	logger.Printf("Action: %s\n", feed.Action)
	logger.Printf("CreationDate: %s\n", time.Time(feed.CreationDate).Format("2006-01-02 15:04:05"))
	logger.Printf("UpdatedDate: %s\n", time.Time(feed.UpdatedDate).Format("2006-01-02 15:04:05"))
	logger.Printf("Source: %s\n", feed.Source)
	logger.Printf("TotalRecords: %d\n", feed.TotalRecords)
	logger.Printf("ProcessedRecords: %d\n", feed.ProcessedRecords)
	logger.Printf("FailedRecords: %d\n", feed.FailedRecords)
	logger.Println("FailureReports:")
	logger.Printf("   Format: %s\n", feed.FailureReports.MimeType)
	logger.Printf("   File: %s\n", feed.FailureReports.File)
}

func dumpFeedStatus(feedStatus model.FeedStatus, logger *log.Logger) {
	logger.Println("---------")
	logger.Printf("Feed: %s\n", feedStatus.Feed)
	logger.Printf("Status: %s\n", feedStatus.Status)
	logger.Printf("Action: %s\n", feedStatus.Action)
	logger.Printf("CreationDate: %s\n", time.Time(feedStatus.CreationDate).Format("2006-01-02 15:04:05"))
	logger.Printf("UpdatedDate: %s\n", time.Time(feedStatus.UpdatedDate).Format("2006-01-02 15:04:05"))
	logger.Printf("Source: %s\n", feedStatus.Source)
	logger.Printf("TotalRecords: %d\n", feedStatus.TotalRecords)
	logger.Printf("ProcessedRecords: %d\n", feedStatus.ProcessedRecords)
	logger.Printf("FailedRecords: %d\n", feedStatus.FailedRecords)
	logger.Println("Errors:")
	for _, feedErr := range feedStatus.FeedErrors.Errors {
		logger.Println("---------")
		logger.Printf("	 Code: %s\n", feedErr.Code)
		logger.Printf("	 Message: %s\n", feedErr.Message)
		logger.Printf("	 SellerSku: %s\n", feedErr.SellerSku)
		logger.Printf("	 OrderId: %s\n", feedErr.OrderId)
		logger.Printf("	 OrderItemId: %s\n", feedErr.OrderItemId)
	}

	logger.Println("Warnings:")
	for _, warning := range feedStatus.FeedWarnings.Warnings {
		logger.Println("---------")
		logger.Printf("	 Message: %s\n", warning.Message)
		logger.Printf("	 SellerSku: %s\n", warning.SellerSku)
		logger.Printf("	 OrderId: %s\n", warning.OrderId)
		logger.Printf("	 OrderItemId: %s\n", warning.OrderItemId)
	}
}
