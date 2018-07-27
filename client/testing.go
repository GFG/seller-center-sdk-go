package client

import (
	"log"
)

type FakeClient struct {
	FakeResponse Response
	FakeError    error
}

func (c FakeClient) GetLogger() *log.Logger {
	return nil
}

func (c FakeClient) Call(request Request) (Response, error) {
	return c.FakeResponse, c.FakeError
}
