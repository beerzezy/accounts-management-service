package account

import (
	"net/http"
	"strconv"

	"github.com/beerzezy/accounts-management-service/baseresponse"
	"github.com/beerzezy/accounts-management-service/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (ctrl *controller) GetAccounts(c echo.Context) error {
	return nil
}

func (ctrl *controller) GetAccountById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("GetAccountById ParseUint Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewInvalidRequestWithMessageError(err.Error()))
	}

	result, err := ctrl.AccountService.GetAccountById(ctx, uint64(id))
	if err != nil {
		log.Errorf("GetAccountById Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
		Data:    result,
	})
}
