package model

import (
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
)

type DeliveryType string
type DocumentType string

const (
	DeliveryTypeCrossdocking = DeliveryType("send_to_warehouse")
	DeliveryTypeDropshipping = DeliveryType("dropship")
	DeliveryTypePickup       = DeliveryType("pickup")

	DocumentTypeCarrierManifest = DocumentType("carrierManifest")
	DocumentTypeSerialNumber    = DocumentType("serialNumber")
	DocumentTypeExportInvoice   = DocumentType("exportInvoice")
	DocumentTypeInvoice         = DocumentType("invoice")
	DocumentTypeShippingLabel   = DocumentType("shippingLabel")
	DocumentTypeShippingParcel  = DocumentType("shippingParcel")

	scTimeFormat = "2006-01-02 15:04:05"
)

type Orders struct {
	Orders []Order `json:"Orders"`
}

func (o *Orders) UnmarshalJSON(b []byte) error {
	rawOrders, dataType, _, err := jsonparser.Get(b, "Orders", "Order")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawOrders) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot find order")
	}

	var orders []Order
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawOrders, &orders); nil != err {
			return err
		}
	case jsonparser.Object:
		var order Order
		if err := json.Unmarshal(rawOrders, &order); nil != err {
			return err
		}

		orders = []Order{order}
	}

	*o = Orders{orders}

	return nil
}

type Status []string

func (s *Status) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	raw, dataType, _, err := jsonparser.Get(b, "Status")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(raw) == 0 || dataType == jsonparser.NotExist {
		return nil
	}

	var status []string
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(raw, &status); nil != err {
			return err
		}
	case jsonparser.String:
		status = Status{string(raw)}
	}

	*s = status

	return nil
}

type Address struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Phone     string `json:"Phone"`
	Phone2    string `json:"Phone2"`
	Address1  string `json:"Address1"`
	Address2  string `json:"Address2"`
	Address3  string `json:"Address3"`
	Address4  string `json:"Address4"`
	Address5  string `json:"Address5"`
	City      string `json:"City"`
	Ward      string `json:"Ward"`
	Region    string `json:"Region"`
	PostCode  string `json:"PostCode"`
	Country   string `json:"Country"`
}

type Order struct {
	OrderId                    ScInt       `json:"OrderId"`
	CustomerFirstName          string      `json:"CustomerFirstName"`
	CustomerLastName           string      `json:"CustomerLastName"`
	OrderNumber                string      `json:"OrderNumber"`
	PaymentMethod              string      `json:"PaymentMethod"`
	Remarks                    string      `json:"Remarks"`
	DeliveryInfo               string      `json:"DeliveryInfo"`
	Price                      ScFloat     `json:"Price"`
	GiftOption                 ScBool      `json:"GiftOption"`
	GiftMessage                string      `json:"GiftMessage"`
	VoucherCode                string      `json:"VoucherCode"`
	CreatedAt                  ScTimestamp `json:"CreatedAt"`
	UpdatedAt                  ScTimestamp `json:"UpdatedAt"`
	AddressBilling             Address     `json:"AddressBilling"`
	AddressShipping            Address     `json:"AddressShipping"`
	NationalRegistrationNumber string      `json:"NationalRegistrationNumber"`
	ItemsCount                 ScInt       `json:"ItemsCount"`
	PromisedShippingTime       ScTimestamp `json:"PromisedShippingTime"`
	ExtraAttributes            string      `json:"ExtraAttributes"`
	Statuses                   Status      `json:"Statuses"`
}

type OrdersWithItems struct {
	Orders []OrderWithItems `json:"Orders"`
}

func (o *OrdersWithItems) UnmarshalJSON(b []byte) error {
	rawOrders, dataType, _, err := jsonparser.Get(b, "Orders", "Order")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawOrders) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot find order")
	}

	var orders []OrderWithItems
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawOrders, &orders); nil != err {
			return err
		}
	case jsonparser.Object:
		var order OrderWithItems
		if err := json.Unmarshal(rawOrders, &order); nil != err {
			return err
		}

		orders = []OrderWithItems{order}
	}

	*o = OrdersWithItems{orders}

	return nil
}

