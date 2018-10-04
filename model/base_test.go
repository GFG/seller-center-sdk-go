package model

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

func Test_Unmarshal_ScFloat_NotEmpty(t *testing.T) {
	j := []byte("{\"ScFloat\": \"12.34\"}")

	type s struct {
		ScFloat ScFloat `json:"ScFloat"`
	}

	expected := ScFloat(12.34)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%f` - error: `%s`.", expected, err)
	}

	if expected != c.ScFloat {
		t.Fatalf("unmarshalled doesn't match. expected: `%f` - unmarshalled: `%f`.", expected, c.ScFloat)
	}
}

func Test_Marshal_ScFloat_NotEmpty(t *testing.T) {
	j := ScFloat(12.34)

	expected := []byte(`"12.340000"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScFloat_Empty(t *testing.T) {
	j := []byte("{\"ScFloat\": \"\"}")

	type s struct {
		ScFloat ScFloat `json:"ScFloat"`
	}

	expected := ScFloat(0.0)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%f` - error: `%s`.", expected, err)
	}

	if expected != c.ScFloat {
		t.Fatalf("unmarshalled doesn't match. expected: `%f` - unmarshalled: `%f`.", expected, c.ScFloat)
	}
}

func Test_Marshal_ScFloat_Empty(t *testing.T) {
	j := ScFloat(0.0)

	expected := []byte(`"0.000000"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScBool_False(t *testing.T) {
	j := []byte("{\"ScBool\": \"0\"}")

	type s struct {
		ScBool ScBool `json:"ScBool"'`
	}

	expected := ScBool(false)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%t` - error: `%s`.", expected, err)
	}

	if expected != c.ScBool {
		t.Fatalf("unmarshalled doesn't match. expected: `%t` - unmarshalled: `%t`.", expected, c.ScBool)
	}
}

func Test_Marshal_ScBool_False(t *testing.T) {
	j := ScBool(false)

	expected := []byte(`"0"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScBool_True(t *testing.T) {
	j := []byte("{\"ScBool\": \"1\"}")

	type s struct {
		ScBool ScBool `json:"ScBool"'`
	}

	expected := ScBool(true)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%t` - error: `%s`.", expected, err)
	}

	if expected != c.ScBool {
		t.Fatalf("unmarshalled doesn't match. expected: `%t` - unmarshalled: `%t`.", expected, c.ScBool)
	}
}

func Test_Marshal_ScBool_True(t *testing.T) {
	j := ScBool(true)

	expected := []byte(`"1"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScInt_NotEmpty(t *testing.T) {
	j := []byte("{\"ScInt\": \"12\"}")

	type s struct {
		ScInt ScInt `json:"ScInt"'`
	}

	expected := ScInt(12)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%d` - error: `%s`.", expected, err)
	}

	if expected != c.ScInt {
		t.Fatalf("unmarshalled doesn't match. expected: `%d` - unmarshalled: `%d`.", expected, c.ScInt)
	}
}

func Test_Marshal_ScInt_NotEmpty(t *testing.T) {
	j := ScInt(12)

	expected := []byte(`"12"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScInt_Empty(t *testing.T) {
	j := []byte("{\"ScInt\": \"\"}")

	type s struct {
		ScInt ScInt `json:"ScInt"'`
	}

	expected := ScInt(0)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%d` - error: `%s`.", expected, err)
	}

	if expected != c.ScInt {
		t.Fatalf("unmarshalled doesn't match. expected: `%d` - unmarshalled: `%d`.", expected, c.ScInt)
	}
}

func Test_Marshal_ScInt_Empty(t *testing.T) {
	j := ScInt(0)

	expected := []byte(`"0"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScTimestamp_NotEmpty_ScFormat(t *testing.T) {
	j := []byte("{\"ScTimestamp\": \"2018-07-10 14:26:20\"}")

	type s struct {
		ScTimestamp ScTimestamp `json:"ScTimestamp"'`
	}

	expected := ScTimestamp(time.Date(2018, 7, 10, 14, 26, 20, 0, time.UTC))

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if expected != c.ScTimestamp {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScTimestamp)
	}
}

func Test_Unmarshal_ScTimestamp_NotEmpty_UnmarshaledFormat(t *testing.T) {
	j := []byte("{\"ScTimestamp\": \"2018-07-10T14:26:20Z\"}")

	type s struct {
		ScTimestamp ScTimestamp `json:"ScTimestamp"'`
	}

	expected := ScTimestamp(time.Date(2018, 7, 10, 14, 26, 20, 0, time.UTC))

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if expected != c.ScTimestamp {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScTimestamp)
	}
}

func Test_Marshal_ScTimestamp_NotEmpty(t *testing.T) {
	j := ScTimestamp(time.Date(2018, 7, 10, 14, 26, 20, 0, time.UTC))

	expected := []byte(`"2018-07-10T14:26:20Z"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScTimestamp_Empty(t *testing.T) {
	j := []byte("{\"ScTimestamp\": \"\"}")

	type s struct {
		ScTimestamp ScTimestamp `json:"ScTimestamp"'`
	}

	expected := ScTimestamp(time.Time{})

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if expected != c.ScTimestamp {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScTimestamp)
	}
}

func Test_Marshal_ScTimestamp_Empty(t *testing.T) {
	j := ScTimestamp(time.Time{})

	expected := []byte(`"0001-01-01T00:00:00Z"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScStringSlice_NotEmpty(t *testing.T) {
	j := []byte("{\"ScStringSlice\": \"A,B\"}")

	type s struct {
		ScStringSlice ScStringSlice `json:"ScStringSlice"'`
	}

	expected := ScStringSlice([]string{"A", "B"})

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c.ScStringSlice) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScStringSlice)
	}
}

func Test_Marshal_ScStringSlice_NotEmpty(t *testing.T) {
	j := ScStringSlice([]string{"A", "B"})

	expected := []byte(`"A,B"`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_Unmarshal_ScStringSlice_Empty(t *testing.T) {
	j := []byte("{\"ScStringSlice\": \"\"}")

	type s struct {
		ScStringSlice ScStringSlice `json:"ScStringSlice"'`
	}

	expected := ScStringSlice(nil)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if !reflect.DeepEqual(expected, c.ScStringSlice) {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScStringSlice)
	}
}

func Test_Marshal_ScStringSlice_Empty(t *testing.T) {
	j := ScStringSlice(nil)

	expected := []byte(`""`)
	if actual, err := json.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_MarshalXML_CharData_NotEmpty(t *testing.T) {
	j := CharData("test char data")

	expected := []byte(`<CharData><![CDATA[test char data]]></CharData>`)
	if actual, err := xml.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_MarshalXML_CharData_Empty(t *testing.T) {
	j := CharData("")

	expected := []byte(`<CharData></CharData>`)
	if actual, err := xml.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_MarshalXML_IntSlice_NotEmpty(t *testing.T) {
	j := IntSlice([]int{1, 2})

	expected := []byte(`<IntSlice>1,2</IntSlice>`)
	if actual, err := xml.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}

func Test_MarshalXML_IntSlice_Empty(t *testing.T) {
	j := IntSlice([]int{})

	expected := []byte(nil)
	if actual, err := xml.Marshal(j); nil != err {
		t.Fatalf("can not unmarshal. expected: `%s` - error: `%s`.", expected, err)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unmarshalled doesn't match. expected: `%s` - unmarshalled: `%s`.", expected, actual)
	}
}
