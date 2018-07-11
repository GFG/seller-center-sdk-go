package endpoint

import (
	"encoding/json"
	"fmt"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

type OrderEndpoint struct {
	client client.Client
}

func NewOrder(client client.Client) OrderEndpoint {
	return OrderEndpoint{client: client}
}

func (oe OrderEndpoint) GetOrders(
	createdAfter *time.Time,
	createdBefore *time.Time,
	updatedAfter *time.Time,
	updatedBefore *time.Time,
	status *string,
	limit *int,
	offset *int,
	sortBy *string,
	sortDirection *string) (model.Orders, error) {

	r := client.NewGenericRequest("GetOrders", client.MethodGET)
	r.SetVersion(client.V1)

	if nil != createdAfter {
		r.SetRequestParam("CreatedAfter", createdAfter.Format(time.RFC3339))
	}
	if nil != createdBefore {
		r.SetRequestParam("CreatedBefore", createdBefore.Format(time.RFC3339))
	}
	if nil != updatedAfter {
		r.SetRequestParam("UpdatedAfter", updatedAfter.Format(time.RFC3339))
	}
	if nil != updatedBefore {
		r.SetRequestParam("UpdatedBefore", updatedBefore.Format(time.RFC3339))
	}
	if nil != limit {
		r.SetRequestParam("Limit", strconv.Itoa(*limit))
	}
	if nil != offset {
		r.SetRequestParam("Offset", strconv.Itoa(*offset))
	}
	if nil != status {
		r.SetRequestParam("Status", *status)
	}
	if nil != sortBy {
		r.SetRequestParam("SortBy", *sortBy)
	}
	if nil != sortDirection {
		r.SetRequestParam("SortDirection", *sortDirection)
	}

	response, err := oe.client.Call(r)

	if err != nil {
		return model.Orders{}, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return model.Orders{}, err
		}

		return model.Orders{}, errors.New(errorMessage)
	}

	rawBody := response.GetBody()
	rawOrders, _, _, err := jsonparser.Get(rawBody, "Orders")

	if err != nil {
		return model.Orders{}, err
	}

	orders := model.Orders{}
	if len(rawOrders) == 0 {
		return orders, nil
	}

	err = json.Unmarshal(rawOrders, &orders)

	return orders, err
}

func (oe OrderEndpoint) GetOrder(orderId int) (model.Order, error) {
	r := client.NewGenericRequest("GetOrder", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderId", strconv.Itoa(orderId))

	response, err := oe.client.Call(r)

	if err != nil {
		return model.Order{}, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return model.Order{}, err
		}

		return model.Order{}, errors.New(errorMessage)
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return model.Order{}, err
		}

		return model.Order{}, errors.New(errorMessage)
	}

	rawBody := response.GetBody()
	rawOrders, _, _, err := jsonparser.Get(rawBody, "Orders")

	if err != nil {
		return model.Order{}, err
	}

	rawOrder, _, _, err := jsonparser.Get(rawOrders, "Order")

	if err != nil {
		return model.Order{}, err
	}

	order := model.Order{}
	if len(rawOrder) == 0 {
		return model.Order{}, nil
	}

	err = json.Unmarshal(rawOrder, &order)

	return model.Order{}, err
}

func (oe OrderEndpoint) GetOrderItems(orderId int) (model.OrderItems, error) {
	r := client.NewGenericRequest("GetOrderItems", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderId", strconv.Itoa(orderId))

	response, err := oe.client.Call(r)

	if err != nil {
		return model.OrderItems{}, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return model.OrderItems{}, err
		}

		return model.OrderItems{}, errors.New(errorMessage)
	}

	rawBody := response.GetBody()
	rawOrderItems, _, _, err := jsonparser.Get(rawBody, "OrderItems")

	if err != nil {
		return model.OrderItems{}, err
	}

	orderItems := model.OrderItems{}
	if len(rawOrderItems) == 0 {
		return orderItems, nil
	}

	err = json.Unmarshal(rawOrderItems, &orderItems)

	return orderItems, err
}

func (oe OrderEndpoint) GetDocument(orderItemIds []int, documentType model.DocumentType) (model.Document, error) {
	r := client.NewGenericRequest("GetDocument", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemIds", intSliceToParam(orderItemIds))
	r.SetRequestParam("DocumentType", string(documentType))

	response, err := oe.client.Call(r)

	if err != nil {
		return model.Document{}, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return model.Document{}, err
		}

		return model.Document{}, errors.New(errorMessage)
	}

	rawBody := response.GetBody()
	rawDocuments, _, _, err := jsonparser.Get(rawBody, "Documents")

	if err != nil {
		return model.Document{}, err
	}

	rawDocument, _, _, err := jsonparser.Get(rawDocuments, "Document")

	if err != nil {
		return model.Document{}, err
	}

	document := model.Document{}
	if len(rawDocument) == 0 {
		return model.Document{}, nil
	}

	err = json.Unmarshal(rawDocument, &document)

	return document, err
}

func (oe OrderEndpoint) SetStatusToCanceled(orderItemId int, reason string, reasonDetail string) (bool, error) {
	r := client.NewGenericRequest("SetStatusToCanceled", client.MethodPOST)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemId", strconv.Itoa(orderItemId))
	r.SetRequestParam("Reason", reason)
	r.SetRequestParam("ReasonDetail", reasonDetail)

	response, err := oe.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return false, err
		}

		return false, errors.New(errorMessage)
	}

	return true, nil
}

func (oe OrderEndpoint) SetStatusToPackedByMarketplace(orderItemIds []int, deliveryType model.DeliveryType, shippingProvider string) (bool, error) {
	r := client.NewGenericRequest("SetStatusToPackedByMarketplace", client.MethodPOST)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemIds", intSliceToParam(orderItemIds))
	r.SetRequestParam("DeliveryType", string(deliveryType))
	r.SetRequestParam("ShippingProvider", shippingProvider)

	response, err := oe.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return false, err
		}

		return false, errors.New(errorMessage)
	}

	return true, nil
}

func (oe OrderEndpoint) SetStatusToReadyToShip(orderItemIds []int, deliveryType model.DeliveryType, shippingProvider string, trackingNumber string) (bool, error) {
	r := client.NewGenericRequest("SetStatusToReadyToShip", client.MethodPOST)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemIds", intSliceToParam(orderItemIds))
	r.SetRequestParam("DeliveryType", string(deliveryType))
	r.SetRequestParam("ShippingProvider", shippingProvider)
	r.SetRequestParam("TrackingNumber", trackingNumber)

	response, err := oe.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		rawHead := response.GetHead()
		errorMessage, err := jsonparser.GetString(rawHead, "ErrorMessage")

		if err != nil {
			return false, err
		}

		return false, errors.New(errorMessage)
	}

	return true, nil
}

func intSliceToParam(a []int) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("[%s]", strings.Join(b, ","))
}
