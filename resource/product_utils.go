package resource

import (
	"fmt"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"time"
)

type ProductResource struct {
	client client.Client
}

func NewProduct(client client.Client) ProductResource {
	return ProductResource{client: client}
}

type ProductBuilder struct {
	product productEntry
}

func (pb *ProductBuilder) WithSellerSku(sellerSku string) *ProductBuilder {
	product := pb.product
	product.SellerSku = &sellerSku
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithParentSku(parentSku string) *ProductBuilder {
	product := pb.product
	product.ParentSku = &parentSku
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithStatus(status string) *ProductBuilder {
	product := pb.product
	product.Status = &status
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithName(name string) *ProductBuilder {
	cdName := model.CharData(name)
	product := pb.product
	product.Name = &cdName
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithVariation(variation string) *ProductBuilder {
	product := pb.product
	product.Variation = &variation
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithPrimaryCategory(primaryCategory int) *ProductBuilder {
	product := pb.product
	product.PrimaryCategory = &primaryCategory
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithCategories(categories []int) *ProductBuilder {
	isCategories := model.IntSlice(categories)
	product := pb.product
	product.Categories = &isCategories
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithBrowseNodes(browseNodes []int) *ProductBuilder {
	isBrowseNodes := model.IntSlice(browseNodes)
	product := pb.product
	product.BrowseNodes = &isBrowseNodes
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithDescription(description string) *ProductBuilder {
	cdDescription := model.CharData(description)
	product := pb.product
	product.Description = &cdDescription
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithBrand(brand string) *ProductBuilder {
	product := pb.product
	product.Brand = &brand
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithPrice(price float64) *ProductBuilder {
	product := pb.product
	product.Price = &price
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithSalePrice(salePrice float64) *ProductBuilder {
	product := pb.product
	product.SalePrice = &salePrice
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithSaleStartDate(saleStartDate time.Time) *ProductBuilder {
	t := saleDate(saleStartDate)
	product := pb.product
	product.SaleStartDate = &t
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithSaleEndDate(saleEndDate time.Time) *ProductBuilder {
	t := saleDate(saleEndDate)
	product := pb.product
	product.SaleEndDate = &t
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithTaxClass(taxClass string) *ProductBuilder {
	product := pb.product
	product.TaxClass = &taxClass
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithShipmentType(shipmentType string) *ProductBuilder {
	product := pb.product
	product.ShipmentType = &shipmentType
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithProductId(productId string) *ProductBuilder {
	product := pb.product
	product.ProductId = &productId
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithCondition(condition string) *ProductBuilder {
	product := pb.product
	product.Condition = &condition
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithProductData(productData productDataEntity) *ProductBuilder {
	product := pb.product
	if product.ProductData == nil {
		product.ProductData = &productData
	} else {
		for k, data := range productData {
			(*product.ProductData)[k] = data
		}
	}

	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithQuantity(quantity int) *ProductBuilder {
	product := pb.product
	product.Quantity = &quantity
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithVolumetricWeight(volumetricWeight float64) *ProductBuilder {
	product := pb.product
	product.VolumetricWeight = &volumetricWeight
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithProductGroup(productGroup string) *ProductBuilder {
	product := pb.product
	product.ProductGroup = &productGroup
	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithMainImage(image string) *ProductBuilder {
	product := pb.product
	if product.ProductData == nil {
		productData := productDataEntity{}
		product.ProductData = &productData
	}

	(*product.ProductData)["MainImage"] = image

	pb.product = product

	return pb
}

func (pb *ProductBuilder) WithImage(image string) *ProductBuilder {
	product := pb.product
	if product.images == nil {
		product.images = make([]string, 0)
	}

	if len(product.images) < 7 {
		product.images = append(product.images, image)
	}

	if product.ProductData == nil {
		productData := productDataEntity{}
		product.ProductData = &productData
	}

	for i, image := range product.images {
		fieldName := fmt.Sprintf("Image%d", i+2)

		(*product.ProductData)[fieldName] = image
	}

	pb.product = product

	return pb
}

func (pr ProductResource) InitProduct() *ProductBuilder {
	return &ProductBuilder{product: productEntry{}}
}
