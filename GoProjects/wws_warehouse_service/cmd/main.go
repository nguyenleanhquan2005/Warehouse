//go:build debug
// +build debug

package main

import (
	"app/internal/config"
	"app/internal/domain/service/rest_service/auth_service"
	"app/internal/domain/service/rest_service/user_service"
	"app/internal/domain/usecase"
	"app/internal/framework/route"
	"app/internal/infrastructure/db/mysql"
	"app/internal/infrastructure/handler"
	gormmysql "app/pkg/gorm_mysql"
	"fmt"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := gormmysql.InitDatabase(gormmysql.Config{
		User:         cfg.Mysql.DBUser,
		Password:     cfg.Mysql.DBPass,
		Host:         cfg.Mysql.DBHost,
		Port:         cfg.Mysql.DBPort,
		DatabaseName: cfg.Mysql.DBName,
	})
	if err != nil {
		panic(err)
	}

	// init auth service
	auth_service.NewAuthService(
		cfg.AuthService.URL,
		cfg.AuthService.APITimeOut,
	)

	// init user service
	user_service.NewUserService(
		cfg.UserService.URL,
		cfg.UserService.APITimeOut,
	)

	// init repository
	balanceRepo := mysql.NewBalanceRepository(db)

	// init usecase
	balanceListByUserId := usecase.NewBalanceListByUserIdUsecase(balanceRepo)
	balanceGetUsecase := usecase.NewBalanceGetUsecase(balanceRepo)
	// init handler
	balanceHandler := handler.NewBalanceHandler(balanceListByUserId, balanceGetUsecase)

	r := route.NewRouter(
		cfg,
		balanceHandler,
	)

	if err := r.Run(":" + cfg.PORT); err != nil {
		panic(fmt.Errorf("failed to run server: %w", err))
	}
}

//nolint:unused
func mainDebug() {

	//cfg, err := config.GetConfig()
	//if err != nil {
	//	panic(err)
	//}

	//db, err := gormmysql.InitDatabase(gormmysql.Config{
	//	User:         cfg.Mysql.DBUser,
	//	Password:     cfg.Mysql.DBPass,
	//	Host:         cfg.Mysql.DBHost,
	//	Port:         cfg.Mysql.DBPort,
	//	DatabaseName: cfg.Mysql.DBName,
	//})
	//if err != nil {
	//	panic(err)
	//}

	//ctx := context.Background()
}
