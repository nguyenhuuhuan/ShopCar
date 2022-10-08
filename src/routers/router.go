package routers

import (
	"Improve/src/configs"
	"Improve/src/controllers"
	"Improve/src/logger"
	"Improve/src/repositories"
	"Improve/src/services"
	"Improve/src/token"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title ShopCar Web
// @version 1.0
// @description This is a server of ShopCar Web.
// @termsOfService http://swagger.io/terms/

// @contact.name Huan Nguyen
// @contact.url http://www.swagger.io/support
// @contact.email huuhuan19031999@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /shop-car

// @securityDefinitions.apikey BasicAuthToken
// @in header
// @name Authorization

// @securityDefinitions.apikey JWTAccessToken
// @in header
// @name Authorization
func InitRouter(ctx context.Context, app *configs.App, db *gorm.DB) (*gin.Engine, error) {
	// declare Repositories
	var (
		userRepo     = repositories.NewUserRepository(db)
		roleRepo     = repositories.NewRoleRepository(db)
		userRoleRepo = repositories.NewUserRoleRepository(db)
	)

	maker, err := token.NewJWTMaker(app.JWT.SecretKey)
	if err != nil {
		logger.Context(ctx).Errorf("Cannot create token %v", err)
		return nil, err
	}
	// declare Services
	var (
		authService = services.NewAuthService(*app, maker, userRepo, roleRepo, userRoleRepo)
		roleService = services.NewRoleService(roleRepo)
		userService = services.NewUserService(userRepo)
	)

	// declare Controllers
	var (
		authController = controllers.NewAuthController(authService)
		roleController = controllers.NewRoleController(roleService)
		userController = controllers.NewUserController(userService)
	)

	_, err = newrelic.NewApplication(
		newrelic.ConfigAppName("ShopCar"),
		newrelic.ConfigLicense("037688c44b1de03a303fdd612b5aed3f9004NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		logger.Fatalf(err, "cannot start newrelic application, err : %v", err)
	}

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	v1 := r.Group("shop-car")
	{
		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/register", authController.Register)
			authRouter.POST("/login", authController.Login)
		}
		role := v1.Group("/role")
		{
			role.POST("", roleController.Create)
		}
		user := v1.Group("/user")
		{
			user.GET("/", userController.List)
			user.GET("/:id", userController.GetUser)
		}

	}
	if configs.AppConfig.RunMode == gin.DebugMode && configs.AppConfig.Env != "PRODUCTION" {
		r.GET("/gamification/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r, nil
}
