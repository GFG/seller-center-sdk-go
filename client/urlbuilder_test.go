package client

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

func createUrlBuilder(apiUrl string) ClientUrlBuilder {
	clientConfig := clientConfig{
		Url:  apiUrl,
		User: "abc@sellercenter.net",
		Key:  "1234567890",
	}

	str := "2014-11-12T11:45:26.371Z"
	now, _ := time.Parse(time.RFC3339, str)

	urlBuilder := clientUrlBuilder{
		config: clientConfig,
		hashHmacRequestSignature: hashHmacRequestSignature{
			key: clientConfig.Key,
		},
		datetimeProvider: datetimeProvider{
			now: now,
		},
	}

	return urlBuilder
}

func Test_Can_Build_Url(t *testing.T) {
	urlBuilder := createUrlBuilder("https://my-api.sc.net/")

	requestParams := url.Values{}
	requestParams.Add("Foo", "bar")
	requestParams.Add("Whatever", "you want")

	url, err := urlBuilder.BuildUrl(requestParams)

	expectedUrl := "https://my-api.sc.net/?Foo=bar&Signature=421768b44097516adfde6f2331e1dd609a3070e9007fe06736aef8f7de55182f&Timestamp=2014-11-12T11%3A45%3A26Z&UserID=abc%40sellercenter.net&Whatever=you+want"

	if url != expectedUrl {
		t.Fatalf("can not build url. expected: `%s` - build url: `%s`.", expectedUrl, url)
	}

	if err != nil {
		t.Fatal("can not build url. expected error to be nil.")
	}
}

func Test_Build_Will_Fail_With_Wrong_Url(t *testing.T) {
	urlBuilder := createUrlBuilder("")

	requestParams := url.Values{}
	requestParams.Add("Foo", "bar")
	requestParams.Add("Whatever", "you want")

	url, err := urlBuilder.BuildUrl(requestParams)

	if url != "" {
		t.Fatalf("expected url to be empty string. but build url was: `%s`.", url)
	}

	if reflect.TypeOf(err).String() != "*url.Error" {
		t.Fatal("expected to have an url.Error.")
	}

}
