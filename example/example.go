package main

import (
	"encoding/base64"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/endpoint"
	"github.com/GFG/seller-center-sdk-go/model"
	"io/ioutil"
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

	orderEndpoint := endpoint.NewOrder(scClient)

	createdAfter := time.Date(2014, 2, 25, 0, 0, 0, 0, time.UTC)
	var createdBefore *time.Time
	createdBefore = nil
	var updatedAfter *time.Time
	updatedAfter = nil
	var updatedBefore *time.Time
	updatedBefore = nil
	status := "shipped"
	limit := 5
	offset := 2
	sortBy := "created_at"
	sortDirection := "DESC"

	orders, err := orderEndpoint.GetOrders(
		&createdAfter,
		createdBefore,
		updatedAfter,
		updatedBefore,
		&status,
		&limit,
		&offset,
		&sortBy,
		&sortDirection)

	if err != nil {
		logger.Panicln(err)
	}

	for _, order := range orders.Orders {
		dumpOder(order, logger)
	}

	order, err := orderEndpoint.GetOrder(1)

	if err != nil {
		logger.Panicln(err)
	}

	dumpOder(order, logger)

	orderItems, err := orderEndpoint.GetOrderItems(1)
	if err != nil {
		logger.Panicln(err)
	}

	for _, orderItem := range orderItems.OrderItems {
		dumpOderItem(orderItem, logger)
	}

	orderItemIds := []int{1}
	document, err := orderEndpoint.GetDocument(orderItemIds, model.DocumentTypeShippingLabel)
	if err != nil {
		logger.Panicln(err)
	}

	dumpDocument(document, logger)

	decodedFile, err := base64.StdEncoding.DecodeString(document.File)
	if err != nil {
		logger.Panicln(err)
	}

	file, err := ioutil.TempFile(os.TempDir(), "")
	defer file.Close()

	if err != nil {
		logger.Panicln(err)
	}

	_, err = file.Write(decodedFile)
	if err != nil {
		logger.Panicln(err)
	}

	shippingProivder := "DHL"
	trackingNumber := "TRACK-0001"

	success, err := orderEndpoint.SetStatusToReadyToShip(orderItemIds, model.DeliveryTypeDropshipping, shippingProivder, trackingNumber)
	if false == success {
		logger.Printf("SetStatusToReadyToShip failed: %s\n", err)
	} else {
		logger.Println("SetStatusToReadyToShip succeeded")
	}

	success, err = orderEndpoint.SetStatusToPackedByMarketplace(orderItemIds, model.DeliveryTypeDropshipping, shippingProivder)
	if false == success {
		logger.Printf("SetStatusToPackedByMarketplace failed: %s\n", err)
	} else {
		logger.Println("SetStatusToPackedByMarketplace succeeded")
	}

	reason := "Out of Stock"
	reasonDetail := "No more invetory"
	success, err = orderEndpoint.SetStatusToCanceled(1, reason, reasonDetail)
	if false == success {
		logger.Printf("SetStatusToCanceled failed: %s\n", err)
	} else {
		logger.Println("SetStatusToCanceled succeeded")
	}
}

func dumpDocument(document model.Document, logger *log.Logger) {
	logger.Printf("DocumentType: %s\n", document.DocumentType)
	logger.Printf("MimeType: %s\n", document.MimeType)
	logger.Printf("File: %s\n", document.File)
}

func dumpOderItem(orderItem model.OrderItem, logger *log.Logger) {
	logger.Printf("OrderItemId: %d\n", orderItem.OrderItemId)
	logger.Printf("ShopId: %s\n", orderItem.ShopId)
	logger.Printf("OrderId: %d\n", orderItem.OrderId)
	logger.Printf("Name: %s\n", orderItem.Name)
	logger.Printf("Sku: %s\n", orderItem.Sku)
	logger.Printf("Variation: %s\n", orderItem.Variation)
	logger.Printf("ShopSku: %s\n", orderItem.ShopSku)
	logger.Printf("ShippingType: %s\n", orderItem.ShippingType)
	logger.Printf("ItemPrice: %f\n", orderItem.ItemPrice)
	logger.Printf("PaidPrice: %f\n", orderItem.PaidPrice)
	logger.Printf("Currency: %s\n", orderItem.Currency)
	logger.Printf("WalletCredits: %f\n", orderItem.WalletCredits)
	logger.Printf("TaxAmount: %f\n", orderItem.TaxAmount)
	logger.Printf("CodCollectableAmount: %f\n", orderItem.CodCollectableAmount)
	logger.Printf("ShippingAmount: %f\n", orderItem.ShippingAmount)
	logger.Printf("ShippingServiceCost: %f\n", orderItem.ShippingServiceCost)
	logger.Printf("VoucherAmount: %f\n", orderItem.VoucherAmount)
	logger.Printf("VoucherCode: %s\n", orderItem.VoucherCode)
	logger.Printf("Status: %s\n", orderItem.Status)
	logger.Printf("IsProcessable: %t\n", orderItem.IsProcessable)
	logger.Printf("ShipmentProvider: %s\n", orderItem.ShipmentProvider)
	logger.Printf("IsDigital: %t\n", orderItem.IsDigital)
	logger.Printf("DigitalDeliveryInfo: %s\n", orderItem.DigitalDeliveryInfo)
	logger.Printf("TrackingCode: %s\n", orderItem.TrackingCode)
	logger.Printf("TrackingCodePre: %s\n", orderItem.TrackingCodePre)
	logger.Printf("Reason: %s\n", orderItem.Reason)
	logger.Printf("ReasonDetail: %s\n", orderItem.ReasonDetail)
	logger.Printf("PurchaseOrderId: %d\n", orderItem.PurchaseOrderId)
	logger.Printf("PurchaseOrderNumber: %s\n", orderItem.PurchaseOrderNumber)
	logger.Printf("PackageId: %s\n", orderItem.PackageId)
	logger.Printf("PromisedShippingTime: %s\n", time.Time(orderItem.PromisedShippingTime).Format("2006-01-02 15:04:05"))
	logger.Printf("ExtraAttributes: %s\n", orderItem.ExtraAttributes)
	logger.Printf("ShippingProviderType: %s\n", orderItem.ShippingProviderType)
	logger.Printf("CreatedAt: %s\n", time.Time(orderItem.CreatedAt).Format("2006-01-02 15:04:05"))
	logger.Printf("UpdatedAt: %s\n", time.Time(orderItem.UpdatedAt).Format("2006-01-02 15:04:05"))
	logger.Printf("ReturnStatus: %s\n", orderItem.ReturnStatus)
}

