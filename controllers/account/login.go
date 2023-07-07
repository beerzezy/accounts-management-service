package account

import (
	"net/http"

	"github.com/beerzezy/accounts-management-service/baseresponse"
	"github.com/beerzezy/accounts-management-service/errors"
	"github.com/beerzezy/accounts-management-service/services/account"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (ctrl *controller) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var request account.RequestLoginAccount
	if err := c.Bind(&request); err != nil {
		log.Errorf("CreateAccount bind baseRequest: %v", err)
		return err
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("CreateAccount Validate Request: %v", err)
		return errors.NewInvalidRequireFieldError()
	}

	result, err := ctrl.AccountService.Login(ctx, request)
	if err != nil {
		log.Errorf("CreateAccount Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
		Data:    result,
	})
}
