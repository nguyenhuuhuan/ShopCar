package main

import (
	"Improve/src/configs"
	"Improve/src/logger"
	"Improve/src/routers"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	config, _ := configs.New()
	formatDsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", config.MYSQLUser, config.MYSQLPass, config.MYSQLHost, config.MYSQLPort, config.MYSQLDatabase)
	db, err := sql.Open("mysql", formatDsn)
	fmt.Println(formatDsn)
	if err != nil {
		logger.Fatalf(err, "Connect to database failed %v: ", err)
	}
	gormDB, err := gorm.Open(mysql.Open(formatDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect to gorm failed %v: ", err)
	}

	ctx := context.Background()

	// connect Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.URL(),
		Password: config.Redis.Password,
	})
	err = rdb.Set(ctx, "key", "connect redis successfully", 10*time.Second).Err()
	if err != nil {
		logger.Context(ctx).Fatalf(err, "[Redis] Dial connection %v", err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val)

	defer func() {
		db.Close()
		rdb.Close()
	}()
	fmt.Println("Router")
	// Init Router
	r, err := routers.InitRouter(ctx, config, gormDB)

	if err := r.Run(configs.AppConfig.AddressListener()); err != nil {
		logger.Fatalf(err, "Opening HTTP server: %v", err)
	}
}
