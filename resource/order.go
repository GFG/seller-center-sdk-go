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

type OrderResource struct {
	client client.Client
}

type GetOrdersParams struct {
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
	UpdatedAfter  *time.Time
	UpdatedBefore *time.Time
	Status        *string
	Limit         *int
	Offset        *int
	SortBy        *string
	SortDirection *string
}

func NewOrder(client client.Client) OrderResource {
	return OrderResource{client: client}
}

func (or OrderResource) GetOrders(params GetOrdersParams) (model.Orders, error) {

	r := client.NewGenericRequest("GetOrders", client.MethodGET)
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

	response, err := or.client.Call(r)

	if err != nil {
		return model.Orders{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Orders{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()

	var orders model.Orders
	err = json.Unmarshal(rawBody, &orders)

	return orders, err
}

func (or OrderResource) GetOrder(orderId int) (model.Order, error) {
	r := client.NewGenericRequest("GetOrder", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderId", strconv.Itoa(orderId))

	response, err := or.client.Call(r)

	if err != nil {
		return model.Order{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Order{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	var orders model.Orders
	if err := json.Unmarshal(rawBody, &orders); nil != err {
		return model.Order{}, err
	}

	if len(orders.Orders) == 1 {
		return orders.Orders[0], nil
	}

	return model.Order{}, errors.New("cannot find order")
}

func (or OrderResource) GetOrderItems(orderId int) (model.OrderItems, error) {
	r := client.NewGenericRequest("GetOrderItems", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderId", strconv.Itoa(orderId))

	response, err := or.client.Call(r)

	if err != nil {
		return model.OrderItems{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.OrderItems{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	var orderItems model.OrderItems
	err = json.Unmarshal(rawBody, &orderItems)
	if err != nil {
		return model.OrderItems{}, err
	}

	return orderItems, nil
}

func (or OrderResource) GetMultipleOrderItems(orderIds []int) (model.OrdersWithItems, error) {
	r := client.NewGenericRequest("GetMultipleOrderItems", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderIdList", intSliceToParam(orderIds))

	response, err := or.client.Call(r)

	if err != nil {
		return model.OrdersWithItems{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.OrdersWithItems{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	var ordersWithItems model.OrdersWithItems
	err = json.Unmarshal(rawBody, &ordersWithItems)
	if err != nil {
		return model.OrdersWithItems{}, err
	}

	return ordersWithItems, nil
}

func (or OrderResource) GetDocument(orderItemIds []int, documentType model.DocumentType) (model.Document, error) {
	r := client.NewGenericRequest("GetDocument", client.MethodGET)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemIds", intSliceToParam(orderItemIds))
	r.SetRequestParam("DocumentType", string(documentType))

	response, err := or.client.Call(r)

	if err != nil {
		return model.Document{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return model.Document{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	rawDocument, dataType, _, err := jsonparser.Get(rawBody, "Documents", "Document")

	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return model.Document{}, err
	}

	if len(rawDocument) == 0 || dataType == jsonparser.NotExist {
		return model.Document{}, errors.New("cannot find document")
	}

	document := model.Document{}
	err = json.Unmarshal(rawDocument, &document)
	if err != nil {
		return model.Document{}, err
	}

	return document, nil
}

func (or OrderResource) GetFailureReasons() (map[model.FailureReasonType][]string, error) {
	r := client.NewGenericRequest("GetFailureReasons", client.MethodGET)
	r.SetVersion(client.V1)

	response, err := or.client.Call(r)

	if err != nil {
		return map[model.FailureReasonType][]string{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return map[model.FailureReasonType][]string{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	var failureReasons model.FailureReasons
	err = json.Unmarshal(rawBody, &failureReasons)
	if err != nil {
		return map[model.FailureReasonType][]string{}, err
	}

	ret := make(map[model.FailureReasonType][]string)
	for _, reason := range failureReasons.Reasons {
		reasonType := model.FailureReasonType(reason.Type)
		if _, ok := ret[reasonType]; !ok {
			ret[reasonType] = make([]string, 0)
		}

		ret[reasonType] = append(ret[reasonType], reason.Name)
	}

	return ret, nil
}

func (or OrderResource) SetStatusToCanceled(orderItemId int, reason string, reasonDetail string) (bool, error) {
	r := client.NewGenericRequest("SetStatusToCanceled", client.MethodPOST)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemId", strconv.Itoa(orderItemId))
	r.SetRequestParam("Reason", reason)
	r.SetRequestParam("ReasonDetail", reasonDetail)

	response, err := or.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return false, newApiResponseError(errorResponse.HeadObject)
	}

	return true, nil
}

func (or OrderResource) SetStatusToPackedByMarketplace(orderItemIds []int, deliveryType model.DeliveryType, shippingProvider string) (bool, error) {
	r := client.NewGenericRequest("SetStatusToPackedByMarketplace", client.MethodPOST)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemIds", intSliceToParam(orderItemIds))
	r.SetRequestParam("DeliveryType", string(deliveryType))
	r.SetRequestParam("ShippingProvider", shippingProvider)

	response, err := or.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return false, newApiResponseError(errorResponse.HeadObject)
	}

	return true, nil
}

func (or OrderResource) SetStatusToReadyToShip(orderItemIds []int, deliveryType model.DeliveryType, shippingProvider string, trackingNumber string) (bool, error) {
	r := client.NewGenericRequest("SetStatusToReadyToShip", client.MethodPOST)
	r.SetVersion(client.V1)

	r.SetRequestParam("OrderItemIds", intSliceToParam(orderItemIds))
	r.SetRequestParam("DeliveryType", string(deliveryType))
	r.SetRequestParam("ShippingProvider", shippingProvider)
	r.SetRequestParam("TrackingNumber", trackingNumber)

	response, err := or.client.Call(r)

	if err != nil {
		return false, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)

		return false, newApiResponseError(errorResponse.HeadObject)
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
