package api

import (
	"encoding/json"
	"io"

	"github.com/uptrace/bunrouter"
)

type InputModel interface {
	Validate(r bunrouter.Request) error
}

func ParseRequest[model InputModel](r bunrouter.Request) (*model, error) {
	// Read raw JSON body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// Unmarshal JSON body
	var req model
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, err
	}
	// Validate input
	if err := req.Validate(r); err != nil {
		return nil, err
	}
	return &req, nil
}

func ParseRequestNoValidate[model any](r bunrouter.Request) (*model, error) {
	// Read raw JSON body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// Unmarshal JSON body
	var req model
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func ParseAsString(r bunrouter.Request) (string, error) {
	// Read raw JSON body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	// Unmarshal JSON body
	return string(body), nil
}
