package model

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func Test_FeedList_With_FailureReports(t *testing.T) {
	payloadBody := []byte(`{"Feed":[{"Feed":"83988c5e-c67c-41a8-ae95-0ac21f32fae7","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-24 12:05:05","UpdatedDate":"2018-07-24 12:05:06","Source":"api","TotalRecords":"3","ProcessedRecords":"1","FailedRecords":"2","FailureReports":""},{"Feed":"89e767bc-bf18-4f92-88a9-24368bd6a08c","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-24 11:34:53","UpdatedDate":"2018-07-24 11:34:53","Source":"api","TotalRecords":"1","ProcessedRecords":"1","FailedRecords":"1","FailureReports":{"MimeType":"text/csv","File":"IkVycm9yIjsiV2FybmluZyI7IlNlbGxlclNrdSI7Ik5hbWUiOyJEZXNjcmlwdGlvbiI7IkJyYW5kIjsiVGF4Q2xhc3MiOyJWYXJpYXRpb24iOyJQYXJlbnRTa3UiOyJRdWFudGl0eSI7IlByaWNlIjsiU2FsZVByaWNlIjsiU2FsZVN0YXJ0RGF0ZSI7IlNhbGVFbmREYXRlIjsiU3RhdHVzIjsiUHJvZHVjdElkIjsiVm9sdW1ldHJpY1dlaWdodCI7IlByb2R1Y3RHcm91cCI7Ik1haW5JbWFnZSI7IjAiOyIxIjsiUHJpbWFyeUNhdGVnb3J5IjsiQ2F0ZWdvcmllcyI7IkRlc2NyaXB0aW9uRW4iOyJCcm93c2VOb2RlcyI7IlNoaXBtZW50VHlwZSI7IkNvbmRpdGlvbiIKIkZpZWxkKHMpIFZvbHVtZXRyaWNXZWlnaHQgaXMvYXJlIGludmFsaWQiOyIiOyJTZWxsZXIgU2t1IjsiTmFtZSI7IlRoaXMgaXMgYSA8Yj5ib2xkPC9iPiBwcm9kdWN0LiI7IkJyYW5kIjsiZGVmYXVsdCI7IlhYUyI7IlBhcmVudCBTa3UiOyI0IjsiNDAiOyIzMyI7IjIwMTUtMTEtMDQgMTA6MzA6NDkiOyIyMDE1LTExLTA5IDEwOjMwOjQ5IjsiYWN0aXZlIjsiUHJvZHVjdCBJZCI7IjEwLjU1IjsicHJvZHVjdCBncm91cCI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UxLmpwZyI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UyLmpwZyI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UzLmpwZyI7IjEiOyIyLDMiOyJJIGFtIGEgZGVzY3JpcHRpb24gZm9yIHRoZSBuZXcgcHJvZHVjdCBhZ2FpbiI7IjUsNiI7ImNyb3NzZG9ja2luZyI7Im5ldyIK"}}]}`)

	decodedFile, _ := base64.StdEncoding.DecodeString("IkVycm9yIjsiV2FybmluZyI7IlNlbGxlclNrdSI7Ik5hbWUiOyJEZXNjcmlwdGlvbiI7IkJyYW5kIjsiVGF4Q2xhc3MiOyJWYXJpYXRpb24iOyJQYXJlbnRTa3UiOyJRdWFudGl0eSI7IlByaWNlIjsiU2FsZVByaWNlIjsiU2FsZVN0YXJ0RGF0ZSI7IlNhbGVFbmREYXRlIjsiU3RhdHVzIjsiUHJvZHVjdElkIjsiVm9sdW1ldHJpY1dlaWdodCI7IlByb2R1Y3RHcm91cCI7Ik1haW5JbWFnZSI7IjAiOyIxIjsiUHJpbWFyeUNhdGVnb3J5IjsiQ2F0ZWdvcmllcyI7IkRlc2NyaXB0aW9uRW4iOyJCcm93c2VOb2RlcyI7IlNoaXBtZW50VHlwZSI7IkNvbmRpdGlvbiIKIkZpZWxkKHMpIFZvbHVtZXRyaWNXZWlnaHQgaXMvYXJlIGludmFsaWQiOyIiOyJTZWxsZXIgU2t1IjsiTmFtZSI7IlRoaXMgaXMgYSA8Yj5ib2xkPC9iPiBwcm9kdWN0LiI7IkJyYW5kIjsiZGVmYXVsdCI7IlhYUyI7IlBhcmVudCBTa3UiOyI0IjsiNDAiOyIzMyI7IjIwMTUtMTEtMDQgMTA6MzA6NDkiOyIyMDE1LTExLTA5IDEwOjMwOjQ5IjsiYWN0aXZlIjsiUHJvZHVjdCBJZCI7IjEwLjU1IjsicHJvZHVjdCBncm91cCI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UxLmpwZyI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UyLmpwZyI7Imh0dHBzOi8vc2VsbGVyYXBpLnNlbGxlcmNlbnRlci5uZXQvaW1hZ2UzLmpwZyI7IjEiOyIyLDMiOyJJIGFtIGEgZGVzY3JpcHRpb24gZm9yIHRoZSBuZXcgcHJvZHVjdCBhZ2FpbiI7IjUsNiI7ImNyb3NzZG9ja2luZyI7Im5ldyIK")

	expected := FeedList{
		[]Feed{
			{
				"83988c5e-c67c-41a8-ae95-0ac21f32fae7",
				"Finished",
				"ProductCreate",
				ScTimestamp(time.Date(2018, 7, 24, 12, 5, 5, 00, time.UTC)),
				ScTimestamp(time.Date(2018, 7, 24, 12, 5, 6, 00, time.UTC)),
				"api",
				ScInt(3),
				ScInt(1),
				ScInt(2),
				FeedFailureReportFile{},
			},
			{
				"89e767bc-bf18-4f92-88a9-24368bd6a08c",
				"Finished",
				"ProductCreate",
				ScTimestamp(time.Date(2018, 7, 24, 11, 34, 53, 00, time.UTC)),
				ScTimestamp(time.Date(2018, 7, 24, 11, 34, 53, 00, time.UTC)),
				"api",
				ScInt(1),
				ScInt(1),
				ScInt(1),
				FeedFailureReportFile{
					"text/csv",
					string(decodedFile),
				},
			},
		},
	}

	var feedList FeedList
	if err := json.Unmarshal(payloadBody, &feedList); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedList) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedList)
	}
}

