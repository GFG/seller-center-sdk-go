package model

import (
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
)

type Products struct {
	Products []Product `json:"Products"`
}

func (p *Products) UnmarshalJSON(b []byte) error {
	rawProducts, dataType, _, err := jsonparser.Get(b, "Products", "Product")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawProducts) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot find order product")
	}

	var products []Product
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawProducts, &products); nil != err {
			return err
		}
	case jsonparser.Object:
		var product Product
		if err := json.Unmarshal(rawProducts, &product); nil != err {
			return err
		}

		products = []Product{product}
	}

	*p = Products{products}

	return nil
}

type Images []string

func (i *Images) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	raw, dataType, _, err := jsonparser.Get(b, "Image")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(raw) == 0 || dataType == jsonparser.NotExist {
		return nil
	}

	var images []string
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(raw, &images); nil != err {
			return err
		}
	case jsonparser.String:
		images = Images{string(raw)}
	}

	*i = images

	return nil
}

type Product struct {
	SellerSku                string                 `json:"SellerSku"`
	ShopSku                  string                 `json:"ShopSku"`
	Name                     string                 `json:"Name"`
	Description              string                 `json:"Description"`
	Brand                    string                 `json:"Brand"`
	TaxClass                 string                 `json:"TaxClass"`
	Variation                string                 `json:"Variation"`
	ParentSku                string                 `json:"ParentSku"`
	Quantity                 ScInt                  `json:"Quantity"`
	FulfillmentByNonSellable ScBool                 `json:"FulfillmentByNonSellable"`
	Available                ScBool                 `json:"Available"`
	Price                    ScFloat                `json:"Price"`
	SalePrice                ScFloat                `json:"SalePrice"`
	SaleStartDate            ScTimestamp            `json:"SaleStartDate"`
	SaleEndDate              ScTimestamp            `json:"SaleEndDate"`
	Status                   string                 `json:"Status"`
	ProductId                string                 `json:"ProductId"`
	Url                      string                 `json:"Url"`
	MainImage                string                 `json:"MainImage"`
	Images                   Images                 `json:"Images"`
	PrimaryCategory          string                 `json:"PrimaryCategory"`
	Categories               ScStringSlice          `json:"Categories"`
	ProductData              map[string]interface{} `json:"ProductData"`
	BrowseNodes              ScStringSlice          `json:"BrowseNodes"`
	ShipmentType             string                 `json:"ShipmentType"`
	Condition                string                 `json:"Condition"`
}

type Categories struct {
	Categories []Category `json:"Category"`
}

func (c *Categories) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	raw, dataType, _, err := jsonparser.Get(b, "Category")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(raw) == 0 || dataType == jsonparser.NotExist {
		return nil
	}

	var categories []Category
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(raw, &categories); nil != err {
			return err
		}
	case jsonparser.Object:
		var category Category
		if err := json.Unmarshal(raw, &category); nil != err {
			return err
		}

		categories = []Category{category}
	}

	c.Categories = categories

	return nil
}

type Category struct {
	Name             string     `json:"Name"`
	CategoryId       ScInt      `json:"CategoryId"`
	GlobalIdentifier string     `json:"GlobalIdentifier"`
	Children         Categories `json:"Children"`
}

type Brands struct {
	Brands []Brand `json:"Brand"`
}

type Brand struct {
	BrandId          ScInt  `json:"BrandId"`
	Name             string `json:"Name"`
	GlobalIdentifier string `json:"GlobalIdentifier"`
}

type Attributes struct {
	Attributes []Attribute `json:"Attribute"`
}

type Attribute struct {
	Label             string           `json:"Label"`
	Name              string           `json:"Name"`
	FeedName          string           `json:"FeedName"`
	GlobalIdentifier  string           `json:"GlobalIdentifier"`
	IsMandatory       ScBool           `json:"IsMandatory"`
	IsGlobalAttribute ScBool           `json:"IsGlobalAttribute"`
	Description       string           `json:"Description"`
	ProductType       string           `json:"ProductType"`
	InputType         string           `json:"InputType"`
	AttributeType     string           `json:"AttributeType"`
	ExampleValue      string           `json:"ExampleValue"`
	MaxLength         ScInt            `json:"MaxLength"`
	Options           AttributeOptions `json:"Options"`
}

type AttributeOptions []AttributeOption

func (ao *AttributeOptions) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	raw, dataType, _, err := jsonparser.Get(b, "Option")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(raw) == 0 || dataType == jsonparser.NotExist {
		return nil
	}

	var options []AttributeOption
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(raw, &options); nil != err {
			return err
		}
	case jsonparser.Object:
		var option AttributeOption
		if err := json.Unmarshal(raw, &option); nil != err {
			return err
		}

		options = AttributeOptions{option}
	}

	*ao = options

	return nil
}

type AttributeOption struct {
	GlobalIdentifier string `json:"GlobalIdentifier"`
	Name             string `json:"Name"`
	IsDefault        ScBool `json:"IsDefault"`
}
