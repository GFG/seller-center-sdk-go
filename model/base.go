package model

import (
	"encoding/xml"
	"fmt"
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
	"time"
)

type ScFloat float64

func (f *ScFloat) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	if w, err := strconv.ParseFloat(raw, 64); err != nil {
		return err
	} else {
		*f = ScFloat(w)
	}

	return nil
}

func (f ScFloat) MarshalJSON() ([]byte, error) {
	asString := fmt.Sprintf(`"%f"`, f)

	return []byte(asString), nil
}

type ScBool bool

func (t *ScBool) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	*t = ScBool("1" == raw)

	return nil
}

func (t ScBool) MarshalJSON() ([]byte, error) {
	asString := `"1"`

	if t == false {
		asString = `"0"`
	}

	return []byte(asString), nil
}

type ScInt int

func (i *ScInt) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	if w, err := strconv.Atoi(raw); err != nil {
		return err
	} else {
		*i = ScInt(w)
	}

	return nil
}

func (i ScInt) MarshalJSON() ([]byte, error) {
	asString := fmt.Sprintf(`"%d"`, i)

	return []byte(asString), nil
}

type ScTimestamp time.Time

func (t *ScTimestamp) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	if w, err := time.Parse(scTimeFormat, string(raw)); err == nil {
		*t = ScTimestamp(w)
	} else if w, err := time.Parse(time.RFC3339, string(raw)); err == nil {
		*t = ScTimestamp(w)
	} else {
		return err
	}

	return nil
}

func (t ScTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format(time.RFC3339))), nil
}

type ScIntSlice []int

func (i *ScIntSlice) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	var rawStrings = strings.Split(raw, ",")

	values := make(ScIntSlice, 0)
	for _, rawString := range rawStrings {
		if w, err := strconv.Atoi(rawString); err != nil {
			return err
		} else {
			values = append(values, w)
		}
	}

	*i = values

	return nil
}

func (i ScIntSlice) MarshalJSON() ([]byte, error) {
	asString := strings.Trim(strings.Replace(fmt.Sprint(i), " ", ",", -1), "[]")

	return []byte(fmt.Sprintf(`"%s"`, asString)), nil
}

type ScStringSlice []string

func (s *ScStringSlice) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	*s = ScStringSlice(strings.Split(raw, ","))

	return nil
}

func (s ScStringSlice) MarshalJSON() ([]byte, error) {
	asString := fmt.Sprintf(`"%s"`, strings.Join(s, ","))

	return []byte(asString), nil
}

type CharData string

func (cd CharData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		S string `xml:",cdata"`
	}{
		S: string(cd),
	}, start)
}

type IntSlice []int

func (is IntSlice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	intSliceString := ""
	if len(is) == 0 {
		return nil
	}

	b := make([]string, len(is))
	for i, v := range is {
		b[i] = strconv.Itoa(v)
	}

	intSliceString = strings.Join(b, ",")
	e.EncodeElement(intSliceString, start)

	return nil
}
