package main

import (
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/GFG/seller-center-sdk-go/resource"
	"log"
	"os"
	"strings"
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

	productResource := resource.NewProduct(scClient)

	brands, err := productResource.GetBrands()

	if err != nil {
		logger.Panicln(err)
	}

	for _, brand := range brands.Brands {
		dumpBrand(brand, logger)
	}

	categories, err := productResource.GetCategoryTree()

	if err != nil {
		logger.Panicln(err)
	}

	for _, category := range categories.Categories {
		dumpCategory(category, logger, 0)
	}

	attributes, err := productResource.GetCategoryAttributes(3835)

	if err != nil {
		logger.Panicln(err)
	}

	for _, attribute := range attributes.Attributes {
		dumpAttribute(attribute, logger)
	}

	createdAfter := time.Date(2014, 2, 25, 0, 0, 0, 0, time.UTC)
	status := "inactive"
	limit := 2
	offset := 0
	sortBy := "created_at"
	sortDirection := "DESC"
	search := "Yellow"
	filter := "all"
	skuSellerList := []string{"SELLERSKU1", "SELLERSKU2"}
	globalIdentifier := true

	params := resource.GetProductsParams{
		CreatedAfter:     &createdAfter,
		Status:           &status,
		Limit:            &limit,
		Offset:           &offset,
		SortBy:           &sortBy,
		SortDirection:    &sortDirection,
		Search:           &search,
		Filter:           &filter,
		SkuSellerList:    &skuSellerList,
		GlobalIdentifier: &globalIdentifier,
	}

	products, err := productResource.GetProducts(params)

	if err != nil {
		logger.Panicln(err)
	}

	for _, product := range products.Products {
		dumpProduct(product, logger)
	}

	sellerSku := "SELLER SKU"
	images := model.Images([]string{
		"https://sellerapi.sellercenter.net/image1.jpg",
		"https://sellerapi.sellercenter.net/image2.jpg",
	})

	requestId, err := productResource.ProductImage(sellerSku, images)
	if err != nil {
		logger.Printf("ProductImage failed: %s\n", err)
	} else {
		logger.Printf("ProductImage succeeded, RequestId: %s\n", requestId)
	}

	brand := "A Little 7"
	primaryCategory := 1
	sellerSku0 := "Api test product"
	sellerSku1 := "Api test product again"

	newProductBuilder := make([]resource.ProductBuilder, 2)
	newProductBuilder[0] = productResource.InitProduct().
		WithName("New Product"). // Automatic CDATA encapsulation
		WithSellerSku(sellerSku0).
		WithStatus("active").
		WithVariation("XXL").
		WithPrimaryCategory(primaryCategory).
		WithDescription(`This is a <b>bold</b> product.`). // Automatic CDATA encapsulation
		WithBrand(brand).
		WithPrice(40.00).
		WithSalePrice(33).
		WithSaleStartDate(time.Now()).
		WithSaleEndDate(time.Now().AddDate(0, 0, 5)).
		WithTaxClass("default").
		WithMainImage("https://sellerapi.sellercenter.net/image1.jpg").
		WithProductData(
			map[string]interface{}{
				"DescriptionEn": model.CharData(`I am a description for the new product`), // Explicit CDATA encapsulation
				"NameEn":        `I am a new product`,
				"PackageType":   `Parcel`,
			})

	newProductBuilder[1] = productResource.InitProduct().
		WithName("New Product Again"). // Automatic CDATA encapsulation
		WithSellerSku(sellerSku1).
		WithStatus("active").
		WithVariation("XXS").
		WithPrimaryCategory(primaryCategory).
		WithDescription(`This is a <b>bold</b> product.`). // Automatic CDATA encapsulation
		WithBrand(brand).
		WithPrice(40.00).
		WithSalePrice(33).
		WithSaleStartDate(time.Now()).
		WithSaleEndDate(time.Now().AddDate(0, 0, 5)).
		WithTaxClass("default").
		WithMainImage("https://sellerapi.sellercenter.net/image1.jpg").
		WithImage("https://sellerapi.sellercenter.net/image2.jpg").
		WithImage("https://sellerapi.sellercenter.net/image3.jpg").
		WithProductData(
			map[string]interface{}{
				"DescriptionEn": model.CharData(`I am a description for the new product again`), // Explicit CDATA encapsulation
				"NameEn":        `I am a new product Again`,
				"PackageType":   `Parcel`,
			})

	requestId, err = productResource.ProductCreate(newProductBuilder)
	if err != nil {
		logger.Printf("ProductCreate failed: %s\n", err)
	} else {
		logger.Printf("ProductCreate succeeded, RequestId: %s\n", requestId)
	}

	updateProductBuilder := make([]resource.ProductBuilder, 1)
	updateProductBuilder[0] = productResource.InitProduct().
		WithSellerSku(sellerSku0).
		WithName("Updated Product"). // Automatic CDATA encapsulation
		WithStatus("inactive").
		WithProductData(
			map[string]interface{}{
				"DescriptionEn": model.CharData(`I am an updated description for the old product`), // Explicit CDATA encapsulation
			})

	requestId, err = productResource.ProductUpdate(updateProductBuilder)
	if err != nil {
		logger.Printf("ProductUpdate failed: %s\n", err)
	} else {
		logger.Printf("ProductUpdate succeeded, RequestId: %s\n", requestId)
	}

}

