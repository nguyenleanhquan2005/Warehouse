package route

import (
	"app/internal/config"
	"app/internal/infrastructure/handler"
	middlewarepkg "app/pkg/middleware"

	_ "app/internal/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// NewRouter define mapping routes
// @title product service backend
// @version 1.0
// @description This is the project of product service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @SecurityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(
	cfg *config.Config,
	balanceHandler *handler.BalanceHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: middlewarepkg.LogFormatterJSON,
			Output:    gin.DefaultWriter,
			SkipPaths: []string{"/", "/api/healthz"},
		}),

		middlewarepkg.Recovery(),
		middlewarepkg.Secure(),
		middlewarepkg.Headers,
	)

	if cfg.ENV == config.ENVProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	if cfg.ENV == config.ENVDevelopment || cfg.ENV == config.ENVProduction {
		router.Use(middlewarepkg.CorsMiddleware(cfg.CORS.AllowHosts))
	}

	router.GET("/api/healthz", middlewarepkg.Health)

	router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// validators
	//roleValidator := middleware.NewRolesValidator(middleware.RoleAdmin)

	// middlewares
	//jwt := middlewarepkg.NewJWT(cfg.JWTSecret)
	//auth := middleware.NewAuth(
	//	auth_service.NewAuthService(
	//		cfg.AuthService.URL,
	//		cfg.AuthService.APITimeOut,
	//	),
	//	user_service.NewUserService(cfg.UserService.URL, cfg.UserService.APITimeOut),
	//	jwt,
	//	roleValidator,
	//)

	v1Auth := router.Group("/api/v1")

	//
	//v1Auth := router.Group("/api/v1")
	//
	//adminGroup := v1Auth.Group("/admin")

	//Todo : using auth.Authenticate to verify before using balance-service
	warehouseGroup := v1Auth.Group("/warehouses")
	{
		warehouseGroup.GET("", warehouseHandler.ListBalanceByUserID)
		warehouseGroup.GET("/:user_id/:currency", balanceHandler.GetBalance)
		warehouseGroup.PUT("", balanceHandler.UpdateBalanceByUserID)
	}

	return router
}
