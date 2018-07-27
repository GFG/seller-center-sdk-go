package resource

import (
	"errors"
	"github.com/GFG/seller-center-sdk-go/client"
	"github.com/GFG/seller-center-sdk-go/model"
	"reflect"
	"testing"
)

func Test_Get_FeedList_Returns_Client_Error(t *testing.T) {
	clientError := errors.New("Timeout")

	fakeClient := client.FakeClient{
		FakeResponse: nil,
		FakeError:    clientError,
	}

	resource := NewFeed(fakeClient)

	feedList, err := resource.FeedList()

	if !reflect.DeepEqual(clientError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", clientError, err)
	}

	empty := model.FeedList{}

	if !reflect.DeepEqual(empty, feedList) {
		t.Fatalf("did not return empty feedList. received: `%v`.", feedList)
	}
}

func Test_Get_FeedList_Can_Handle_ErrorResponse(t *testing.T) {
	clientErrorResponse := client.ErrorResponse{
		HeadObject: client.HeadErrorResponse{
			ErrorCode:    "E27",
			ErrorMessage: "Wrong signature",
		},
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientErrorResponse,
		FakeError:    nil,
	}

	resource := NewFeed(fakeClient)

	feedList, err := resource.FeedList()

	expectedError := &ApiResponseError{
		"E27",
		"Wrong signature",
	}

	if !reflect.DeepEqual(expectedError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", expectedError, err)
	}

	empty := model.FeedList{}

	if !reflect.DeepEqual(empty, feedList) {
		t.Fatalf("did not return empty feedList. received: `%v`.", feedList)
	}
}

