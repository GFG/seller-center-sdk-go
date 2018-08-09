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

type ScBool bool

func (t *ScBool) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	*t = ScBool("1" == raw)

	return nil
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

type ScTimestamp time.Time

func (t *ScTimestamp) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	if w, err := time.Parse(scTimeFormat, string(raw)); err != nil {
		return err
	} else {
		*t = ScTimestamp(w)
	}

	return nil
}

func (t ScTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))), nil
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
