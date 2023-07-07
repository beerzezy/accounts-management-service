package account

import (
	"net/http"
	"strconv"

	"github.com/beerzezy/accounts-management-service/baseresponse"
	"github.com/beerzezy/accounts-management-service/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (ctrl *controller) GetAccountById(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("GetAccountById ParseUint Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewInvalidRequestWithMessageError(err.Error()))
	}

	if id == 0 {
		log.Errorf("GetAccountById Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewInvalidRequestWithMessageError("Invalid Require ID"))
	}

	result, err := ctrl.AccountService.GetAccountById(uint64(id))
	if err != nil {
		log.Errorf("GetAccountById Error: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewExceptionError(err.Error()))
	}

	return c.JSON(http.StatusOK, baseresponse.BaseResponse{
		Message: "success",
		Data:    result,
	})
}