func dumpOder(order model.Order, logger *log.Logger) {
	logger.Printf("OrderId: %d\n", order.OrderId)
	logger.Printf("OrderNumber: %s\n", order.OrderNumber)
	logger.Printf("PaymentMethod: %s\n", order.PaymentMethod)
	logger.Printf("Remarks: %s\n", order.Remarks)
	logger.Printf("DeliveryInfo: %s\n", order.DeliveryInfo)
	logger.Printf("Price: %f\n", order.Price)
	logger.Printf("GiftOption: %t\n", order.GiftOption)
	logger.Printf("GiftMessage: %s\n", order.GiftMessage)
	logger.Printf("VoucherCode: %s\n", order.VoucherCode)
	logger.Printf("CreatedAt: %s\n", time.Time(order.CreatedAt).Format("2006-01-02 15:04:05"))
	logger.Printf("UpdatedAt: %s\n", time.Time(order.UpdatedAt).Format("2006-01-02 15:04:05"))
	logger.Printf("AddressBilling FirstName: %s\n", order.AddressBilling.FirstName)
	logger.Printf("AddressBilling LastName: %s\n", order.AddressBilling.LastName)
	logger.Printf("AddressBilling Phone: %s\n", order.AddressBilling.Phone)
	logger.Printf("AddressBilling Phone2: %s\n", order.AddressBilling.Phone2)
	logger.Printf("AddressBilling Address1: %s\n", order.AddressBilling.Address1)
	logger.Printf("AddressBilling Address2: %s\n", order.AddressBilling.Address2)
	logger.Printf("AddressBilling Address3: %s\n", order.AddressBilling.Address3)
	logger.Printf("AddressBilling Address4: %s\n", order.AddressBilling.Address4)
	logger.Printf("AddressBilling Address5: %s\n", order.AddressBilling.Address5)
	logger.Printf("AddressBilling City: %s\n", order.AddressBilling.City)
	logger.Printf("AddressBilling Ward: %s\n", order.AddressBilling.Ward)
	logger.Printf("AddressBilling Region: %s\n", order.AddressBilling.Region)
	logger.Printf("AddressBilling PostCode: %s\n", order.AddressBilling.PostCode)
	logger.Printf("AddressBilling Country: %s\n", order.AddressBilling.Country)
	logger.Printf("AddressShipping FirstName: %s\n", order.AddressShipping.FirstName)
	logger.Printf("AddressShipping LastName: %s\n", order.AddressShipping.LastName)
	logger.Printf("AddressShipping Phone: %s\n", order.AddressShipping.Phone)
	logger.Printf("AddressShipping Phone2: %s\n", order.AddressShipping.Phone2)
	logger.Printf("AddressShipping Address1: %s\n", order.AddressShipping.Address1)
	logger.Printf("AddressShipping Address2: %s\n", order.AddressShipping.Address2)
	logger.Printf("AddressShipping Address3: %s\n", order.AddressShipping.Address3)
	logger.Printf("AddressShipping Address4: %s\n", order.AddressShipping.Address4)
	logger.Printf("AddressShipping Address5: %s\n", order.AddressShipping.Address5)
	logger.Printf("AddressShipping City: %s\n", order.AddressShipping.City)
	logger.Printf("AddressShipping Ward: %s\n", order.AddressShipping.Ward)
	logger.Printf("AddressShipping Region: %s\n", order.AddressShipping.Region)
	logger.Printf("AddressShipping PostCode: %s\n", order.AddressShipping.PostCode)
	logger.Printf("AddressShipping Country: %s\n", order.AddressShipping.Country)
	logger.Printf("NationalRegistrationNumber: %s\n", order.NationalRegistrationNumber)
	logger.Printf("ItemsCount: %d\n", order.ItemsCount)
	logger.Printf("PromisedShippingTime: %s\n", time.Time(order.PromisedShippingTime).Format("2006-01-02 15:04:05"))
	logger.Printf("ExtraAttributes: %s\n", order.ExtraAttributes)
	logger.Printf("Status: %s\n", order.Statuses.Status)
}
