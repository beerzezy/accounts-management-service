package routes

import (
	"github.com/beerzezy/accounts-management-service/appconfig"
	accountController "github.com/beerzezy/accounts-management-service/controllers/account"
	"github.com/beerzezy/accounts-management-service/repositories"
	accountService "github.com/beerzezy/accounts-management-service/services/account"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg appconfig.Config) {

	serverContext := cfg.Server.Context
	// loginPath := fmt.Sprintf("%s/login", serverContext)
	// registerPath := fmt.Sprintf("%s/registration", serverContext)

	// provide repositories
	repo := repositories.NewRepository(db)

	// provide controller
	accountController := accountController.NewController(accountService.NewAccountService(repo))

	v1 := e.Group(serverContext + "/api/v1")

	v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())
	// k.Use(echojwt.WithConfig(echojwt.Config{
	// 	Skipper: func(c echo.Context) bool {
	// 		isSkip := false
	// 		if c.Request().URL.Path == loginPath || c.Request().URL.Path == registerPath {
	// 			isSkip = true
	// 		}l............................................................................,
	// 		return isSkip
	// 	},
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(jwtCustomClaims)
	// 	},
	// 	SigningKey: []byte(cfg.JWT.SigningKey),
	// }))

	v1.POST("/login", accountController.Login)
	v1.POST("/create-account", accountController.CreateAccount)
	v1.GET("/account/:id", accountController.GetAccountById)
	v1.PUT("/update-account", accountController.UpdateAccount)
}
