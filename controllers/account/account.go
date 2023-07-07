package account

import (
	accountService "github.com/beerzezy/accounts-management-service/services/account"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Login(c echo.Context) error
	CreateAccount(c echo.Context) error
	UpdateAccount(c echo.Context) error
	GetAccountById(c echo.Context) error
}

type controller struct {
	accountService.AccountService
}

func NewController(accountService accountService.AccountService) Controller {
	return &controller{
		accountService,
	}
}
