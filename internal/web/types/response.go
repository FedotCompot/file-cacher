package types

import "time"

type PageResponse struct {
	Url string    `json:"url"`
	Exp time.Time `json:"exp"`
}
