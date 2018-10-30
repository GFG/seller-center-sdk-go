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
	pb.product.SellerSku = &sellerSku

	return pb
}

func (pb *ProductBuilder) WithParentSku(parentSku string) *ProductBuilder {
	pb.product.ParentSku = &parentSku

	return pb
}

func (pb *ProductBuilder) WithStatus(status string) *ProductBuilder {
	pb.product.Status = &status

	return pb
}

func (pb *ProductBuilder) WithName(name string) *ProductBuilder {
	cdName := model.CharData(name)
	pb.product.Name = &cdName

	return pb
}

func (pb *ProductBuilder) WithVariation(variation string) *ProductBuilder {
	pb.product.Variation = &variation

	return pb
}

func (pb *ProductBuilder) WithPrimaryCategory(primaryCategory int) *ProductBuilder {
	pb.product.PrimaryCategory = &primaryCategory

	return pb
}

func (pb *ProductBuilder) WithCategories(categories []int) *ProductBuilder {
	isCategories := model.IntSlice(categories)
	pb.product.Categories = &isCategories

	return pb
}

func (pb *ProductBuilder) WithBrowseNodes(browseNodes []int) *ProductBuilder {
	isBrowseNodes := model.IntSlice(browseNodes)
	pb.product.BrowseNodes = &isBrowseNodes

	return pb
}

func (pb *ProductBuilder) WithDescription(description string) *ProductBuilder {
	cdDescription := model.CharData(description)
	pb.product.Description = &cdDescription

	return pb
}

func (pb *ProductBuilder) WithBrand(brand string) *ProductBuilder {
	pb.product.Brand = &brand

	return pb
}

func (pb *ProductBuilder) WithPrice(price float64) *ProductBuilder {
	pb.product.Price = &price

	return pb
}

func (pb *ProductBuilder) WithSalePrice(salePrice float64) *ProductBuilder {
	pb.product.SalePrice = &salePrice

	return pb
}

func (pb *ProductBuilder) WithSaleStartDate(saleStartDate time.Time) *ProductBuilder {
	t := saleDate(saleStartDate)
	pb.product.SaleStartDate = &t

	return pb
}

func (pb *ProductBuilder) WithSaleEndDate(saleEndDate time.Time) *ProductBuilder {
	t := saleDate(saleEndDate)
	pb.product.SaleEndDate = &t

	return pb
}

func (pb *ProductBuilder) WithTaxClass(taxClass string) *ProductBuilder {
	pb.product.TaxClass = &taxClass

	return pb
}

func (pb *ProductBuilder) WithShipmentType(shipmentType string) *ProductBuilder {
	pb.product.ShipmentType = &shipmentType

	return pb
}

func (pb *ProductBuilder) WithProductId(productId string) *ProductBuilder {
	pb.product.ProductId = &productId

	return pb
}

func (pb *ProductBuilder) WithCondition(condition string) *ProductBuilder {
	pb.product.Condition = &condition

	return pb
}

func (pb *ProductBuilder) WithProductData(productData productDataEntity) *ProductBuilder {
	if pb.product.ProductData == nil {
		pb.product.ProductData = &productData
	} else {
		for k, data := range productData {
			(*pb.product.ProductData)[k] = data
		}
	}

	return pb
}

func (pb *ProductBuilder) WithQuantity(quantity int) *ProductBuilder {
	pb.product.Quantity = &quantity

	return pb
}

func (pb *ProductBuilder) WithVolumetricWeight(volumetricWeight float64) *ProductBuilder {
	pb.product.VolumetricWeight = &volumetricWeight

	return pb
}

func (pb *ProductBuilder) WithProductGroup(productGroup string) *ProductBuilder {
	pb.product.ProductGroup = &productGroup

	return pb
}

func (pb *ProductBuilder) WithMainImage(image string) *ProductBuilder {
	if pb.product.ProductData == nil {
		productData := productDataEntity{}
		pb.product.ProductData = &productData
	}

	(*pb.product.ProductData)["MainImage"] = image

	return pb
}

func (pb *ProductBuilder) WithImage(image string) *ProductBuilder {
	if pb.product.images == nil {
		pb.product.images = make([]string, 0)
	}

	if len(pb.product.images) < 7 {
		pb.product.images = append(pb.product.images, image)
	}

	if pb.product.ProductData == nil {
		productData := productDataEntity{}
		pb.product.ProductData = &productData
	}

	for i, image := range pb.product.images {
		fieldName := fmt.Sprintf("Image%d", i+2)

		(*pb.product.ProductData)[fieldName] = image
	}

	return pb
}

func (pr ProductResource) InitProduct() *ProductBuilder {
	return &ProductBuilder{product: productEntry{}}
}
