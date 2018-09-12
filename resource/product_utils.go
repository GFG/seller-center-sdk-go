package resource

import (
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

func (pb ProductBuilder) WithSellerSku(sellerSku string) ProductBuilder {
	c := pb
	c.product.SellerSku = &sellerSku

	return c
}

func (pb ProductBuilder) WithParentSku(parentSku string) ProductBuilder {
	c := pb
	c.product.ParentSku = &parentSku

	return c
}

func (pb ProductBuilder) WithStatus(status string) ProductBuilder {
	c := pb
	c.product.Status = &status

	return c
}
func (pb ProductBuilder) WithName(name string) ProductBuilder {
	cdName := model.CharData(name)
	c := pb
	c.product.Name = &cdName

	return c
}

func (pb ProductBuilder) WithVariation(variation string) ProductBuilder {
	c := pb
	c.product.Variation = &variation

	return c
}

func (pb ProductBuilder) WithPrimaryCategory(primaryCategory int) ProductBuilder {
	c := pb
	c.product.PrimaryCategory = &primaryCategory

	return c
}

func (pb ProductBuilder) WithCategories(categories []int) ProductBuilder {
	isCategories := model.IntSlice(categories)
	c := pb
	c.product.Categories = &isCategories

	return c
}

func (pb ProductBuilder) WithBrowseNodes(browseNodes []int) ProductBuilder {
	isBrowseNodes := model.IntSlice(browseNodes)
	c := pb
	c.product.BrowseNodes = &isBrowseNodes

	return c
}

func (pb ProductBuilder) WithDescription(description string) ProductBuilder {
	cdDescription := model.CharData(description)
	c := pb
	c.product.Description = &cdDescription

	return c
}

func (pb ProductBuilder) WithBrand(brand string) ProductBuilder {
	c := pb
	c.product.Brand = &brand

	return c
}

func (pb ProductBuilder) WithPrice(price float64) ProductBuilder {
	c := pb
	c.product.Price = &price

	return c
}

func (pb ProductBuilder) WithSalePrice(salePrice float64) ProductBuilder {
	c := pb
	c.product.SalePrice = &salePrice

	return c
}

func (pb ProductBuilder) WithSaleStartDate(saleStartDate time.Time) ProductBuilder {
	t := saleDate(saleStartDate)
	c := pb
	c.product.SaleStartDate = &t

	return c
}

func (pb ProductBuilder) WithSaleEndDate(saleEndDate time.Time) ProductBuilder {
	t := saleDate(saleEndDate)
	c := pb
	c.product.SaleEndDate = &t

	return c
}

func (pb ProductBuilder) WithTaxClass(taxClass string) ProductBuilder {
	c := pb
	c.product.TaxClass = &taxClass

	return c
}

func (pb ProductBuilder) WithShipmentType(shipmentType string) ProductBuilder {
	c := pb
	c.product.ShipmentType = &shipmentType

	return c
}

func (pb ProductBuilder) WithProductId(productId string) ProductBuilder {
	c := pb
	c.product.ProductId = &productId

	return c
}

func (pb ProductBuilder) WithCondition(condition string) ProductBuilder {
	c := pb
	c.product.Condition = &condition

	return c
}

func (pb ProductBuilder) WithProductData(productData productDataEntity) ProductBuilder {
	c := pb
	c.product.ProductData = &productData

	return c
}

func (pb ProductBuilder) WithQuantity(quantity int) ProductBuilder {
	c := pb
	c.product.Quantity = &quantity

	return c
}

func (pb ProductBuilder) WithMainImage(mainImage string) ProductBuilder {
	c := pb
	c.product.MainImage = &mainImage

	return c
}

func (pb ProductBuilder) WithVolumetricWeight(volumetricWeight float64) ProductBuilder {
	c := pb
	c.product.VolumetricWeight = &volumetricWeight

	return c
}

func (pb ProductBuilder) WithProductGroup(productGroup string) ProductBuilder {
	c := pb
	c.product.ProductGroup = &productGroup

	return c
}

func (pb ProductBuilder) WithImage(image string) ProductBuilder {
	c := pb
	if c.product.Images == nil {
		images := make([]string, 0)
		c.product.Images = &productImagesEntries{Image: images}
	}

	c.product.Images.Image = append(c.product.Images.Image, image)

	return c
}

func (pr ProductResource) InitProduct() ProductBuilder {
	return ProductBuilder{product: productEntry{}}
}
