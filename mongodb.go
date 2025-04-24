package oneGroup

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

func InitMongoDB() {
	mon := NaCos.Mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", mon.User, mon.Passwd, mon.Hort, mon.Port)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("mongodb初始化成功")
}
