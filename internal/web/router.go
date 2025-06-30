package web

import (
	"log/slog"
	"net/http"

	"github.com/FedotCompot/file-cacher/internal/cache"
	"github.com/FedotCompot/file-cacher/internal/web/api"
	pages "github.com/FedotCompot/file-cacher/internal/web/api/pages"
	"github.com/FedotCompot/file-cacher/internal/web/middlewares"
	"github.com/FedotCompot/file-cacher/internal/web/types"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func getRouter() *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.WithMiddleware(reqlog.NewMiddleware(reqlog.WithVerbose(true))),
		bunrouter.WithMiddleware(errorHandler),
	)

	router.GET("/readyz", func(w http.ResponseWriter, r bunrouter.Request) error {
		return api.RenderStatus(w, http.StatusOK)
	})
	router.GET("/livez", func(w http.ResponseWriter, r bunrouter.Request) error {
		return api.RenderStatus(w, http.StatusOK)
	})

	// Internal API
	router.WithGroup("/api/v1", func(router *bunrouter.Group) {
		// JWT Authentication
		router = router.Use(middlewares.JWTAuthMiddleware())

		router.POST("/upload", pages.UploadPage)
	})

	// Serving static files from the cache
	router.GET("/*path", func(w http.ResponseWriter, r bunrouter.Request) error {
		path := r.URL.Path
		slog.Info("Fetching page from cache", "path", path)
		data, err := cache.GetPage(path)
		if err != nil {
			return api.RenderStatus(w, http.StatusNotFound)
		}

		return api.RenderContent(w, http.StatusOK, *data)
	})

	return router
}

func errorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)
		if err != nil {
			slog.ErrorContext(req.Context(),
				"error: "+err.Error(),
			)
		}
		switch err := err.(type) {
		case nil:
			// no error
		case types.Error: // already a types.Error
			w.WriteHeader(err.StatusCode)
			_ = bunrouter.JSON(w, err)
		default:
			httpErr := types.NewError(err)
			w.WriteHeader(httpErr.StatusCode)
			_ = bunrouter.JSON(w, httpErr)
		}
		return nil
	}
}
