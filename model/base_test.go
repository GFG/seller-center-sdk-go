package model

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func Test_ScFloat_NotEmpty(t *testing.T) {
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

func Test_ScFloat_Empty(t *testing.T) {
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

func Test_ScBool_False(t *testing.T) {
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

func Test_ScBool_True(t *testing.T) {
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

func Test_ScInt_NotEmpty(t *testing.T) {
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

func Test_ScInt_Empty(t *testing.T) {
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

func Test_ScTimestamp_NotEmpty(t *testing.T) {
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

func Test_ScTimestamp_Empty(t *testing.T) {
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
func Test_ScStringSlice_NotEmpty(t *testing.T) {
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

func Test_ScStringSlice_Empty(t *testing.T) {
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
