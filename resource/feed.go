package resource

import (
	"encoding/json"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"github.com/buger/jsonparser"
)

type FeedResource struct {
	client client.Client
}

func NewFeed(client client.Client) FeedResource {
	return FeedResource{client: client}
}

func (fr FeedResource) FeedList() (model.FeedList, error) {
	request := client.NewGenericRequest("FeedList", client.MethodGET)
	request.SetVersion(client.V1)

	response, err := fr.client.Call(request)

	if err != nil {
		return model.FeedList{}, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)
		return model.FeedList{}, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	rawFeeds, _, _, err := jsonparser.Get(rawBody, "Feed")

	feedList := model.FeedList{}

	if err != nil {
		if err == jsonparser.KeyPathNotFoundError {
			return feedList, nil
		}

		return feedList, err
	}

	if len(rawFeeds) == 0 {
		return feedList, nil
	}

	err = json.Unmarshal(rawBody, &feedList)
	if err != nil {
		return model.FeedList{}, err
	}

	return feedList, nil
}

func (fr FeedResource) FeedStatus(feedIdentifier string) (model.FeedStatus, error) {
	request := client.NewGenericRequest("FeedStatus", client.MethodGET)
	request.SetVersion(client.V1)
	request.SetRequestParam("FeedID", feedIdentifier)

	response, err := fr.client.Call(request)

	feedStatus := model.FeedStatus{}

	if err != nil {
		return feedStatus, err
	}

	if response.IsError() {
		errorResponse, _ := response.(client.ErrorResponse)
		return feedStatus, newApiResponseError(errorResponse.HeadObject)
	}

	rawBody := response.GetBody()
	rawFeedStatus, _, _, err := jsonparser.Get(rawBody, "FeedDetail")

	if err != nil {
		if err == jsonparser.KeyPathNotFoundError {
			return feedStatus, nil
		}

		return feedStatus, err
	}

	if len(rawFeedStatus) == 0 {
		return feedStatus, nil
	}

	err = json.Unmarshal(rawFeedStatus, &feedStatus)
	if err != nil {
		return model.FeedStatus{}, err
	}

	return feedStatus, nil
}