func Test_Get_FeedList_Can_Handle_EmptyResponseBody(t *testing.T) {
	clientResponse := client.SuccessResponse{
		Body: []byte(``),
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewFeed(fakeClient)

	feedList, err := resource.FeedList()

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	empty := model.FeedList{}

	if !reflect.DeepEqual(empty, feedList) {
		t.Fatalf("did not return empty feedList. received: `%v`.", feedList)
	}
}

func Test_Get_FeedList_Can_Handle_Response(t *testing.T) {
	payloadBody := []byte(`{"Feed":[{"Feed":"83988c5e-c67c-41a8-ae95-0ac21f32fae7","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-24 12:05:05","UpdatedDate":"2018-07-24 12:05:06","Source":"api","TotalRecords":"3","ProcessedRecords":"1","FailedRecords":"2","FailureReports":""},{"Feed":"89e767bc-bf18-4f92-88a9-24368bd6a08c","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-24 11:34:53","UpdatedDate":"2018-07-24 11:34:53","Source":"api","TotalRecords":"1","ProcessedRecords":"1","FailedRecords":"1","FailureReports":{"MimeType":"text/csv","File":"IkVycm9yIjsiV2FybmluZyI7IlNlbGxlclNrdSI7Ik5hbWUiOyJEZXNjcmlwdGlvbiI7IkJyYW5kIjsiVGF4Q2xhc3MiOyJWYXJpYXRpb24iOyJQYXJlbnRTa3UiOyJRdWFudGl0eSI7IlByaWNlIjsiU2FsZVByaWNlIjsiU2FsZVN0YXJ0RGF0ZSI7IlNhbGVFbmREYXRlIjsiU3RhdHVzIjsiUHJvZHVjdElkIjsiVm9sdW1ldHJpY1dlaWdodCI7IlByb2R1Y3RHcm91cCI7Ik1haW5JbWFnZSI7IjAiOyIxIjsiUHJpbWFyeUNhdGVnb3J5IjsiQ2F0ZWdvcmllcyI7IkRlc2NyaXB0aW9uRW4iOyJCcm93c2VOb2RlcyI7IlNoaXBtZW50VHlwZSI7IkNvbmRpdGlvbiIKIkZpZWxkKHMpIFZvbHVtZXRyaWNXZWlnaHQgaXMvYXJlIGludmFsaWQiOyIiOyJTZWxsZXIgU2t1IjsiTmFtZSI7IlRoaXMgaXMgYSA8Yj5ib2xkPC9iPiBwcm9kdWN0LiI7IkJyYW5kIjsiZGVmYXVsdCI7IlhYUyI7IlBhcmVudCBTa3UiOyI0IjsiNDAiOyIzMyI7IjIwMTUtMTEtMDQgMTA6MzA6NDkiOyIyMDE1LTExLTA5IDEwOjMwOjQ5IjsiYWN0aXZlIjsiUHJvZHVjdCBJZCI7IjEwLjU1IjsicHJvZHVjdCBncm91cCI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UxLmpwZyI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UyLmpwZyI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UzLmpwZyI7IjEiOyIyLDMiOyJJIGFtIGEgZGVzY3JpcHRpb24gZm9yIHRoZSBuZXcgcHJvZHVjdCBhZ2FpbiI7IjUsNiI7ImNyb3NzZG9ja2luZyI7Im5ldyIK"}}]}`)

	clientResponse := client.SuccessResponse{
		Body: payloadBody,
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewFeed(fakeClient)

	feedList, err := resource.FeedList()

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	if len(feedList.Feeds) != 2 {
		t.Fatalf("did not receive feedList with 2 items. received: `%d`.", len(feedList.Feeds))
	}
}

func Test_Get_FeedStatus_Returns_Client_Error(t *testing.T) {
	clientError := errors.New("Timeout")

	fakeClient := client.FakeClient{
		FakeResponse: nil,
		FakeError:    clientError,
	}

	resource := NewFeed(fakeClient)

	feedStatus, err := resource.FeedStatus("1234")

	if !reflect.DeepEqual(clientError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", clientError, err)
	}

	empty := model.FeedStatus{}

	if !reflect.DeepEqual(empty, feedStatus) {
		t.Fatalf("did not return empty feedStatus. received: `%v`.", feedStatus)
	}
}

func Test_Get_FeedStatus_Can_Handle_ErrorResponse(t *testing.T) {
	clientErrorResponse := client.ErrorResponse{
		HeadObject: client.HeadErrorResponse{
			ErrorCode:    "E27",
			ErrorMessage: "Wrong signature",
		},
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientErrorResponse,
		FakeError:    nil,
	}

	resource := NewFeed(fakeClient)

	feedList, err := resource.FeedStatus("1234")

	expectedError := &ApiResponseError{
		"E27",
		"Wrong signature",
	}

	if !reflect.DeepEqual(expectedError, err) {
		t.Fatalf("client error was not returned. expected: `%v` - received: `%v`.", expectedError, err)
	}

	empty := model.FeedStatus{}

	if !reflect.DeepEqual(empty, feedList) {
		t.Fatalf("did not return empty feedStatus. received: `%v`.", feedList)
	}
}

func Test_Get_FeedStatus_Can_Handle_EmptyResponseBody(t *testing.T) {
	clientResponse := client.SuccessResponse{
		Body: []byte(``),
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewFeed(fakeClient)

	feedStatus, err := resource.FeedStatus("1234")

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	empty := model.FeedStatus{}

	if !reflect.DeepEqual(empty, feedStatus) {
		t.Fatalf("did not return empty feedStatus. received: `%v`.", feedStatus)
	}
}

func Test_Get_FeedStatus_Can_Handle_Response(t *testing.T) {
	payloadBody := []byte(`{"FeedDetail":{"Feed":"992955e2-af2d-4d23-9b5a-91230a86b50d","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-25 14:13:15","UpdatedDate":"2018-07-25 15:11:17","Source":"api","TotalRecords":"1","ProcessedRecords":"1","FailedRecords":"1","FeedWarnings":{"Warning":{"Message":"2438 is not a valid category id.","SellerSku":"1234567890"}}}}`)

	clientResponse := client.SuccessResponse{
		Body: payloadBody,
	}

	fakeClient := client.FakeClient{
		FakeResponse: clientResponse,
		FakeError:    nil,
	}

	resource := NewFeed(fakeClient)

	feedStatus, err := resource.FeedStatus("1234")

	if err != nil {
		t.Fatalf("client error was expected to be nil. received: `%v`.", err)
	}

	// just check 2 attributes as the json decode is already tested in the model test
	if feedStatus.Feed != "992955e2-af2d-4d23-9b5a-91230a86b50d" || feedStatus.Status != "Finished" {
		t.Fatalf("did not receive proper feedStatus. received: `%v`.", feedStatus)
	}
}