func Test_FeedList_With_One_Item(t *testing.T) {
	payloadBody := []byte(`{"Feed":{"Feed":"992955e2-af2d-4d23-9b5a-91230a86b50d","Status":"Queued","Action":"ProductCreate","CreationDate":"2018-07-25 14:13:15","UpdatedDate":"2018-07-25 14:13:15","Source":"api","TotalRecords":"1","ProcessedRecords":"0","FailedRecords":"0","FailureReports":""}}`)

	expected := FeedList{
		[]Feed{
			{
				"992955e2-af2d-4d23-9b5a-91230a86b50d",
				"Queued",
				"ProductCreate",
				ScTimestamp(time.Date(2018, 7, 25, 14, 13, 15, 00, time.UTC)),
				ScTimestamp(time.Date(2018, 7, 25, 14, 13, 15, 00, time.UTC)),
				"api",
				ScInt(1),
				ScInt(0),
				ScInt(0),
				FeedFailureReportFile{},
			},
		},
	}

	var feedList FeedList
	if err := json.Unmarshal(payloadBody, &feedList); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedList) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedList)
	}
}

func Test_FeedStatus_Without_Error_Or_Warning(t *testing.T) {
	rawFeedStatus := []byte(`{"Feed":"fb407b7a-93a8-42d1-9a58-797f4ccac3d6","Status":"Queued","Action":"ProductCreate","CreationDate":"2018-07-26 10:43:17","UpdatedDate":"2018-07-26 10:43:17","Source":"api","TotalRecords":"1","ProcessedRecords":"0","FailedRecords":"0"}`)

	expected := FeedStatus{
		"fb407b7a-93a8-42d1-9a58-797f4ccac3d6",
		"Queued",
		"ProductCreate",
		ScTimestamp(time.Date(2018, 7, 26, 10, 43, 17, 00, time.UTC)),
		ScTimestamp(time.Date(2018, 7, 26, 10, 43, 17, 00, time.UTC)),
		"api",
		ScInt(1),
		ScInt(0),
		ScInt(0),
		FeedErrors{},
		FeedWarnings{},
	}

	var feedStatus FeedStatus
	if err := json.Unmarshal(rawFeedStatus, &feedStatus); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedStatus) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedStatus)
	}
}

func Test_FeedStatus_With_One_Error(t *testing.T) {
	rawFeedStatus := []byte(`{"Feed":"89e767bc-bf18-4f92-88a9-24368bd6a08c","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-24 11:34:53","UpdatedDate":"2018-07-24 11:34:53","Source":"api","TotalRecords":"3","ProcessedRecords":"2","FailedRecords":"1","FeedErrors":{"Error":{"Code":"0","Message":"Field(s) VolumetricWeight is/are invalid","SellerSku":"Seller Sku"}}}`)

	expected := FeedStatus{
		"89e767bc-bf18-4f92-88a9-24368bd6a08c",
		"Finished",
		"ProductCreate",
		ScTimestamp(time.Date(2018, 7, 24, 11, 34, 53, 00, time.UTC)),
		ScTimestamp(time.Date(2018, 7, 24, 11, 34, 53, 00, time.UTC)),
		"api",
		ScInt(3),
		ScInt(2),
		ScInt(1),
		FeedErrors{
			[]FeedError{
				{
					"0",
					"Field(s) VolumetricWeight is/are invalid",
					"Seller Sku",
					"",
					"",
				},
			},
		},
		FeedWarnings{},
	}

	var feedStatus FeedStatus
	if err := json.Unmarshal(rawFeedStatus, &feedStatus); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedStatus) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedStatus)
	}
}

