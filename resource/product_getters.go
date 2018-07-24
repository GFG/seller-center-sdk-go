package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
	"time"
)

type GetProductsParams struct {
	CreatedAfter     *time.Time
	CreatedBefore    *time.Time
	UpdatedAfter     *time.Time
	UpdatedBefore    *time.Time
	Status           *string
	Limit            *int
	Offset           *int
	SortBy           *string
	SortDirection    *string
	Search           *string
	Filter           *string
	SkuSellerList    *[]string
	GlobalIdentifier *bool
}

func (pr ProductResource) GetBrands() (model.Brands, error) {
	r := client.NewGenericRequest("GetBrands", client.MethodGET)
	r.SetVersion(client.V1)

	response, err := pr.client.Call(r)

	if err != nil {
		return model.Brands{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Brands{}, errors.New(errorResponse.HeadObject.ErrorMessage)
	}

	rawBody := response.GetBody()
	rawBrands, _, _, err := jsonparser.Get(rawBody, "Brands")

	if err != nil {
		return model.Brands{}, err
	}

	brands := model.Brands{}
	if len(rawBrands) == 0 {
		return brands, nil
	}

	err = json.Unmarshal(rawBrands, &brands)
	if err != nil {
		return model.Brands{}, err
	}

	return brands, nil
}

func (pr ProductResource) GetCategoryTree() (model.Categories, error) {
	r := client.NewGenericRequest("GetCategoryTree", client.MethodGET)
	r.SetVersion(client.V1)

	response, err := pr.client.Call(r)

	if err != nil {
		return model.Categories{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Categories{}, errors.New(errorResponse.HeadObject.ErrorMessage)
	}

	rawBody := response.GetBody()
	rawCategories, _, _, err := jsonparser.Get(rawBody, "Categories")

	if err != nil {
		return model.Categories{}, err
	}

	categories := model.Categories{}
	if len(rawCategories) == 0 {
		return categories, nil
	}

	err = json.Unmarshal(rawCategories, &categories)
	if err != nil {
		return model.Categories{}, err
	}

	return categories, nil
}

func (pr ProductResource) GetCategoryAttributes(categoryId int) (model.Attributes, error) {
	r := client.NewGenericRequest("GetCategoryAttributes", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("PrimaryCategory", strconv.Itoa(categoryId))

	response, err := pr.client.Call(r)

	if err != nil {
		return model.Attributes{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Attributes{}, errors.New(errorResponse.HeadObject.ErrorMessage)
	}

	rawBody := response.GetBody()

	attributes := model.Attributes{}
	if len(rawBody) == 0 {
		return attributes, nil
	}

	err = json.Unmarshal(rawBody, &attributes)
	if err != nil {
		return model.Attributes{}, err
	}

	return attributes, nil
}

func (pr ProductResource) GetProducts(params GetProductsParams) (model.Products, error) {

	r := client.NewGenericRequest("GetProducts", client.MethodGET)
	r.SetVersion(client.V1)

	if nil != params.CreatedAfter {
		r.SetRequestParam("CreatedAfter", params.CreatedAfter.Format(time.RFC3339))
	}
	if nil != params.CreatedBefore {
		r.SetRequestParam("CreatedBefore", params.CreatedBefore.Format(time.RFC3339))
	}
	if nil != params.UpdatedAfter {
		r.SetRequestParam("UpdatedAfter", params.UpdatedAfter.Format(time.RFC3339))
	}
	if nil != params.UpdatedBefore {
		r.SetRequestParam("UpdatedBefore", params.UpdatedBefore.Format(time.RFC3339))
	}
	if nil != params.Limit {
		r.SetRequestParam("Limit", strconv.Itoa(*params.Limit))
	}
	if nil != params.Offset {
		r.SetRequestParam("Offset", strconv.Itoa(*params.Offset))
	}
	if nil != params.Status {
		r.SetRequestParam("Status", *params.Status)
	}
	if nil != params.SortBy {
		r.SetRequestParam("SortBy", *params.SortBy)
	}
	if nil != params.SortDirection {
		r.SetRequestParam("SortDirection", *params.SortDirection)
	}

	if nil != params.Search {
		r.SetRequestParam("Search", *params.Search)
	}

	if nil != params.Filter {
		r.SetRequestParam("Filter", *params.Filter)
	}

	if nil != params.SkuSellerList {
		cache := make([]string, len(*params.SkuSellerList))
		for i, skuSeller := range *params.SkuSellerList {
			cache[i] = fmt.Sprintf(`"%s"`, skuSeller)
		}

		param := fmt.Sprintf("[%s]", strings.Join(cache, ","))
		r.SetRequestParam("SkuSellerList", param)
	}

	if nil != params.GlobalIdentifier {
		var param string
		if *params.GlobalIdentifier {
			param = "1"
		} else {
			param = "0"

		}

		r.SetRequestParam("GlobalIdentifier", param)
	}

	response, err := pr.client.Call(r)

	if err != nil {
		return model.Products{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Products{}, errors.New(errorResponse.HeadObject.ErrorMessage)
	}

	rawBody := response.GetBody()
	var products model.Products
	if err := json.Unmarshal(rawBody, &products); nil != err {
		return model.Products{}, err
	}

	return products, nil
}
