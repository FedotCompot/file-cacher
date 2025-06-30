package cache

import (
	"context"
	"encoding/json"

	"github.com/FedotCompot/file-cacher/internal/config"
	"github.com/FedotCompot/file-cacher/internal/web/types"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func Initialize(ctx context.Context) error {
	options, err := redis.ParseURL(config.Data.RedisUrl)
	if err != nil {
		return err
	}
	rdb = redis.NewClient(options)
	_, err = rdb.Ping(ctx).Result()
	return err
}

func Close() error {
	if rdb == nil {
		return nil
	}
	return rdb.Close()
}

func GetPage(path string) (*types.Data, error) {
	data, err := rdb.Get(ctx, path).Result()
	if err != nil {
		return nil, err // Other error
	}
	var page types.Data
	err = json.Unmarshal([]byte(data), &page)
	return &page, err
}

func UploadPage(page *types.UploadRequest) error {
	ttl := config.Data.CacheTTL
	if page.TTLOverride != nil {
		ttl = *page.TTLOverride
	}
	data, err := json.Marshal(page.Data)
	if err != nil {
		return err
	}
	_, err = rdb.Set(ctx, page.Path, string(data), ttl).Result()
	return err
}
