package account

import (
	"net/http"

	"github.com/beerzezy/accounts-management-service/baseresponse"
	"github.com/beerzezy/accounts-management-service/errors"
	"github.com/beerzezy/accounts-management-service/services/account"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (ctrl *controller) CreateAccount(c echo.Context) error {

	var request account.RequestCreateAccount
	if err := c.Bind(&request); err != nil {
		log.Errorf("CreateAccount bind baseRequest: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewCannotBindRequestStructError())
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("CreateAccount Validate Request: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewInvalidRequireFieldError())
	}

	result, err := ctrl.AccountService.CreateAccount(request)
	if err != nil {
		log.Errorf("CreateAccount Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
		Data:    result,
	})
}
