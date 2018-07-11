package model

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_scFloat_NotEmpty(t *testing.T) {
	j := []byte("{\"scFloat\": \"12.34\"}")

	type s struct {
		ScFloat scFloat `json:"scFloat"`
	}

	expected := scFloat(12.34)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%f` - error: `%s`.", expected, err)
	}

	if expected != c.ScFloat {
		t.Fatalf("unmarshalled doesn't match. expected: `%f` - unmarshalled: `%f`.", expected, c.ScFloat)
	}
}

func Test_scFloat_Empty(t *testing.T) {
	j := []byte("{\"scFloat\": \"\"}")

	type s struct {
		ScFloat scFloat `json:"scFloat"`
	}

	expected := scFloat(0.0)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%f` - error: `%s`.", expected, err)
	}

	if expected != c.ScFloat {
		t.Fatalf("unmarshalled doesn't match. expected: `%f` - unmarshalled: `%f`.", expected, c.ScFloat)
	}
}

func Test_scBool_False(t *testing.T) {
	j := []byte("{\"scBool\": \"0\"}")

	type s struct {
		ScBool scBool `json:"scBool"'`
	}

	expected := scBool(false)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%t` - error: `%s`.", expected, err)
	}

	if expected != c.ScBool {
		t.Fatalf("unmarshalled doesn't match. expected: `%t` - unmarshalled: `%t`.", expected, c.ScBool)
	}
}

func Test_scBool_True(t *testing.T) {
	j := []byte("{\"scBool\": \"1\"}")

	type s struct {
		ScBool scBool `json:"scBool"'`
	}

	expected := scBool(true)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%t` - error: `%s`.", expected, err)
	}

	if expected != c.ScBool {
		t.Fatalf("unmarshalled doesn't match. expected: `%t` - unmarshalled: `%t`.", expected, c.ScBool)
	}
}

func Test_scInt_NotEmpty(t *testing.T) {
	j := []byte("{\"scInt\": \"12\"}")

	type s struct {
		ScInt scInt `json:"scInt"'`
	}

	expected := scInt(12)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%d` - error: `%s`.", expected, err)
	}

	if expected != c.ScInt {
		t.Fatalf("unmarshalled doesn't match. expected: `%d` - unmarshalled: `%d`.", expected, c.ScInt)
	}
}

func Test_scInt_Empty(t *testing.T) {
	j := []byte("{\"scInt\": \"\"}")

	type s struct {
		ScInt scInt `json:"scInt"'`
	}

	expected := scInt(0)

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%d` - error: `%s`.", expected, err)
	}

	if expected != c.ScInt {
		t.Fatalf("unmarshalled doesn't match. expected: `%d` - unmarshalled: `%d`.", expected, c.ScInt)
	}
}

func Test_scTimestamp_NotEmpty(t *testing.T) {
	j := []byte("{\"scTimestamp\": \"2018-07-10 14:26:20\"}")

	type s struct {
		ScTimestamp scTimestamp `json:"scTimestamp"'`
	}

	expected := scTimestamp(time.Date(2018, 7, 10, 14, 26, 20, 0, time.UTC))

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if expected != c.ScTimestamp {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScTimestamp)
	}
}

func Test_scTimestamp_Empty(t *testing.T) {
	j := []byte("{\"scTimestamp\": \"\"}")

	type s struct {
		ScTimestamp scTimestamp `json:"scTimestamp"'`
	}

	expected := scTimestamp(time.Time{})

	var c s
	if err := json.Unmarshal(j, &c); nil != err {
		t.Fatalf("can not unmarshal. expected: `%v` - error: `%s`.", expected, err)
	}

	if expected != c.ScTimestamp {
		t.Fatalf("unmarshalled doesn't match. expected: `%v` - unmarshalled: `%v`.", expected, c.ScTimestamp)
	}
}
