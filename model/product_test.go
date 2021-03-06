package model

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func Test_ProductsEmpty(t *testing.T) {
	j := []byte(`{}`)

	expected := Products{[]Product{}}

	var c Products
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_ProductsSingle(t *testing.T) {
	j := []byte(`{"Products":{"Product":{"SellerSku":"SellerSku 1","ShopSku":"ShopSku 1","Name":"Name 1","Description":"Description 1","Brand":"Brand 1","TaxClass":"TaxClass 1","Variation":"Variation 1","ParentSku":"ParentSku 1","Quantity":"1","FulfillmentByNonSellable":"1","Available":"1","Price":"10.10","SalePrice":"20.20","SaleStartDate":"2015-11-04 10:30:49","SaleEndDate":"2015-11-05 10:30:49","Status":"active","ProductId":"ProductId 1","Url":"Url 1","MainImage":"MainImage 1","Images":{"Image":"Image 1"},"PrimaryCategory":"PrimaryCategory 1","PrimaryCategoryId":"73","Categories":"Category 1,Category 2","CategoriesIds":"77,83","ProductData":{"ProductData 1":"ProductData 1"},"BrowseNodes":"BrowseNode 1","ShipmentType":"ShipmentType 1","Condition":"Condition 1"} } }`)

	expected := Products{[]Product{
		{
			SellerSku:                "SellerSku 1",
			ShopSku:                  "ShopSku 1",
			Name:                     "Name 1",
			Description:              "Description 1",
			Brand:                    "Brand 1",
			TaxClass:                 "TaxClass 1",
			Variation:                "Variation 1",
			ParentSku:                "ParentSku 1",
			Quantity:                 ScInt(1),
			FulfillmentByNonSellable: ScBool(true),
			Available:                ScBool(true),
			Price:                    ScFloat(10.10),
			SalePrice:                ScFloat(20.20),
			SaleStartDate:            ScTimestamp(time.Date(2015, 11, 4, 10, 30, 49, 00, time.UTC)),
			SaleEndDate:              ScTimestamp(time.Date(2015, 11, 5, 10, 30, 49, 00, time.UTC)),
			Status:                   "active",
			ProductId:                "ProductId 1",
			Url:                      "Url 1",
			MainImage:                "MainImage 1",
			Images:                   Images{"Image 1"},
			PrimaryCategory:          "PrimaryCategory 1",
			PrimaryCategoryId:        ScInt(73),
			Categories:               ScStringSlice{"Category 1", "Category 2"},
			CategoriesIds:            ScIntSlice{77, 83},
			ProductData:              map[string]interface{}{"ProductData 1": "ProductData 1"},
			BrowseNodes:              ScStringSlice{"BrowseNode 1"},
			ShipmentType:             "ShipmentType 1",
			Condition:                "Condition 1",
		},
	},
	}

	var c Products
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_ProductsSingle_Only_Mandatory_Attributes(t *testing.T) {
	j := []byte(`{"Products":{"Product":{"SellerSku":"minimalSellerSKU","ShopSku":"","Name":"minimal product","Variation":"S","ParentSku":"","Quantity":"0","Available":"0","Price":"888.00","SalePrice":"","SaleStartDate":"","SaleEndDate":"","Status":"active","ProductId":"","Url":"","MainImage":"","Images":"","Description":"","TaxClass":"","Brand":"Test MP Brand","PrimaryCategory":"Dresses","PrimaryCategoryId":"73","ProductData":{}}}}`)

	var emptyTime ScTimestamp

	expected := Products{Products: []Product{
		{
			Brand:             "Test MP Brand",
			Description:       "",
			Name:              "minimal product",
			Price:             ScFloat(888.00),
			PrimaryCategory:   "Dresses",
			PrimaryCategoryId: ScInt(73),
			SellerSku:         "minimalSellerSKU",
			Variation:         "S",
			Status:            "active",
			ShopSku:           "",
			ParentSku:         "",
			Quantity:          ScInt(0),
			Available:         ScBool(false),
			SalePrice:         ScFloat(0),
			SaleStartDate:     emptyTime,
			SaleEndDate:       emptyTime,
			ProductId:         "",
			Url:               "",
			MainImage:         "",
			Images:            Images{},
			TaxClass:          "",
			ProductData:       map[string]interface{}{},
		},
	},
	}

	var c Products
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_ProductsMultiple(t *testing.T) {
	j := []byte(`{"Products":{"Product":[{"SellerSku":"SellerSku 1","ShopSku":"ShopSku 1","Name":"Name 1","Description":"Description 1","Brand":"Brand 1","TaxClass":"TaxClass 1","Variation":"Variation 1","ParentSku":"ParentSku 1","Quantity":"1","FulfillmentByNonSellable":"1","Available":"1","Price":"10.10","SalePrice":"20.20","SaleStartDate":"2015-11-04 10:30:49","SaleEndDate":"2015-11-05 10:30:49","Status":"active","ProductId":"ProductId 1","Url":"Url 1","MainImage":"MainImage 1","Images":{"Image":"Image 1"},"PrimaryCategory":"PrimaryCategory 1","PrimaryCategoryId":"73","Categories":"Category 1","CategoriesIds":"77","ProductData":{"ProductData 1":"ProductData 1"},"BrowseNodes":"BrowseNode 1","ShipmentType":"ShipmentType 1","Condition":"Condition 1"},{"SellerSku":"SellerSku 2","ShopSku":"ShopSku 2","Name":"Name 2","Description":"Description 2","Brand":"Brand 2","TaxClass":"TaxClass 2","Variation":"Variation 2","ParentSku":"ParentSku 2","Quantity":"2","FulfillmentByNonSellable":"0","Available":"0","Price":"110.10","SalePrice":"120.20","SaleStartDate":"2016-11-04 10:30:49","SaleEndDate":"2016-11-05 10:30:49","Status":"inactive","ProductId":"ProductId 2","Url":"Url 2","MainImage":"MainImage 2","Images":{"Image":["Image 2","Image 3"]},"PrimaryCategory":"PrimaryCategory 2","PrimaryCategoryId":"74","Categories":"Category 2,Category 3","CategoriesIds":"83,84","ProductData":{"ProductData 2":"ProductData 2","ProductData 3":"ProductData 3"},"BrowseNodes":"BrowseNode 2,BrowseNode 3","ShipmentType":"ShipmentType 2","Condition":"Condition 2"}] } }`)

	expected := Products{[]Product{
		{
			SellerSku:                "SellerSku 1",
			ShopSku:                  "ShopSku 1",
			Name:                     "Name 1",
			Description:              "Description 1",
			Brand:                    "Brand 1",
			TaxClass:                 "TaxClass 1",
			Variation:                "Variation 1",
			ParentSku:                "ParentSku 1",
			Quantity:                 ScInt(1),
			FulfillmentByNonSellable: ScBool(true),
			Available:                ScBool(true),
			Price:                    ScFloat(10.10),
			SalePrice:                ScFloat(20.20),
			SaleStartDate:            ScTimestamp(time.Date(2015, 11, 4, 10, 30, 49, 00, time.UTC)),
			SaleEndDate:              ScTimestamp(time.Date(2015, 11, 5, 10, 30, 49, 00, time.UTC)),
			Status:                   "active",
			ProductId:                "ProductId 1",
			Url:                      "Url 1",
			MainImage:                "MainImage 1",
			Images:                   Images{"Image 1"},
			PrimaryCategory:          "PrimaryCategory 1",
			PrimaryCategoryId:        ScInt(73),
			Categories:               ScStringSlice{"Category 1"},
			CategoriesIds:            ScIntSlice{77},
			ProductData:              map[string]interface{}{"ProductData 1": "ProductData 1"},
			BrowseNodes:              ScStringSlice{"BrowseNode 1"},
			ShipmentType:             "ShipmentType 1",
			Condition:                "Condition 1",
		},
		{
			SellerSku:                "SellerSku 2",
			ShopSku:                  "ShopSku 2",
			Name:                     "Name 2",
			Description:              "Description 2",
			Brand:                    "Brand 2",
			TaxClass:                 "TaxClass 2",
			Variation:                "Variation 2",
			ParentSku:                "ParentSku 2",
			Quantity:                 ScInt(2),
			FulfillmentByNonSellable: ScBool(false),
			Available:                ScBool(false),
			Price:                    ScFloat(110.10),
			SalePrice:                ScFloat(120.20),
			SaleStartDate:            ScTimestamp(time.Date(2016, 11, 4, 10, 30, 49, 00, time.UTC)),
			SaleEndDate:              ScTimestamp(time.Date(2016, 11, 5, 10, 30, 49, 00, time.UTC)),
			Status:                   "inactive",
			ProductId:                "ProductId 2",
			Url:                      "Url 2",
			MainImage:                "MainImage 2",
			Images:                   Images{"Image 2", "Image 3"},
			PrimaryCategory:          "PrimaryCategory 2",
			PrimaryCategoryId:        ScInt(74),
			Categories:               ScStringSlice{"Category 2", "Category 3"},
			CategoriesIds:            ScIntSlice{83, 84},
			ProductData:              map[string]interface{}{"ProductData 2": "ProductData 2", "ProductData 3": "ProductData 3"},
			BrowseNodes:              ScStringSlice{"BrowseNode 2", "BrowseNode 3"},
			ShipmentType:             "ShipmentType 2",
			Condition:                "Condition 2",
		},
	},
	}

	var c Products
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_CategoriesEmpty(t *testing.T) {
	j := []byte(`{}`)

	expected := Categories{[]Category{}}

	var c Categories
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_Categories(t *testing.T) {
	j := []byte(`{"Category":{"Name":"Name 1","CategoryId":"1","GlobalIdentifier":"GlobalIdentifier 1","AttributeSetId":"12","Children":{"Category":{"Name":"Name 2","CategoryId":"2","AttributeSetId":"8","GlobalIdentifier":"GlobalIdentifier 2","Children":{"Category":[{"Name":"Name 3","CategoryId":"3","AttributeSetId":"8","GlobalIdentifier":"GlobalIdentifier 3","Children":""},{"Name":"Name 4","CategoryId":"4","AttributeSetId":"7","GlobalIdentifier":"GlobalIdentifier 4","Children":""}]}}}}}`)

	expected := Categories{[]Category{
		{
			"Name 1",
			ScInt(1),
			ScInt(12),
			"GlobalIdentifier 1",
			Categories{
				[]Category{
					{
						"Name 2",
						ScInt(2),
						ScInt(8),
						"GlobalIdentifier 2",
						Categories{
							[]Category{
								{
									"Name 3",
									ScInt(3),
									ScInt(8),
									"GlobalIdentifier 3",
									Categories{[]Category{}},
								},
								{
									"Name 4",
									ScInt(4),
									ScInt(7),
									"GlobalIdentifier 4",
									Categories{[]Category{}},
								},
							},
						},
					},
				},
			},
		},
	},
	}

	var c Categories
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}

func Test_AttributeOptionsEmpty(t *testing.T) {
	j := []byte(`{}`)

	expected := Attributes{}

	var c Attributes
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Errorf("unmarshalled doesn't match. expected: `%#v` - unmarshalled: `%#v`.", expected, c)
	}
}

func Test_AttributeOptions(t *testing.T) {
	j := []byte(`{ "Attribute": [{ "Label": "Label 1", "Name": "Name 1", "FeedName": "FeedName 1", "GlobalIdentifier": "GlobalIdentifier 1", "GroupName": "GroupName 1", "IsMandatory": "0", "IsGlobalAttribute": "0", "Description": "Description 1", "ProductType": "config", "InputType": "textfield", "AttributeType": "value", "ExampleValue": "ExampleValue 1", "MaxLength": "255", "Options": "" }, { "Label": "Label 2", "Name": "Name 2", "FeedName": "FeedName 2", "GlobalIdentifier": "GlobalIdentifier 2", "GroupName": "GroupName 2", "IsMandatory": "1", "IsGlobalAttribute": "1", "Description": "Description 2", "ProductType": "config", "InputType": "dropdown", "AttributeType": "option", "ExampleValue": "Options ExampleValue 2", "MaxLength": "0", "Options": {"Option": { "GlobalIdentifier": "Options GlobalIdentifier 2", "Name": "Options ExampleValue 2", "IsDefault": "1" }  } }, { "Label": "Label 3", "Name": "Name 3", "FeedName": "FeedName 3", "GlobalIdentifier": "GlobalIdentifier 3", "GroupName": "GroupName 3", "IsMandatory": "0", "IsGlobalAttribute": "1", "Description": "Description 3", "ProductType": "config", "InputType": "dropdown", "AttributeType": "option", "ExampleValue": "Spring,Summer,Autumn,Winter", "MaxLength": "", "Options": { "Option": [{ "GlobalIdentifier": "Options GlobalIdentifier 3-1", "Name": "Spring", "IsDefault": "0" }, { "GlobalIdentifier": "Options GlobalIdentifier 3-2", "Name": "Summer", "IsDefault": "1" }, { "GlobalIdentifier": "Options GlobalIdentifier 3-3", "Name": "Autumn", "IsDefault": "0" }, { "GlobalIdentifier": "Options GlobalIdentifier 3-4", "Name": "Winter", "IsDefault": "0" } ] } } ] }`)

	expected := Attributes{[]Attribute{
		{
			"Label 1",
			"Name 1",
			"FeedName 1",
			"GlobalIdentifier 1",
			ScBool(false),
			ScBool(false),
			"Description 1",
			"config",
			"textfield",
			"value",
			"ExampleValue 1",
			ScInt(255),
			AttributeOptions([]AttributeOption{}),
		},
		{
			"Label 2",
			"Name 2",
			"FeedName 2",
			"GlobalIdentifier 2",
			ScBool(true),
			ScBool(true),
			"Description 2",
			"config",
			"dropdown",
			"option",
			"Options ExampleValue 2",
			ScInt(0),
			[]AttributeOption{
				{
					"Options GlobalIdentifier 2",
					"Options ExampleValue 2",
					ScBool(true),
				},
			},
		},
		{
			"Label 3",
			"Name 3",
			"FeedName 3",
			"GlobalIdentifier 3",
			ScBool(false),
			ScBool(true),
			"Description 3",
			"config",
			"dropdown",
			"option",
			"Spring,Summer,Autumn,Winter",
			ScInt(0),
			[]AttributeOption{
				{
					"Options GlobalIdentifier 3-1",
					"Spring",
					ScBool(false),
				},
				{
					"Options GlobalIdentifier 3-2",
					"Summer",
					ScBool(true),
				},
				{
					"Options GlobalIdentifier 3-3",
					"Autumn",
					ScBool(false),
				},
				{
					"Options GlobalIdentifier 3-4",
					"Winter",
					ScBool(false),
				},
			},
		},
	},
	}

	var c Attributes
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Errorf("unmarshalled doesn't match. expected: `%#v` - unmarshalled: `%#v`.", expected, c)
	}
}

func Test_Brands(t *testing.T) {
	j := []byte(`{ "Brand": [{ "BrandId": "1", "Name": "Name 1", "GlobalIdentifier": "GlobalIdentifier 1"},{ "BrandId": "2", "Name": "Name 2", "GlobalIdentifier": "GlobalIdentifier 2"} ] }`)

	expected := Brands{[]Brand{
		{
			ScInt(1),
			"Name 1",
			"GlobalIdentifier 1",
		},
		{
			ScInt(2),
			"Name 2",
			"GlobalIdentifier 2",
		},
	},
	}

	var c Brands
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c)
	}
}
