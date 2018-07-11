package model

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
	Orders []Order `json:"Order"`
}

type Status struct {
	Status string `json:"Status"`
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
	OrderId                    scInt       `json:"OrderId"`
	CustomerFirstName          string      `json:"CustomerFirstName"`
	CustomerLastName           string      `json:"CustomerLastName"`
	OrderNumber                string      `json:"OrderNumber"`
	PaymentMethod              string      `json:"PaymentMethod"`
	Remarks                    string      `json:"Remarks"`
	DeliveryInfo               string      `json:"DeliveryInfo"`
	Price                      scFloat     `json:"Price"`
	GiftOption                 scBool      `json:"GiftOption"`
	GiftMessage                string      `json:"GiftMessage"`
	VoucherCode                string      `json:"VoucherCode"`
	CreatedAt                  scTimestamp `json:"CreatedAt"`
	UpdatedAt                  scTimestamp `json:"UpdatedAt"`
	AddressBilling             Address     `json:"AddressBilling"`
	AddressShipping            Address     `json:"AddressShipping"`
	NationalRegistrationNumber string      `json:"NationalRegistrationNumber"`
	ItemsCount                 scInt       `json:"ItemsCount"`
	PromisedShippingTime       scTimestamp `json:"PromisedShippingTime"`
	ExtraAttributes            string      `json:"ExtraAttributes"`
	Statuses                   Status      `json:"Statuses"`
}

type OrderItems struct {
	OrderItems []OrderItem `json:"OrderItem"`
}

type OrderItem struct {
	OrderItemId          scInt       `json:"OrderItemId"`
	ShopId               string      `json:"ShopId"`
	OrderId              scInt       `json:"OrderId"`
	Name                 string      `json:"Name"`
	Sku                  string      `json:"Sku"`
	Variation            string      `json:"Variation"`
	ShopSku              string      `json:"ShopSku"`
	ShippingType         string      `json:"ShippingType"`
	ItemPrice            scFloat     `json:"ItemPrice"`
	PaidPrice            scFloat     `json:"PaidPrice"`
	Currency             string      `json:"Currency"`
	WalletCredits        scFloat     `json:"WalletCredits"`
	TaxAmount            scFloat     `json:"TaxAmount"`
	CodCollectableAmount scFloat     `json:"CodCollectableAmount"`
	ShippingAmount       scFloat     `json:"ShippingAmount"`
	ShippingServiceCost  scFloat     `json:"ShippingServiceCost"`
	VoucherAmount        scFloat     `json:"VoucherAmount"`
	VoucherCode          string      `json:"VoucherCode"`
	Status               string      `json:"Status"`
	IsProcessable        scBool      `json:"IsProcessable"`
	ShipmentProvider     string      `json:"ShipmentProvider"`
	IsDigital            scBool      `json:"IsDigital"`
	DigitalDeliveryInfo  string      `json:"DigitalDeliveryInfo"`
	TrackingCode         string      `json:"TrackingCode"`
	TrackingCodePre      string      `json:"TrackingCodePre"`
	Reason               string      `json:"Reason"`
	ReasonDetail         string      `json:"ReasonDetail"`
	PurchaseOrderId      scInt       `json:"PurchaseOrderId"`
	PurchaseOrderNumber  string      `json:"PurchaseOrderNumber"`
	PackageId            string      `json:"PackageId"`
	PromisedShippingTime scTimestamp `json:"PromisedShippingTime"`
	ExtraAttributes      string      `json:"ExtraAttributes"`
	ShippingProviderType string      `json:"ShippingProviderType"`
	CreatedAt            scTimestamp `json:"CreatedAt"`
	UpdatedAt            scTimestamp `json:"UpdatedAt"`
	ReturnStatus         string      `json:"ReturnStatus"`
}

type Document struct {
	DocumentType string `json:"DocumentType"`
	MimeType     string `json:"MimeType"`
	File         string `json:"File"`
}
