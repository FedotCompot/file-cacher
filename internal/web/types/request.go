package types

import "time"

type UploadRequest struct {
	Path        string         `json:"path"`
	Data        Data           `json:"data"`
	TTLOverride *time.Duration `json:"ttl,omitempty"`
}

type Data struct {
	Data        string `json:"data"`
	ContentType string `json:"content_type,omitempty"`
}
