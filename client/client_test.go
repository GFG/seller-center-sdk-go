package client

import (
	"encoding/xml"
	"testing"
)

func Test_Can_Get_Request_Params_From_Generic_Request(t *testing.T) {
	action := "GetProducts"

	genericRequest := NewGenericRequest(action, MethodGET)

	requestParams := genericRequest.GetRequestParams()

	queryString := requestParams.Encode()

	expected := "Action=GetProducts&Format=JSON&Version=1.0"

	if queryString != expected {
		t.Fatalf("can not GetRequestParams. expected: `%s` - actual: `%s`.", expected, queryString)
	}
}

func Test_Can_Set_Version_Of_Generic_Request(t *testing.T) {
	action := "GetBrands"

	genericRequest := NewGenericRequest(action, MethodGET)

	genericRequest.SetVersion(V2)

	requestParams := genericRequest.GetRequestParams()

	queryString := requestParams.Encode()

	expected := "Action=GetBrands&Format=JSON&Version=2.0"

	if queryString != expected {
		t.Fatalf("can not GetRequestParams. expected: `%s` - actual: `%s`.", expected, queryString)
	}
}

func Test_Can_Set_Additional_Request_Params_To_Generic_Request(t *testing.T) {
	action := "GetBrands"

	genericRequest := NewGenericRequest(action, MethodGET)

	genericRequest.SetRequestParam("Whatever", "you want")
	genericRequest.SetRequestParam("Foo", "bar")

	requestParams := genericRequest.GetRequestParams()

	queryString := requestParams.Encode()

	expected := "Action=GetBrands&Foo=bar&Format=JSON&Version=1.0&Whatever=you+want"

	if queryString != expected {
		t.Fatalf("can not GetRequestParams. expected: `%s` - actual: `%s`.", expected, queryString)
	}
}

func Test_Can_Generate_Empty_Post_Xml_Generic_Request(t *testing.T) {
	action := "Whatever"

	genericRequest := NewGenericRequest(action, MethodPOST)

	xml, err := genericRequest.GeneratePostXml()

	if string(xml) != "" {
		t.Fatalf("can not get empty post xml, expected xml to be empty. actual `%s`", xml)
	}

	if err != nil {
		t.Fatal("can not get empty post xml, expected err to be nil.")
	}
}

func Test_Can_Generate_Post_Xml_Generic_Request(t *testing.T) {
	action := "Whatever"
	genericRequest := NewGenericRequest(action, MethodPOST)

	type myStruct struct {
		XMLName xml.Name `xml:"Node"`
		Id      int      `xml:"Id"`
		Msg     string   `xml:"Notification"`
	}

	data := myStruct{Id: 123, Msg: "Hello"}

	genericRequest.SetPostData(data)

	xml, err := genericRequest.GeneratePostXml()

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" + "\n" +
		"<Request>" + "\n" +
		"    <Node>" + "\n" +
		"        <Id>123</Id>" + "\n" +
		"        <Notification>Hello</Notification>" + "\n" +
		"    </Node>" + "\n" +
		"</Request>"

	if string(xml) != expected {
		t.Fatalf("can not get post xml, expected `%s`. actual `%s`", expected, xml)
	}

	if err != nil {
		t.Fatal("can not get post xml, expected err to be nil.")
	}
}
