package global

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"miaosha-jjl/common/config"
)

var (
	DB           *gorm.DB
	Rdb          *redis.Client
	GlobalConfig *config.Config
	Es           *elasticsearch.Client
	CTX          = context.Background()
)
