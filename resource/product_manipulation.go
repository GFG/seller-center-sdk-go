package resource

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
	"time"
)

type warning struct {
	Field   string `json:"Field"`
	Message string `json:"Message"`
	Value   string `json:"Value"`
}

func newApiWarningError(w []warning) error {
	return &ApiWarningError{w}
}

type ApiWarningError struct {
	w []warning
}

func (e *ApiWarningError) Error() string {
	messages := make([]string, 0, len(e.w))
	for _, warning := range e.w {
		messages = append(messages, fmt.Sprintf(`Field "%s": %s (%s)`, warning.Field, warning.Message, warning.Value))
	}

	return strings.Join(messages, "\n")
}

func (pr ProductResource) ProductImage(sellerSku string, images model.Images) (string, error) {
	r := client.NewGenericRequest("Image", client.MethodPOST)
	r.SetVersion(client.V1)

	postData := productImageXmlBody{
		SellerSku: sellerSku,
		Images:    productImagesEntries{Image: images},
	}

	r.SetPostData(postData)

	response, err := pr.client.Call(r)

	if err != nil {
		return "", err
	}

	return extractPostReponseReturnValues(response)
}

func (pr ProductResource) ProductCreate(productBuilders []ProductBuilder) (string, error) {
	r := client.NewGenericRequest("ProductCreate", client.MethodPOST)
	r.SetVersion(client.V1)

	products := make([]productEntry, len(productBuilders))
	for i, productBuilder := range productBuilders {
		products[i] = productBuilder.product

	}

	postData := products

	r.SetPostData(postData)

	response, err := pr.client.Call(r)

	if err != nil {
		return "", err
	}

	return extractPostReponseReturnValues(response)
}

func (pr ProductResource) ProductUpdate(productBuilders []ProductBuilder) (string, error) {
	r := client.NewGenericRequest("ProductUpdate", client.MethodPOST)
	r.SetVersion(client.V1)

	products := make([]productEntry, len(productBuilders))
	for i, productBuilder := range productBuilders {
		products[i] = productBuilder.product

	}

	postData := products

	r.SetPostData(postData)

	response, err := pr.client.Call(r)

	if err != nil {
		return "", err
	}

	return extractPostReponseReturnValues(response)
}

type productImagesEntries struct {
	XMLName xml.Name `xml:"Images"`
	Image   []string `xml:"Images>Image`
}

type productEntry struct {
	XMLName          xml.Name  `xml:"Product"`
	SellerSku        *string   `xml:"SellerSku"`
	Name             *CharData `xml:"Name"`
	Description      *CharData `xml:"Description"`
	Brand            *string   `xml:"Brand"`
	TaxClass         *string   `xml:"TaxClass"`
	Variation        *string   `xml:"Variation"`
	ParentSku        *string   `xml:"ParentSku"`
	Quantity         *int      `xml:"Quantity"`
	Price            *float64  `xml:"Price"`
	SalePrice        *float64  `xml:"SalePrice"`
	SaleStartDate    *saleDate `xml:"SaleStartDate"`
	SaleEndDate      *saleDate `xml:"SaleEndDate"`
	Status           *string   `xml:"Status"`
	ProductId        *string   `xml:"ProductId"`
	VolumetricWeight *float64  `xml:"VolumetricWeight"`
	ProductGroup     *string   `xml:"ProductGroup"`
	MainImage        *string   `xml:"MainImage"`
	Images           *productImagesEntries
	PrimaryCategory  *int               `xml:"PrimaryCategory"`
	Categories       *intSlice          `xml:"Categories"`
	ProductData      *productDataEntity `xml:"ProductData"`
	BrowseNodes      *intSlice          `xml:"BrowseNodes"`
	ShipmentType     *string            `xml:"ShipmentType"`
	Condition        *string            `xml:"Condition"`
}

type productImageXmlBody struct {
	XMLName   xml.Name `xml:"ProductImage"`
	SellerSku string   `xml:"SellerSku`
	Images    productImagesEntries
}

func extractPostReponseReturnValues(response client.Response) (string, error) {
	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return "", errors.New(errorResponse.HeadObject.ErrorMessage)
	}

	rawBody := response.GetBody()

	warnings, dataType, _, err := jsonparser.Get(rawBody, "WarningDetail")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return extractRequestId(response)
	}

	if len(warnings) == 0 || dataType == jsonparser.NotExist {
		return extractRequestId(response)
	}

	var cache []warning
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(warnings, &cache); nil != err {
			return "", err
		}
	case jsonparser.Object:
		var w warning
		if err := json.Unmarshal(warnings, &w); nil != err {
			return "", err
		}

		cache = []warning{w}
	}

	if len(cache) == 0 {
		return extractRequestId(response)
	}

	return "", newApiWarningError(cache)

}

func extractRequestId(response client.Response) (string, error) {
	rawHead := response.GetHead()
	requestId, err := jsonparser.GetString(rawHead, "RequestId")

	if err != nil {
		return "", err
	}

	return requestId, nil
}

type CharData string

func (cd CharData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		S string `xml:",cdata"`
	}{
		S: string(cd),
	}, start)
}

type productDataEntity map[string]interface{}

func (pd productDataEntity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}
	for k, v := range pd {
		switch v.(type) {
		case CharData:
			v := fmt.Sprintf("[CDATA[%s]]", v)
			t := xml.StartElement{Name: xml.Name{"", k}}
			tokens = append(tokens, t, xml.Directive(v), xml.EndElement{t.Name})
		default:
			v := fmt.Sprintf("%v", v)
			t := xml.StartElement{Name: xml.Name{"", k}}
			tokens = append(tokens, t, xml.CharData(v), xml.EndElement{t.Name})
		}
	}

	tokens = append(tokens, xml.EndElement{start.Name})
	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

type saleDate time.Time

func (sd saleDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	dateString := time.Time(sd).Format(saleDateTimeFormat)
	e.EncodeElement(dateString, start)

	return nil
}

type intSlice []int

func (is intSlice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	intSliceString := ""
	if len(is) == 0 {
		return nil
	}

	b := make([]string, len(is))
	for i, v := range is {
		b[i] = strconv.Itoa(v)
	}

	intSliceString = strings.Join(b, ",")
	e.EncodeElement(intSliceString, start)

	return nil
}
