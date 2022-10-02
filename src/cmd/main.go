package main

import (
	"Improve/src/configs"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	fmt.Println("Huan")

	config := configs.App{}
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@(%v:%v)/%v", config.User, config.Pass, config.Host, config.Port, config.Database))
	if err != nil {
		log.Fatalf("Connect to database failed %v: ", err)
	} else {
		fmt.Println("Connect to database successfully")
	}

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pass,
	})
	err = rdb.Set(ctx, "key", "value", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val)
	defer func() {
		db.Close()
		rdb.Close()
	}()

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	r.Run(":8080")
}