type OrderWithItems struct {
	OrderId     ScInt                      `json:"OrderId"`
	OrderNumber string                     `json:"OrderNumber"`
	OrderItems  OrderItemsInOrderWithItems `json:"OrderItems"`
}

type OrderItems struct {
	OrderItems []OrderItem `json:"OrderItems"`
}

func (oi *OrderItems) UnmarshalJSON(b []byte) error {
	rawOrderItems, dataType, _, err := jsonparser.Get(b, "OrderItems", "OrderItem")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawOrderItems) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot find order items")
	}

	var orderItems []OrderItem
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawOrderItems, &orderItems); nil != err {
			return err
		}
	case jsonparser.Object:
		var orderItem OrderItem
		if err := json.Unmarshal(rawOrderItems, &orderItem); nil != err {
			return err
		}

		orderItems = []OrderItem{orderItem}
	}

	*oi = OrderItems{orderItems}

	return nil
}

type OrderItemsInOrderWithItems struct {
	Items []OrderItem `json:"OrderItems"`
}

func (oi *OrderItemsInOrderWithItems) UnmarshalJSON(b []byte) error {
	rawOrderItems, dataType, _, err := jsonparser.Get(b, "OrderItem")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawOrderItems) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot find order items")
	}

	var orderItems []OrderItem
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawOrderItems, &orderItems); nil != err {
			return err
		}
	case jsonparser.Object:
		var orderItem OrderItem
		if err := json.Unmarshal(rawOrderItems, &orderItem); nil != err {
			return err
		}

		orderItems = []OrderItem{orderItem}
	}

	*oi = OrderItemsInOrderWithItems{orderItems}

	return nil
}

type OrderItem struct {
	OrderItemId          ScInt       `json:"OrderItemId"`
	ShopId               string      `json:"ShopId"`
	OrderId              ScInt       `json:"OrderId"`
	Name                 string      `json:"Name"`
	Sku                  string      `json:"Sku"`
	Variation            string      `json:"Variation"`
	ShopSku              string      `json:"ShopSku"`
	ShippingType         string      `json:"ShippingType"`
	ItemPrice            ScFloat     `json:"ItemPrice"`
	PaidPrice            ScFloat     `json:"PaidPrice"`
	Currency             string      `json:"Currency"`
	WalletCredits        ScFloat     `json:"WalletCredits"`
	TaxAmount            ScFloat     `json:"TaxAmount"`
	CodCollectableAmount ScFloat     `json:"CodCollectableAmount"`
	ShippingAmount       ScFloat     `json:"ShippingAmount"`
	ShippingServiceCost  ScFloat     `json:"ShippingServiceCost"`
	VoucherAmount        ScFloat     `json:"VoucherAmount"`
	VoucherCode          string      `json:"VoucherCode"`
	Status               string      `json:"Status"`
	IsProcessable        ScBool      `json:"IsProcessable"`
	ShipmentProvider     string      `json:"ShipmentProvider"`
	IsDigital            ScBool      `json:"IsDigital"`
	DigitalDeliveryInfo  string      `json:"DigitalDeliveryInfo"`
	TrackingCode         string      `json:"TrackingCode"`
	TrackingCodePre      string      `json:"TrackingCodePre"`
	Reason               string      `json:"Reason"`
	ReasonDetail         string      `json:"ReasonDetail"`
	PurchaseOrderId      ScInt       `json:"PurchaseOrderId"`
	PurchaseOrderNumber  string      `json:"PurchaseOrderNumber"`
	PackageId            string      `json:"PackageId"`
	PromisedShippingTime ScTimestamp `json:"PromisedShippingTime"`
	ExtraAttributes      string      `json:"ExtraAttributes"`
	ShippingProviderType string      `json:"ShippingProviderType"`
	CreatedAt            ScTimestamp `json:"CreatedAt"`
	UpdatedAt            ScTimestamp `json:"UpdatedAt"`
	ReturnStatus         string      `json:"ReturnStatus"`
}

type Document struct {
	DocumentType string `json:"DocumentType"`
	MimeType     string `json:"MimeType"`
	File         string `json:"File"`
}