func dumpProduct(product model.Product, logger *log.Logger) {
	logger.Println("---------")
	logger.Printf("SellerSku: %s\n", product.SellerSku)
	logger.Printf("ShopSku: %s\n", product.ShopSku)
	logger.Printf("Name: %s\n", product.Name)
	logger.Printf("Description: %s\n", product.Description)
	logger.Printf("Brand: %s\n", product.Brand)
	logger.Printf("TaxClass: %s\n", product.TaxClass)
	logger.Printf("Variation: %s\n", product.Variation)
	logger.Printf("ParentSku: %s\n", product.ParentSku)
	logger.Printf("Quantity: %d\n", product.Quantity)
	logger.Printf("FulfillmentByNonSellable: %t\n", product.FulfillmentByNonSellable)
	logger.Printf("Available: %t\n", product.Available)
	logger.Printf("Price: %f\n", product.Price)
	logger.Printf("SalePrice: %f\n", product.SalePrice)
	logger.Printf("SaleStartDate: %s\n", time.Time(product.SaleStartDate).Format("2006-01-02 15:04:05"))
	logger.Printf("SaleEndDate: %s\n", time.Time(product.SaleEndDate).Format("2006-01-02 15:04:05"))
	logger.Printf("Status: %s\n", product.Status)
	logger.Printf("ProductId: %s\n", product.ProductId)
	logger.Printf("Url: %s\n", product.Url)
	logger.Printf("MainImage: %s\n", product.MainImage)
	logger.Println("Images:")
	for _, image := range product.Images {
		logger.Printf("	%s\n", image)
	}
	logger.Printf("PrimaryCategory: %s\n", product.PrimaryCategory)
	logger.Println("Categories:")
	for _, category := range product.Categories {
		logger.Printf("	%d\n", category)
	}
	logger.Println("ProductData:")
	for key, value := range product.ProductData {
		logger.Printf("	%s: %v\n", key, value)
	}
	logger.Println("BrowseNodes:")
	for _, browseNode := range product.BrowseNodes {
		logger.Printf("	%d\n", browseNode)
	}
	logger.Printf("ShipmentType: %s\n", product.ShipmentType)
	logger.Printf("Condition: %s\n", product.Condition)
}

func dumpAttribute(attribute model.Attribute, logger *log.Logger) {
	logger.Println("---------")
	logger.Printf("Name: %s\n", attribute.Name)
	logger.Printf("AttributeType: %s\n", attribute.AttributeType)
	logger.Printf("Description: %s\n", attribute.Description)
	logger.Printf("ExampleValue: %s\n", attribute.ExampleValue)
	logger.Printf("FeedName: %s\n", attribute.FeedName)
	logger.Printf("GlobalIdentifier: %s\n", attribute.GlobalIdentifier)
	logger.Printf("InputType: %s\n", attribute.InputType)
	logger.Printf("IsGlobalAttribute: %t\n", attribute.IsGlobalAttribute)
	logger.Printf("Label: %s\n", attribute.Label)
	logger.Printf("MaxLength: %d\n", attribute.MaxLength)
	logger.Printf("ProductType: %s\n", attribute.ProductType)
	logger.Println("Options:")

	for _, option := range attribute.Options {
		logger.Println("	---------")
		logger.Printf("	GlobalIdentifier: %s\n", option.GlobalIdentifier)
		logger.Printf("	Name: %s\n", option.Name)
	}

}
func dumpBrand(brand model.Brand, logger *log.Logger) {
	logger.Println("---------")
	logger.Printf("BrandId: %d\n", brand.BrandId)
	logger.Printf("Name: %s\n", brand.Name)
	logger.Printf("GlobalIdentifier: %s\n", brand.GlobalIdentifier)
}

func dumpCategory(category model.Category, logger *log.Logger, padding int) {
	indent := strings.Repeat("	", padding)
	logger.Printf("%s---------\n", indent)
	logger.Printf("%sName: %s\n", indent, category.Name)
	logger.Printf("%sCategoryId: %d\n", indent, category.CategoryId)
	logger.Printf("%sGlobalIdentifier: %s\n", indent, category.GlobalIdentifier)
	logger.Printf("%sChildren:\n", indent)

	padding = padding + 1
	for _, child := range category.Children.Categories {
		dumpCategory(child, logger, padding)
	}
}
