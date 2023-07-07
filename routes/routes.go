package routes

import (
	"fmt"

	"github.com/beerzezy/accounts-management-service/appconfig"
	accountController "github.com/beerzezy/accounts-management-service/controllers/account"
	accountMiddleware "github.com/beerzezy/accounts-management-service/middleware"
	"github.com/beerzezy/accounts-management-service/repositories"
	accountService "github.com/beerzezy/accounts-management-service/services/account"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg appconfig.Config) {

	// provide repositories
	repo := repositories.NewRepository(db)

	// provide controller
	accountController := accountController.NewController(accountService.NewAccountService(cfg, repo))

	serverContext := cfg.Server.Context
	v1 := e.Group(serverContext + "/api/v1")
	v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	// jwt middleware
	loginPath := fmt.Sprintf("%s/api/v1/login", serverContext)
	createAccountPath := fmt.Sprintf("%s/api/v1/create-account", serverContext)
	v1.Use(accountMiddleware.NewJWTMiddleware(
		cfg.JWT.SigningKey,
		[]string{
			loginPath,
			createAccountPath,
		},
	).Handler())

	v1.POST("/login", accountController.Login)
	v1.POST("/create-account", accountController.CreateAccount)
	v1.GET("/account/:id", accountController.GetAccountById)
	v1.PUT("/update-account", accountController.UpdateAccount)
}
