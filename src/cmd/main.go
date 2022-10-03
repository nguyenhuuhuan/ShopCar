package main

import (
	"Improve/src/configs"
	"Improve/src/controllers"
	"Improve/src/repositories"
	"Improve/src/services"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	config := &configs.App{}
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Database))
	if err != nil {
		log.Fatalf("Connect to database failed %v: ", err)
	} else {
		fmt.Println("Connect database successfully")
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DefaultStringSize:         256,  // default size for string fields
		DisableDatetimePrecision:  true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect to gorm failed %v: ", err)
	}
	ctx := context.Background()

	// connect Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pass,
	})
	err = rdb.Set(ctx, "key", "connect redis successfully", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val)

	defer func() {
		db.Close()
		rdb.Close()
	}()

	// declare Repositories
	var (
		userRepo = repositories.NewUserRepository(gormDB)
	)

	// declare Services
	var (
		authService = services.NewAuthService(userRepo)
	)

	// declare Controllers
	var (
		authController = controllers.NewAuthController(authService)
	)

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	authRouter := r.Group("api/auth")
	{
		authRouter.POST("/register", authController.Register)
		authRouter.GET("/login", authController.Login)
	}

	r.Run(":8080")
}
