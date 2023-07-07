package account

import (
	"net/http"

	"github.com/beerzezy/accounts-management-service/baseresponse"
	"github.com/beerzezy/accounts-management-service/errors"
	"github.com/beerzezy/accounts-management-service/services/account"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (ctrl *controller) UpdateAccount(c echo.Context) error {

	var request account.RequestUpdateAccount
	if err := c.Bind(&request); err != nil {
		log.Errorf("UpdateAccount bind baseRequest: %v", err)
		return err
	}

	if err := c.Validate(request); err != nil {
		log.Errorf("UpdateAccount Validate Request: %v", err)
		return errors.NewInvalidRequireFieldError()
	}

	err := ctrl.AccountService.UpdateAccount(request)
	if err != nil {
		log.Errorf("UpdateAccount Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
	})
}
