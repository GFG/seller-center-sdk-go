package resource

import (
	"encoding/xml"
	"github.com/GFG/seller-center-sdk-go/client"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

const (
	scApiBaseUrl = "https://sellerapi.sellercenter.net/"
	scApiUser    = "user@sellercenter.net"
	scApiKey     = "000000000000000000000000000000000000000"
)

func TestProductBuilder(t *testing.T) {
	logger := log.New(ioutil.Discard, "SC SDK", log.LstdFlags)

	clientConfig, err := client.NewClientConfig(
		scApiBaseUrl,
		scApiUser,
		scApiKey,
		logger,
	)

	if err != nil {
		t.Fatalf("unexpected error:`%s`.", err)
	}

	scClient := client.NewClient(*clientConfig, logger)

	if nil == scClient {
		t.Fatal("no client available")
	}

	productResource := NewProduct(scClient)

	productBuilder := *productResource.InitProduct().
		WithName("Name"). // Automatic CDATA encapsulation
		WithSellerSku("Seller Sku").
		WithParentSku("Parent Sku").
		WithStatus("active").
		WithVariation("XXS").
		WithPrimaryCategory(1).
		WithCategories([]int{2, 3}).
		WithBrowseNodes([]int{5, 6}).
		WithDescription(`This is a <b>bold</b> product.`). // Automatic CDATA encapsulation
		WithBrand("Brand").
		WithPrice(40.00).
		WithSalePrice(33).
		WithSaleStartDate(time.Date(2015, 11, 4, 10, 30, 49, 00, time.UTC)).
		WithSaleEndDate(time.Date(2015, 11, 4, 10, 30, 49, 00, time.UTC).AddDate(0, 0, 5)).
		WithTaxClass("default").
		WithMainImage("https://sellerapi.sellercenter.net/image1.jpg").
		WithImage("https://sellerapi.sellercenter.net/image2.jpg").
		WithImage("https://sellerapi.sellercenter.net/image3.jpg").
		WithProductId("Product Id").
		WithQuantity(4).
		WithShipmentType("crossdocking").
		WithCondition("new").
		WithVolumetricWeight(10.55).
		WithProductGroup("product group").
		WithProductData(
			map[string]interface{}{
				"DescriptionEn": CharData(`I am a description for the new product again`), // Explicit CDATA encapsulation
			})

	expected := `<Product><SellerSku>Seller Sku</SellerSku><Name><![CDATA[Name]]></Name><Description><![CDATA[This is a <b>bold</b> product.]]></Description><Brand>Brand</Brand><TaxClass>default</TaxClass><Variation>XXS</Variation><ParentSku>Parent Sku</ParentSku><Quantity>4</Quantity><Price>40</Price><SalePrice>33</SalePrice><SaleStartDate>2015-11-04 10:30:49</SaleStartDate><SaleEndDate>2015-11-09 10:30:49</SaleEndDate><Status>active</Status><ProductId>Product Id</ProductId><VolumetricWeight>10.55</VolumetricWeight><ProductGroup>product group</ProductGroup><MainImage>https://sellerapi.sellercenter.net/image1.jpg</MainImage><Images><Image>https://sellerapi.sellercenter.net/image2.jpg</Image><Image>https://sellerapi.sellercenter.net/image3.jpg</Image></Images><PrimaryCategory>1</PrimaryCategory><Categories>2,3</Categories><ProductData><DescriptionEn><![CDATA[I am a description for the new product again]]></DescriptionEn></ProductData><BrowseNodes>5,6</BrowseNodes><ShipmentType>crossdocking</ShipmentType><Condition>new</Condition></Product>`

	xml, err := xml.MarshalIndent(productBuilder.product, "", "")
	if err != nil {
		t.Fatalf("can not marshal. expected:`%s` - error:`%s`.", expected, err)
	}

	if string(xml) != expected {
		t.Fatalf("marshalled doesn't match. expected: `%s` - marshalled: `%s`.", expected, xml)
	}
}
