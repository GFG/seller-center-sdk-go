package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
	"strings"
)

type FeedList struct {
	Feeds []Feed `json:"Feeds"`
}

func (fl *FeedList) UnmarshalJSON(b []byte) error {
	rawFeeds, dataType, _, err := jsonparser.Get(b, "Feed")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(rawFeeds) == 0 || dataType == jsonparser.NotExist {
		*fl = FeedList{[]Feed{}}
		return nil
	}

	var feeds []Feed
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(rawFeeds, &feeds); nil != err {
			return err
		}
	case jsonparser.Object:
		var feed Feed
		if err := json.Unmarshal(rawFeeds, &feed); nil != err {
			return err
		}

		feeds = []Feed{feed}
	}

	*fl = FeedList{feeds}

	return nil
}

type Feed struct {
	Feed             string                `json:"Feed"`
	Status           string                `json:"Status"`
	Action           string                `json:"Action"`
	CreationDate     ScTimestamp           `json:"CreationDate"`
	UpdatedDate      ScTimestamp           `json:"UpdatedDate"`
	Source           string                `json:"Source"`
	TotalRecords     ScInt                 `json:"TotalRecords"`
	ProcessedRecords ScInt                 `json:"ProcessedRecords"`
	FailedRecords    ScInt                 `json:"FailedRecords"`
	FailureReports   FeedFailureReportFile `json:"FailureReports"`
}

type FeedFailureReportFile struct {
	MimeType string `json:"MimeType"`
	File     string `json:"File"`
}

func (ffrf *FeedFailureReportFile) UnmarshalJSON(b []byte) error {
	rawMimetype, dataType, _, err := jsonparser.Get(b, "MimeType")
	if err != nil {
		if err == jsonparser.KeyPathNotFoundError {
			return nil
		}
		return err
	}

	if len(rawMimetype) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot parse FeedFailureReportFile.Mimetype")
	}

	rawFile, dataType, _, err := jsonparser.Get(b, "File")
	if err != nil {
		return err
	}
	if len(rawFile) == 0 || dataType == jsonparser.NotExist {
		return errors.New("cannot parse FeedFailureReportFile.File")
	}

	fileString := string(rawFile)

	fileString = strings.Replace(fileString, "\\", "", -1)

	decodedFile, err := base64.StdEncoding.DecodeString(fileString)
	if err != nil {
		return err
	}

	*ffrf = FeedFailureReportFile{
		MimeType: string(rawMimetype),
		File:     string(decodedFile),
	}

	return nil
}

type FeedStatus struct {
	Feed             string       `json:"Feed"`
	Status           string       `json:"Status"`
	Action           string       `json:"Action"`
	CreationDate     ScTimestamp  `json:"CreationDate"`
	UpdatedDate      ScTimestamp  `json:"UpdatedDate"`
	Source           string       `json:"Source"`
	TotalRecords     ScInt        `json:"TotalRecords"`
	ProcessedRecords ScInt        `json:"ProcessedRecords"`
	FailedRecords    ScInt        `json:"FailedRecords"`
	FeedErrors       FeedErrors   `json:"FeedErrors"`
	FeedWarnings     FeedWarnings `json:"FeedWarnings"`
}

type FeedErrors struct {
	Errors []FeedError
}

func (fe *FeedErrors) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	raw, dataType, _, err := jsonparser.Get(b, "Error")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(raw) == 0 || dataType == jsonparser.NotExist {
		*fe = FeedErrors{[]FeedError{}}
		return nil
	}

	var errorItems []FeedError
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(raw, &errorItems); nil != err {
			return err
		}
	case jsonparser.Object:
		var errorItem FeedError
		if err := json.Unmarshal(raw, &errorItem); nil != err {
			return err
		}

		errorItems = []FeedError{errorItem}
	}

	*fe = FeedErrors{errorItems}

	return nil
}

type FeedError struct {
	Code        string `json:"Code"`
	Message     string `json:"Message"`
	SellerSku   string `json:"SellerSku"`
	OrderId     string `json:"OrderId"`
	OrderItemId string `json:"OrderItemId"`
}

type FeedWarnings struct {
	Warnings []FeedWarning
}

func (fw *FeedWarnings) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	raw, dataType, _, err := jsonparser.Get(b, "Warning")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return err
	}

	if len(raw) == 0 || dataType == jsonparser.NotExist {
		*fw = FeedWarnings{[]FeedWarning{}}
		return nil
	}

	var warningItems []FeedWarning
	switch dataType {
	case jsonparser.Array:
		if err := json.Unmarshal(raw, &warningItems); nil != err {
			return err
		}
	case jsonparser.Object:
		var warningItem FeedWarning
		if err := json.Unmarshal(raw, &warningItem); nil != err {
			return err
		}

		warningItems = []FeedWarning{warningItem}
	}

	*fw = FeedWarnings{warningItems}

	return nil
}

type FeedWarning struct {
	Message     string `json:"Message"`
	SellerSku   string `json:"SellerSku"`
	OrderId     string `json:"OrderId"`
	OrderItemId string `json:"OrderItemId"`
}
