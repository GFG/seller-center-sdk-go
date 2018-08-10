package model

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func Test_OrdersEmpty(t *testing.T) {
	j := []byte("{}")

	expected := Orders{[]Order{}}

	var c Orders
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrdersSingle(t *testing.T) {
	j := []byte("{\"Orders\":{\"Order\":{\"OrderId\":\"1\",\"CustomerFirstName\":\"CustomerFirstName 1\",\"CustomerLastName\":\"CustomerLastName 1\",\"OrderNumber\":\"01\",\"PaymentMethod\":\"CashOnDelivery 1\",\"Remarks\":\"Remarks 1\",\"DeliveryInfo\":\"DeliveryInfo 1\",\"Price\":\"380.00\",\"GiftOption\":\"0\",\"GiftMessage\":\"GiftMessage 1\",\"VoucherCode\":\"VoucherCode 1\",\"CreatedAt\":\"2015-11-04 10:30:49\",\"UpdatedAt\":\"2015-11-05 10:30:49\",\"AddressUpdatedAt\":\"2015-11-06 10:30:49\",\"AddressBilling\":{\"FirstName\":\"FirstName 1\",\"LastName\":\"LastName 1\",\"Phone\":\"00111000\",\"Phone2\":\"00222000\",\"Address1\":\"Address1 1\",\"Address2\":\"Address2 1\",\"Address3\":\"Address3 1\",\"Address4\":\"Address4 1\",\"Address5\":\"Address5 1\",\"CustomerEmail\":\"CustomerEmail 1\",\"City\":\"City 1\",\"Ward\":\"Ward 1\",\"Region\":\"Region 1\",\"PostCode\":\"000001\",\"Country\":\"Country 1\" },\"AddressShipping\":{\"FirstName\":\"FirstName 2\",\"LastName\":\"LastName 2\",\"Phone\":\"00333000\",\"Phone2\":\"00444000\",\"Address1\":\"Address1 2\",\"Address2\":\"Address2 2\",\"Address3\":\"Address3 2\",\"Address4\":\"Address4 2\",\"Address5\":\"Address5 2\",\"CustomerEmail\":\"CustomerEmail 2\",\"City\":\"City 2\",\"Ward\":\"Ward 2\",\"Region\":\"Region 2\",\"PostCode\":\"000002\",\"Country\":\"Country 2\" },\"NationalRegistrationNumber\":\"NationalRegistrationNumber 1\",\"ItemsCount\":\"1\",\"PromisedShippingTime\":\"2015-11-07 10:30:49\",\"ExtraAttributes\":\"ExtraAttributes 1\",\"Statuses\":{\"Status\":[\"ready_to_ship\",\"shipped\"] } } } }")

	expected := Orders{[]Order{
		{
			ScInt(1),
			"CustomerFirstName 1",
			"CustomerLastName 1",
			"01",
			"CashOnDelivery 1",
			"Remarks 1",
			"DeliveryInfo 1",
			ScFloat(380.0),
			ScBool(false),
			"GiftMessage 1",
			"VoucherCode 1",
			ScTimestamp(time.Date(2015, 11, 4, 10, 30, 49, 00, time.UTC)),
			ScTimestamp(time.Date(2015, 11, 5, 10, 30, 49, 00, time.UTC)),
			Address{
				"FirstName 1",
				"LastName 1",
				"00111000",
				"00222000",
				"Address1 1",
				"Address2 1",
				"Address3 1",
				"Address4 1",
				"Address5 1",
				"City 1",
				"Ward 1",
				"Region 1",
				"000001",
				"Country 1",
			},
			Address{
				"FirstName 2",
				"LastName 2",
				"00333000",
				"00444000",
				"Address1 2",
				"Address2 2",
				"Address3 2",
				"Address4 2",
				"Address5 2",
				"City 2",
				"Ward 2",
				"Region 2",
				"000002",
				"Country 2",
			},
			"NationalRegistrationNumber 1",
			ScInt(1),
			ScTimestamp(time.Date(2015, 11, 7, 10, 30, 49, 00, time.UTC)),
			"ExtraAttributes 1",
			Status{"ready_to_ship", "shipped"},
		},
	},
	}

	var c Orders
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrdersMultiple(t *testing.T) {
	j := []byte("{\"Orders\":{\"Order\":[{\"OrderId\":\"1\",\"CustomerFirstName\":\"CustomerFirstName 1\",\"CustomerLastName\":\"CustomerLastName 1\",\"OrderNumber\":\"01\",\"PaymentMethod\":\"CashOnDelivery 1\",\"Remarks\":\"Remarks 1\",\"DeliveryInfo\":\"DeliveryInfo 1\",\"Price\":\"380.00\",\"GiftOption\":\"0\",\"GiftMessage\":\"GiftMessage 1\",\"VoucherCode\":\"VoucherCode 1\",\"CreatedAt\":\"2015-11-04 10:30:49\",\"UpdatedAt\":\"2015-11-05 10:30:49\",\"AddressUpdatedAt\":\"2015-11-06 10:30:49\",\"AddressBilling\":{\"FirstName\":\"FirstName 1\",\"LastName\":\"LastName 1\",\"Phone\":\"00111000\",\"Phone2\":\"00222000\",\"Address1\":\"Address1 1\",\"Address2\":\"Address2 1\",\"Address3\":\"Address3 1\",\"Address4\":\"Address4 1\",\"Address5\":\"Address5 1\",\"CustomerEmail\":\"CustomerEmail 1\",\"City\":\"City 1\",\"Ward\":\"Ward 1\",\"Region\":\"Region 1\",\"PostCode\":\"000001\",\"Country\":\"Country 1\" },\"AddressShipping\":{\"FirstName\":\"FirstName 2\",\"LastName\":\"LastName 2\",\"Phone\":\"00333000\",\"Phone2\":\"00444000\",\"Address1\":\"Address1 2\",\"Address2\":\"Address2 2\",\"Address3\":\"Address3 2\",\"Address4\":\"Address4 2\",\"Address5\":\"Address5 2\",\"CustomerEmail\":\"CustomerEmail 2\",\"City\":\"City 2\",\"Ward\":\"Ward 2\",\"Region\":\"Region 2\",\"PostCode\":\"000002\",\"Country\":\"Country 2\" },\"NationalRegistrationNumber\":\"NationalRegistrationNumber 1\",\"ItemsCount\":\"1\",\"PromisedShippingTime\":\"2015-11-07 10:30:49\",\"ExtraAttributes\":\"ExtraAttributes 1\",\"Statuses\":{\"Status\":\"shipped\" } },{\"OrderId\":\"2\",\"CustomerFirstName\":\"CustomerFirstName 2\",\"CustomerLastName\":\"CustomerLastName 2\",\"OrderNumber\":\"02\",\"PaymentMethod\":\"CashOnDelivery 2\",\"Remarks\":\"Remarks 2\",\"DeliveryInfo\":\"DeliveryInfo 2\",\"Price\":\"75.00\",\"GiftOption\":\"1\",\"GiftMessage\":\"GiftMessage 2\",\"VoucherCode\":\"VoucherCode 2\",\"CreatedAt\":\"2016-11-04 10:30:49\",\"UpdatedAt\":\"2016-11-05 10:30:49\",\"AddressUpdatedAt\":\"2016-11-06 10:30:49\",\"AddressBilling\":{\"FirstName\":\"FirstName 3\",\"LastName\":\"LastName 3\",\"Phone\":\"00555000\",\"Phone2\":\"00666000\",\"Address1\":\"Address1 3\",\"Address2\":\"Address2 3\",\"Address3\":\"Address3 3\",\"Address4\":\"Address4 3\",\"Address5\":\"Address5 3\",\"CustomerEmail\":\"CustomerEmail 3\",\"City\":\"City 3\",\"Ward\":\"Ward 3\",\"Region\":\"Region 3\",\"PostCode\":\"000003\",\"Country\":\"Country 3\" },\"AddressShipping\":{\"FirstName\":\"FirstName 4\",\"LastName\":\"LastName 4\",\"Phone\":\"00777000\",\"Phone2\":\"00888000\",\"Address1\":\"Address1 4\",\"Address2\":\"Address2 4\",\"Address3\":\"Address3 4\",\"Address4\":\"Address4 4\",\"Address5\":\"Address5 4\",\"CustomerEmail\":\"CustomerEmail 4\",\"City\":\"City 4\",\"Ward\":\"Ward 4\",\"Region\":\"Region 4\",\"PostCode\":\"000004\",\"Country\":\"Country 4\" },\"NationalRegistrationNumber\":\"NationalRegistrationNumber 2\",\"ItemsCount\":\"2\",\"PromisedShippingTime\":\"2016-11-07 10:30:49\",\"ExtraAttributes\":\"ExtraAttributes 2\",\"Statuses\":{\"Status\":\"pending\" } } ] } }")

	expected := Orders{[]Order{
		{
			ScInt(1),
			"CustomerFirstName 1",
			"CustomerLastName 1",
			"01",
			"CashOnDelivery 1",
			"Remarks 1",
			"DeliveryInfo 1",
			ScFloat(380.0),
			ScBool(false),
			"GiftMessage 1",
			"VoucherCode 1",
			ScTimestamp(time.Date(2015, 11, 4, 10, 30, 49, 00, time.UTC)),
			ScTimestamp(time.Date(2015, 11, 5, 10, 30, 49, 00, time.UTC)),
			Address{
				"FirstName 1",
				"LastName 1",
				"00111000",
				"00222000",
				"Address1 1",
				"Address2 1",
				"Address3 1",
				"Address4 1",
				"Address5 1",
				"City 1",
				"Ward 1",
				"Region 1",
				"000001",
				"Country 1",
			},
			Address{
				"FirstName 2",
				"LastName 2",
				"00333000",
				"00444000",
				"Address1 2",
				"Address2 2",
				"Address3 2",
				"Address4 2",
				"Address5 2",
				"City 2",
				"Ward 2",
				"Region 2",
				"000002",
				"Country 2",
			},
			"NationalRegistrationNumber 1",
			ScInt(1),
			ScTimestamp(time.Date(2015, 11, 7, 10, 30, 49, 00, time.UTC)),
			"ExtraAttributes 1",
			Status{"shipped"},
		},
		{
			ScInt(2),
			"CustomerFirstName 2",
			"CustomerLastName 2",
			"02",
			"CashOnDelivery 2",
			"Remarks 2",
			"DeliveryInfo 2",
			ScFloat(75.0),
			ScBool(true),
			"GiftMessage 2",
			"VoucherCode 2",
			ScTimestamp(time.Date(2016, 11, 4, 10, 30, 49, 00, time.UTC)),
			ScTimestamp(time.Date(2016, 11, 5, 10, 30, 49, 00, time.UTC)),
			Address{
				"FirstName 3",
				"LastName 3",
				"00555000",
				"00666000",
				"Address1 3",
				"Address2 3",
				"Address3 3",
				"Address4 3",
				"Address5 3",
				"City 3",
				"Ward 3",
				"Region 3",
				"000003",
				"Country 3",
			},
			Address{
				"FirstName 4",
				"LastName 4",
				"00777000",
				"00888000",
				"Address1 4",
				"Address2 4",
				"Address3 4",
				"Address4 4",
				"Address5 4",
				"City 4",
				"Ward 4",
				"Region 4",
				"000004",
				"Country 4",
			},
			"NationalRegistrationNumber 2",
			ScInt(2),
			ScTimestamp(time.Date(2016, 11, 7, 10, 30, 49, 00, time.UTC)),
			"ExtraAttributes 2",
			Status{"pending"},
		},
	},
	}

	var c Orders
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrdersWithItemsEmpty(t *testing.T) {
	j := []byte("{}")

	expected := OrdersWithItems{[]OrderWithItems{}}

	var c OrdersWithItems
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrdersWithItemsSingle(t *testing.T) {
	j := []byte("{\"Orders\":{\"Order\":{\"OrderId\":\"1\",\"OrderNumber\":\"01\", \"OrderItems\":{ \"OrderItem\": { \"OrderItemId\": \"1\",	\"ShopId\": \"ShopId 1\",	\"OrderId\": \"1\",	\"Name\": \"Name 1\",	\"Sku\": \"Sku 1\",	\"Variation\": \"Variation 1\",	\"ShopSku\": \"ShopSku 1\",	\"ShippingType\": \"Dropshipping 1\",	\"ItemPrice\": \"180.00\",	\"PaidPrice\": \"280.00\",	\"Currency\": \"USD\",	\"WalletCredits\": \"380.00\",	\"TaxAmount\": \"18.32\",	\"CodCollectableAmount\": \"19.32\",	\"ShippingAmount\": \"20.32\",	\"ShippingServiceCost\": \"21.32\",	\"VoucherAmount\": \"22.32\",	\"VoucherCode\": \"VoucherCode 1\",	\"Status\": \"shipped\",	\"IsProcessable\": \"1\",	\"ShipmentProvider\": \"DHL\",	\"IsDigital\": \"0\",	\"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 1\",	\"TrackingCode\": \"TrackingCode 1\",	\"TrackingCodePre\": \"TrackingCodePre 1\",	\"Reason\": \"Reason 1\",	\"ReasonDetail\": \"ReasonDetail 1\",	\"PurchaseOrderId\": \"1\",	\"PurchaseOrderNumber\": \"PurchaseOrderNumber 1\",	\"PackageId\": \"PackageId 1\",	\"PromisedShippingTime\": \"2015-11-04 10:30:57\",	\"ExtraAttributes\": \"ExtraAttributes 1\",	\"ShippingProviderType\": \"ShippingProviderType 1\",	\"CreatedAt\": \"2015-11-05 10:30:57\",	\"UpdatedAt\": \"2015-11-06 10:30:57\",	\"ReturnStatus\": \"ReturnStatus 1\"} } } } }")

	expected := OrdersWithItems{[]OrderWithItems{
		{
			ScInt(1),
			"01",
			OrderItems{
				[]OrderItem{
					{
						ScInt(1),
						"ShopId 1",
						ScInt(1),
						"Name 1",
						"Sku 1",
						"Variation 1",
						"ShopSku 1",
						"Dropshipping 1",
						ScFloat(180.00),
						ScFloat(280.00),
						"USD",
						ScFloat(380.00),
						ScFloat(18.32),
						ScFloat(19.32),
						ScFloat(20.32),
						ScFloat(21.32),
						ScFloat(22.32),
						"VoucherCode 1",
						"shipped",
						ScBool(true),
						"DHL",
						ScBool(false),
						"DigitalDeliveryInfo 1",
						"TrackingCode 1",
						"TrackingCodePre 1",
						"Reason 1",
						"ReasonDetail 1",
						ScInt(1),
						"PurchaseOrderNumber 1",
						"PackageId 1",
						ScTimestamp(time.Date(2015, 11, 4, 10, 30, 57, 00, time.UTC)),
						"ExtraAttributes 1",
						"ShippingProviderType 1",
						ScTimestamp(time.Date(2015, 11, 5, 10, 30, 57, 00, time.UTC)),
						ScTimestamp(time.Date(2015, 11, 6, 10, 30, 57, 00, time.UTC)),
						"ReturnStatus 1",
					},
				},
			},
		},
	},
	}

	var c OrdersWithItems
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrdersWithItemsMultiple(t *testing.T) {
	j := []byte("{\"Orders\":{\"Order\":[ {\"OrderId\":\"1\",\"OrderNumber\":\"01\", \"OrderItems\":{ \"OrderItem\": { \"OrderItemId\": \"1\", \"ShopId\": \"ShopId 1\", \"OrderId\": \"1\", \"Name\": \"Name 1\", \"Sku\": \"Sku 1\", \"Variation\": \"Variation 1\", \"ShopSku\": \"ShopSku 1\", \"ShippingType\": \"Dropshipping 1\", \"ItemPrice\": \"180.00\", \"PaidPrice\": \"280.00\", \"Currency\": \"USD\", \"WalletCredits\": \"380.00\", \"TaxAmount\": \"18.32\", \"CodCollectableAmount\": \"19.32\", \"ShippingAmount\": \"20.32\", \"ShippingServiceCost\": \"21.32\", \"VoucherAmount\": \"22.32\", \"VoucherCode\": \"VoucherCode 1\", \"Status\": \"shipped\", \"IsProcessable\": \"0\", \"ShipmentProvider\": \"DHL 1\", \"IsDigital\": \"1\", \"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 1\", \"TrackingCode\": \"TrackingCode 1\", \"TrackingCodePre\": \"TrackingCodePre 1\", \"Reason\": \"Reason 1\", \"ReasonDetail\": \"ReasonDetail 1\", \"PurchaseOrderId\": \"1\", \"PurchaseOrderNumber\": \"PurchaseOrderNumber 1\", \"PackageId\": \"PackageId 1\", \"PromisedShippingTime\": \"2015-11-04 10:30:57\", \"ExtraAttributes\": \"ExtraAttributes 1\", \"ShippingProviderType\": \"ShippingProviderType 1\", \"CreatedAt\": \"2015-11-05 10:30:57\", \"UpdatedAt\": \"2015-11-06 10:30:57\", \"ReturnStatus\": \"ReturnStatus 1\"} } }, {\"OrderId\":\"2\",\"OrderNumber\":\"02\", \"OrderItems\":{ \"OrderItem\": [{ \"OrderItemId\": \"2\", \"ShopId\": \"ShopId 2-1\", \"OrderId\": \"2\", \"Name\": \"Name 2-1\", \"Sku\": \"Sku 2-1\", \"Variation\": \"Variation 2-1\", \"ShopSku\": \"ShopSku 2-1\", \"ShippingType\": \"Dropshipping 2-1\", \"ItemPrice\": \"21180.00\", \"PaidPrice\": \"21280.00\", \"Currency\": \"EUR\", \"WalletCredits\": \"21380.00\", \"TaxAmount\": \"2118.32\", \"CodCollectableAmount\": \"2119.32\", \"ShippingAmount\": \"2120.32\", \"ShippingServiceCost\": \"2121.32\", \"VoucherAmount\": \"2122.32\", \"VoucherCode\": \"VoucherCode 2-1\", \"Status\": \"returned\", \"IsProcessable\": \"1\", \"ShipmentProvider\": \"DHL 2-1\", \"IsDigital\": \"0\", \"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 2-1\", \"TrackingCode\": \"TrackingCode 2-1\", \"TrackingCodePre\": \"TrackingCodePre 2-1\", \"Reason\": \"Reason 2-1\", \"ReasonDetail\": \"ReasonDetail 2-1\", \"PurchaseOrderId\": \"2\", \"PurchaseOrderNumber\": \"PurchaseOrderNumber 2-1\", \"PackageId\": \"PackageId 2-1\", \"PromisedShippingTime\": \"2016-11-04 10:30:57\", \"ExtraAttributes\": \"ExtraAttributes 2-1\", \"ShippingProviderType\": \"ShippingProviderType 2-1\", \"CreatedAt\": \"2016-11-05 10:30:57\", \"UpdatedAt\": \"2016-11-06 10:30:57\", \"ReturnStatus\": \"ReturnStatus 2-1\"}, { \"OrderItemId\": \"3\", \"ShopId\": \"ShopId 2-2\", \"OrderId\": \"2\", \"Name\": \"Name 2-2\", \"Sku\": \"Sku 2-2\", \"Variation\": \"Variation 2-2\", \"ShopSku\": \"ShopSku 2-2\", \"ShippingType\": \"Dropshipping 2-2\", \"ItemPrice\": \"22180.00\", \"PaidPrice\": \"22280.00\", \"Currency\": \"KRW\", \"WalletCredits\": \"22380.00\", \"TaxAmount\": \"2218.32\", \"CodCollectableAmount\": \"2219.32\", \"ShippingAmount\": \"2220.32\", \"ShippingServiceCost\": \"2221.32\", \"VoucherAmount\": \"2222.32\", \"VoucherCode\": \"VoucherCode 2-2\", \"Status\": \"canceled\", \"IsProcessable\": \"0\", \"ShipmentProvider\": \"DHL 2-2\", \"IsDigital\": \"1\", \"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 2-2\", \"TrackingCode\": \"TrackingCode 2-2\", \"TrackingCodePre\": \"TrackingCodePre 2-2\", \"Reason\": \"Reason 2-2\", \"ReasonDetail\": \"ReasonDetail 2-2\", \"PurchaseOrderId\": \"3\", \"PurchaseOrderNumber\": \"PurchaseOrderNumber 2-2\", \"PackageId\": \"PackageId 2-2\", \"PromisedShippingTime\": \"2017-11-04 10:30:57\", \"ExtraAttributes\": \"ExtraAttributes 2-2\", \"ShippingProviderType\": \"ShippingProviderType 2-2\", \"CreatedAt\": \"2017-11-05 10:30:57\", \"UpdatedAt\": \"2017-11-06 10:30:57\", \"ReturnStatus\": \"ReturnStatus 2-2\"} ] } } ] } }")

	expected := OrdersWithItems{[]OrderWithItems{
		{
			ScInt(1),
			"01",
			OrderItems{
				[]OrderItem{
					{
						ScInt(1),
						"ShopId 1",
						ScInt(1),
						"Name 1",
						"Sku 1",
						"Variation 1",
						"ShopSku 1",
						"Dropshipping 1",
						ScFloat(180.00),
						ScFloat(280.00),
						"USD",
						ScFloat(380.00),
						ScFloat(18.32),
						ScFloat(19.32),
						ScFloat(20.32),
						ScFloat(21.32),
						ScFloat(22.32),
						"VoucherCode 1",
						"shipped",
						ScBool(false),
						"DHL 1",
						ScBool(true),
						"DigitalDeliveryInfo 1",
						"TrackingCode 1",
						"TrackingCodePre 1",
						"Reason 1",
						"ReasonDetail 1",
						ScInt(1),
						"PurchaseOrderNumber 1",
						"PackageId 1",
						ScTimestamp(time.Date(2015, 11, 4, 10, 30, 57, 00, time.UTC)),
						"ExtraAttributes 1",
						"ShippingProviderType 1",
						ScTimestamp(time.Date(2015, 11, 5, 10, 30, 57, 00, time.UTC)),
						ScTimestamp(time.Date(2015, 11, 6, 10, 30, 57, 00, time.UTC)),
						"ReturnStatus 1",
					},
				},
			},
		}, {
			ScInt(2),
			"02",
			OrderItems{
				[]OrderItem{
					{
						ScInt(2),
						"ShopId 2-1",
						ScInt(2),
						"Name 2-1",
						"Sku 2-1",
						"Variation 2-1",
						"ShopSku 2-1",
						"Dropshipping 2-1",
						ScFloat(21180.00),
						ScFloat(21280.00),
						"EUR",
						ScFloat(21380.00),
						ScFloat(2118.32),
						ScFloat(2119.32),
						ScFloat(2120.32),
						ScFloat(2121.32),
						ScFloat(2122.32),
						"VoucherCode 2-1",
						"returned",
						ScBool(true),
						"DHL 2-1",
						ScBool(false),
						"DigitalDeliveryInfo 2-1",
						"TrackingCode 2-1",
						"TrackingCodePre 2-1",
						"Reason 2-1",
						"ReasonDetail 2-1",
						ScInt(2),
						"PurchaseOrderNumber 2-1",
						"PackageId 2-1",
						ScTimestamp(time.Date(2016, 11, 4, 10, 30, 57, 00, time.UTC)),
						"ExtraAttributes 2-1",
						"ShippingProviderType 2-1",
						ScTimestamp(time.Date(2016, 11, 5, 10, 30, 57, 00, time.UTC)),
						ScTimestamp(time.Date(2016, 11, 6, 10, 30, 57, 00, time.UTC)),
						"ReturnStatus 2-1",
					}, {
						ScInt(3),
						"ShopId 2-2",
						ScInt(2),
						"Name 2-2",
						"Sku 2-2",
						"Variation 2-2",
						"ShopSku 2-2",
						"Dropshipping 2-2",
						ScFloat(22180.00),
						ScFloat(22280.00),
						"KRW",
						ScFloat(22380.00),
						ScFloat(2218.32),
						ScFloat(2219.32),
						ScFloat(2220.32),
						ScFloat(2221.32),
						ScFloat(2222.32),
						"VoucherCode 2-2",
						"canceled",
						ScBool(false),
						"DHL 2-2",
						ScBool(true),
						"DigitalDeliveryInfo 2-2",
						"TrackingCode 2-2",
						"TrackingCodePre 2-2",
						"Reason 2-2",
						"ReasonDetail 2-2",
						ScInt(3),
						"PurchaseOrderNumber 2-2",
						"PackageId 2-2",
						ScTimestamp(time.Date(2017, 11, 4, 10, 30, 57, 00, time.UTC)),
						"ExtraAttributes 2-2",
						"ShippingProviderType 2-2",
						ScTimestamp(time.Date(2017, 11, 5, 10, 30, 57, 00, time.UTC)),
						ScTimestamp(time.Date(2017, 11, 6, 10, 30, 57, 00, time.UTC)),
						"ReturnStatus 2-2",
					},
				},
			},
		},
	},
	}

	var c OrdersWithItems
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}
func Test_OrderItemsEmpty(t *testing.T) {
	j := []byte("{}")

	expected := OrderItems{[]OrderItem{}}

	var c OrderItems
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrderItemsSingle(t *testing.T) {
	j := []byte("{ \"OrderItems\":{ \"OrderItem\": {	\"OrderItemId\": \"1\",	\"ShopId\": \"ShopId 1\",	\"OrderId\": \"1\",	\"Name\": \"Name 1\",	\"Sku\": \"Sku 1\",	\"Variation\": \"Variation 1\",	\"ShopSku\": \"ShopSku 1\",	\"ShippingType\": \"Dropshipping 1\",	\"ItemPrice\": \"180.00\",	\"PaidPrice\": \"280.00\",	\"Currency\": \"USD\",	\"WalletCredits\": \"380.00\",	\"TaxAmount\": \"18.32\",	\"CodCollectableAmount\": \"19.32\",	\"ShippingAmount\": \"20.32\",	\"ShippingServiceCost\": \"21.32\",	\"VoucherAmount\": \"22.32\",	\"VoucherCode\": \"VoucherCode 1\",	\"Status\": \"shipped\",	\"IsProcessable\": \"1\",	\"ShipmentProvider\": \"DHL\",	\"IsDigital\": \"0\",	\"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 1\",	\"TrackingCode\": \"TrackingCode 1\",	\"TrackingCodePre\": \"TrackingCodePre 1\",	\"Reason\": \"Reason 1\",	\"ReasonDetail\": \"ReasonDetail 1\",	\"PurchaseOrderId\": \"1\",	\"PurchaseOrderNumber\": \"PurchaseOrderNumber 1\",	\"PackageId\": \"PackageId 1\",	\"PromisedShippingTime\": \"2015-11-04 10:30:57\",	\"ExtraAttributes\": \"ExtraAttributes 1\",	\"ShippingProviderType\": \"ShippingProviderType 1\",	\"CreatedAt\": \"2015-11-05 10:30:57\",	\"UpdatedAt\": \"2015-11-06 10:30:57\",	\"ReturnStatus\": \"ReturnStatus 1\"} } }")

	expected := OrderItems{[]OrderItem{
		{
			ScInt(1),
			"ShopId 1",
			ScInt(1),
			"Name 1",
			"Sku 1",
			"Variation 1",
			"ShopSku 1",
			"Dropshipping 1",
			ScFloat(180.00),
			ScFloat(280.00),
			"USD",
			ScFloat(380.00),
			ScFloat(18.32),
			ScFloat(19.32),
			ScFloat(20.32),
			ScFloat(21.32),
			ScFloat(22.32),
			"VoucherCode 1",
			"shipped",
			ScBool(true),
			"DHL",
			ScBool(false),
			"DigitalDeliveryInfo 1",
			"TrackingCode 1",
			"TrackingCodePre 1",
			"Reason 1",
			"ReasonDetail 1",
			ScInt(1),
			"PurchaseOrderNumber 1",
			"PackageId 1",
			ScTimestamp(time.Date(2015, 11, 4, 10, 30, 57, 00, time.UTC)),
			"ExtraAttributes 1",
			"ShippingProviderType 1",
			ScTimestamp(time.Date(2015, 11, 5, 10, 30, 57, 00, time.UTC)),
			ScTimestamp(time.Date(2015, 11, 6, 10, 30, 57, 00, time.UTC)),
			"ReturnStatus 1",
		},
	},
	}

	var c OrderItems
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_OrderItemsMultiple(t *testing.T) {
	j := []byte("{ \"OrderItems\":{ \"OrderItem\": [{	\"OrderItemId\": \"1\",	\"ShopId\": \"ShopId 1\",	\"OrderId\": \"1\",	\"Name\": \"Name 1\",	\"Sku\": \"Sku 1\",	\"Variation\": \"Variation 1\",	\"ShopSku\": \"ShopSku 1\",	\"ShippingType\": \"Dropshipping 1\",	\"ItemPrice\": \"180.00\",	\"PaidPrice\": \"280.00\",	\"Currency\": \"USD\",	\"WalletCredits\": \"380.00\",	\"TaxAmount\": \"18.32\",	\"CodCollectableAmount\": \"19.32\",	\"ShippingAmount\": \"20.32\",	\"ShippingServiceCost\": \"21.32\",	\"VoucherAmount\": \"22.32\",	\"VoucherCode\": \"VoucherCode 1\",	\"Status\": \"shipped\",	\"IsProcessable\": \"1\",	\"ShipmentProvider\": \"DHL\",	\"IsDigital\": \"0\",	\"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 1\",	\"TrackingCode\": \"TrackingCode 1\",	\"TrackingCodePre\": \"TrackingCodePre 1\",	\"Reason\": \"Reason 1\",	\"ReasonDetail\": \"ReasonDetail 1\",	\"PurchaseOrderId\": \"1\",	\"PurchaseOrderNumber\": \"PurchaseOrderNumber 1\",	\"PackageId\": \"PackageId 1\",	\"PromisedShippingTime\": \"2015-11-04 10:30:57\",	\"ExtraAttributes\": \"ExtraAttributes 1\",	\"ShippingProviderType\": \"ShippingProviderType 1\",	\"CreatedAt\": \"2015-11-05 10:30:57\",	\"UpdatedAt\": \"2015-11-06 10:30:57\",	\"ReturnStatus\": \"ReturnStatus 1\"},{	\"OrderItemId\": \"2\",	\"ShopId\": \"ShopId 2\",	\"OrderId\": \"1\",	\"Name\": \"Name 2\",	\"Sku\": \"Sku 2\",	\"Variation\": \"Variation 2\",	\"ShopSku\": \"ShopSku 2\",	\"ShippingType\": \"Dropshipping 2\",	\"ItemPrice\": \"1180.00\",	\"PaidPrice\": \"1280.00\",	\"Currency\": \"EUR\",	\"WalletCredits\": \"1380.00\",	\"TaxAmount\": \"118.32\",	\"CodCollectableAmount\": \"119.32\",	\"ShippingAmount\": \"120.32\",	\"ShippingServiceCost\": \"121.32\",	\"VoucherAmount\": \"122.32\",	\"VoucherCode\": \"VoucherCode 2\",	\"Status\": \"canceled\",	\"IsProcessable\": \"0\",	\"ShipmentProvider\": \"UPS\",	\"IsDigital\": \"1\",	\"DigitalDeliveryInfo\": \"DigitalDeliveryInfo 2\",	\"TrackingCode\": \"TrackingCode 2\",	\"TrackingCodePre\": \"TrackingCodePre 2\",	\"Reason\": \"Reason 2\",	\"ReasonDetail\": \"ReasonDetail 2\",	\"PurchaseOrderId\": \"2\",	\"PurchaseOrderNumber\": \"PurchaseOrderNumber 2\",	\"PackageId\": \"PackageId 2\",	\"PromisedShippingTime\": \"2016-11-04 10:30:57\",	\"ExtraAttributes\": \"ExtraAttributes 2\",	\"ShippingProviderType\": \"ShippingProviderType 2\",	\"CreatedAt\": \"2016-11-05 10:30:57\",	\"UpdatedAt\": \"2016-11-06 10:30:57\",	\"ReturnStatus\": \"ReturnStatus 2\"} ] } }")

	expected := OrderItems{[]OrderItem{
		{
			ScInt(1),
			"ShopId 1",
			ScInt(1),
			"Name 1",
			"Sku 1",
			"Variation 1",
			"ShopSku 1",
			"Dropshipping 1",
			ScFloat(180.00),
			ScFloat(280.00),
			"USD",
			ScFloat(380.00),
			ScFloat(18.32),
			ScFloat(19.32),
			ScFloat(20.32),
			ScFloat(21.32),
			ScFloat(22.32),
			"VoucherCode 1",
			"shipped",
			ScBool(true),
			"DHL",
			ScBool(false),
			"DigitalDeliveryInfo 1",
			"TrackingCode 1",
			"TrackingCodePre 1",
			"Reason 1",
			"ReasonDetail 1",
			ScInt(1),
			"PurchaseOrderNumber 1",
			"PackageId 1",
			ScTimestamp(time.Date(2015, 11, 4, 10, 30, 57, 00, time.UTC)),
			"ExtraAttributes 1",
			"ShippingProviderType 1",
			ScTimestamp(time.Date(2015, 11, 5, 10, 30, 57, 00, time.UTC)),
			ScTimestamp(time.Date(2015, 11, 6, 10, 30, 57, 00, time.UTC)),
			"ReturnStatus 1",
		},
		{
			ScInt(2),
			"ShopId 2",
			ScInt(1),
			"Name 2",
			"Sku 2",
			"Variation 2",
			"ShopSku 2",
			"Dropshipping 2",
			ScFloat(1180.00),
			ScFloat(1280.00),
			"EUR",
			ScFloat(1380.00),
			ScFloat(118.32),
			ScFloat(119.32),
			ScFloat(120.32),
			ScFloat(121.32),
			ScFloat(122.32),
			"VoucherCode 2",
			"canceled",
			ScBool(false),
			"UPS",
			ScBool(true),
			"DigitalDeliveryInfo 2",
			"TrackingCode 2",
			"TrackingCodePre 2",
			"Reason 2",
			"ReasonDetail 2",
			ScInt(2),
			"PurchaseOrderNumber 2",
			"PackageId 2",
			ScTimestamp(time.Date(2016, 11, 4, 10, 30, 57, 00, time.UTC)),
			"ExtraAttributes 2",
			"ShippingProviderType 2",
			ScTimestamp(time.Date(2016, 11, 5, 10, 30, 57, 00, time.UTC)),
			ScTimestamp(time.Date(2016, 11, 6, 10, 30, 57, 00, time.UTC)),
			"ReturnStatus 2",
		},
	},
	}

	var c OrderItems
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_Document(t *testing.T) {
	j := []byte("{ \"DocumentType\": \"shippingLabel\", \"MimeType\": \"text/html\", \"File\": \"c2hpcHBpbmdMYWJlbDogdGV4dC9odG1s\" }")

	expected := Document{
		"shippingLabel",
		"text/html",
		"c2hpcHBpbmdMYWJlbDogdGV4dC9odG1s",
	}

	var c Document
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}
