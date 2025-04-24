package oneGroup

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Es  *elasticsearch.Client
	DB  *gorm.DB
	Rdb *redis.Client
)
