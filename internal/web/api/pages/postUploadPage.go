package typicals_api

import (
	"net/http"
	"time"

	"github.com/FedotCompot/file-cacher/internal/cache"
	"github.com/FedotCompot/file-cacher/internal/config"
	"github.com/FedotCompot/file-cacher/internal/web/api"
	"github.com/FedotCompot/file-cacher/internal/web/types"
	"github.com/uptrace/bunrouter"
)

func UploadPage(w http.ResponseWriter, r bunrouter.Request) error {
	page, err := api.ParseRequestNoValidate[types.UploadRequest](r)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err := cache.UploadPage(page); err != nil {
		return err
	}
	return api.RenderJSON(w, http.StatusOK, types.PageResponse{
		Url: config.Data.Host + page.Path,
		Exp: time.Now().Add(config.Data.CacheTTL),
	})
}
