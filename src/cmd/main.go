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

	config, _ := configs.New()
	formatDsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", formatDsn)
	if err != nil {
		log.Fatalf("Connect to database failed %v: ", err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
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
		userRepo     = repositories.NewUserRepository(gormDB)
		roleRepo     = repositories.NewRoleRepository(gormDB)
		userRoleRepo = repositories.NewUserRoleRepository(gormDB)
	)

	// declare Services
	var (
		authService = services.NewAuthService(userRepo, roleRepo, userRoleRepo)
		roleService = services.NewRoleService(roleRepo)
	)

	// declare Controllers
	var (
		authController = controllers.NewAuthController(authService)
		roleController = controllers.NewRoleController(roleService)
	)

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	v1 := r.Group("shop-car")
	{
		authRouter := v1.Group("/api/auth")
		{
			authRouter.POST("/register", authController.Register)
			authRouter.POST("/login", authController.Login)
		}
		role := v1.Group("/")
		{
			role.POST("role", roleController.Create)
		}

	}

	r.Run(":8080")
}