func Test_FeedStatus_With_Multiple_Errors(t *testing.T) {
	rawFeedStatus := []byte(`{"Feed":"992955e2-af2d-4d23-9b5a-91230a86b50d","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-25 14:13:15","UpdatedDate":"2018-07-25 15:11:17","Source":"api","TotalRecords":"1","ProcessedRecords":"1","FailedRecords":"1","FeedErrors":{"Error":[{"Code":"0","Message":"2438 is not a valid category id.","SellerSku":"1234567890"},{"Code":"1","Message":"Price is lower than category min price: 50","SellerSku":"1234567890"},{"Code":"2","Message":"MaterialFilling, Season, Subset is/are not applicable to","SellerSku":"1234567890"}]}}`)

	expected := FeedStatus{
		"992955e2-af2d-4d23-9b5a-91230a86b50d",
		"Finished",
		"ProductCreate",
		ScTimestamp(time.Date(2018, 7, 25, 14, 13, 15, 00, time.UTC)),
		ScTimestamp(time.Date(2018, 7, 25, 15, 11, 17, 00, time.UTC)),
		"api",
		ScInt(1),
		ScInt(1),
		ScInt(1),
		FeedErrors{
			[]FeedError{
				{
					"0",
					"2438 is not a valid category id.",
					"1234567890",
					"",
					"",
				},
				{
					"1",
					"Price is lower than category min price: 50",
					"1234567890",
					"",
					"",
				},
				{
					"2",
					"MaterialFilling, Season, Subset is/are not applicable to",
					"1234567890",
					"",
					"",
				},
			},
		},
		FeedWarnings{},
	}

	var feedStatus FeedStatus
	if err := json.Unmarshal(rawFeedStatus, &feedStatus); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedStatus) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedStatus)
	}
}

func Test_FeedStatus_With_One_Warning(t *testing.T) {
	rawFeedStatus := []byte(`{"Feed":"992955e2-af2d-4d23-9b5a-91230a86b50d","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-25 14:13:15","UpdatedDate":"2018-07-25 15:11:17","Source":"api","TotalRecords":"1","ProcessedRecords":"1","FailedRecords":"1","FeedWarnings":{"Warning":{"Message":"2438 is not a valid category id.","SellerSku":"1234567890"}}}`)

	expected := FeedStatus{
		"992955e2-af2d-4d23-9b5a-91230a86b50d",
		"Finished",
		"ProductCreate",
		ScTimestamp(time.Date(2018, 7, 25, 14, 13, 15, 00, time.UTC)),
		ScTimestamp(time.Date(2018, 7, 25, 15, 11, 17, 00, time.UTC)),
		"api",
		ScInt(1),
		ScInt(1),
		ScInt(1),
		FeedErrors{},
		FeedWarnings{
			[]FeedWarning{
				{
					"2438 is not a valid category id.", //fake warning - we only test the structure
					"1234567890",
					"",
					"",
				},
			},
		},
	}

	var feedStatus FeedStatus
	if err := json.Unmarshal(rawFeedStatus, &feedStatus); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedStatus) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedStatus)
	}
}

func Test_FeedStatus_With_Multiple_Warnings(t *testing.T) {
	rawFeedStatus := []byte(`{"Feed":"992955e2-af2d-4d23-9b5a-91230a86b50d","Status":"Finished","Action":"ProductCreate","CreationDate":"2018-07-25 14:13:15","UpdatedDate":"2018-07-25 15:11:17","Source":"api","TotalRecords":"1","ProcessedRecords":"1","FailedRecords":"1","FeedWarnings":{"Warning":[{"Message":"2438 is not a valid category id.","SellerSku":"1234567890"},{"Message":"Price is lower than category min price: 50","SellerSku":"1234567890"},{"Message":"MaterialFilling, Season, Subset is/are not applicable to","SellerSku":"1234567890"}]}}`)

	expected := FeedStatus{
		"992955e2-af2d-4d23-9b5a-91230a86b50d",
		"Finished",
		"ProductCreate",
		ScTimestamp(time.Date(2018, 7, 25, 14, 13, 15, 00, time.UTC)),
		ScTimestamp(time.Date(2018, 7, 25, 15, 11, 17, 00, time.UTC)),
		"api",
		ScInt(1),
		ScInt(1),
		ScInt(1),
		FeedErrors{},
		FeedWarnings{
			[]FeedWarning{
				{
					"2438 is not a valid category id.", //fake warning - we only test the structure
					"1234567890",
					"",
					"",
				},
				{
					"Price is lower than category min price: 50",
					"1234567890",
					"",
					"",
				},
				{
					"MaterialFilling, Season, Subset is/are not applicable to",
					"1234567890",
					"",
					"",
				},
			},
		},
	}

	var feedStatus FeedStatus
	if err := json.Unmarshal(rawFeedStatus, &feedStatus); nil != err {
		t.Fatalf("can not unmarshal. expected:`%v` - error:`%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, feedStatus) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, feedStatus)
	}
}
