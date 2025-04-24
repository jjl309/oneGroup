package oneGroup

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	data := NaCos.Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     data.Addr,
		Password: data.Password, // no password set
		DB:       data.Db,       // use default DB
	})

	pong, err := Rdb.Ping(Rdb.Context()).Result()
	fmt.Println(pong, err)
	if err != nil {
		panic("redis failed to connect")
	}
	fmt.Println("redis connect success")
	// Output: PONG <nil>
}
