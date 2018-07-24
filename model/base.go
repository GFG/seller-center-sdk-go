package model

import (
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
