package model

import (
	"github.com/buger/jsonparser"
	"strconv"
	"time"
)

type scFloat float64

func (f *scFloat) UnmarshalJSON(b []byte) error {
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
		*f = scFloat(w)
	}

	return nil
}

type scBool bool

func (t *scBool) UnmarshalJSON(b []byte) error {
	var raw, err = jsonparser.GetString(b)
	if err != nil {
		return err
	}

	*t = scBool("1" == raw)

	return nil
}

type scInt int

func (i *scInt) UnmarshalJSON(b []byte) error {
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
		*i = scInt(w)
	}

	return nil
}

type scTimestamp time.Time

func (t *scTimestamp) UnmarshalJSON(b []byte) error {
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
		*t = scTimestamp(w)
	}

	return nil
}
