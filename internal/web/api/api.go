package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/FedotCompot/file-cacher/internal/web/types"
)

func RenderJSON(w http.ResponseWriter, status int, v interface{}) error {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(buf.Bytes())
	return err
}

func RenderContent(w http.ResponseWriter, status int, content types.Data) error {
	w.Header().Set("Content-Type", content.ContentType)
	w.WriteHeader(status)
	data, err := base64.StdEncoding.DecodeString(content.Data)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func RenderStatus(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}
