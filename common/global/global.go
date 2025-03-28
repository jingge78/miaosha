package global

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"miaosha-jjl/common/config"
)

var (
	ConfigData   *config.AppViper
	DB           *gorm.DB
	Rdb          *redis.Client
	Es           *elasticsearch.Client
	CTX          = context.Background()
	ConsulClient *api.Client
	NaCos        *config.T
)
