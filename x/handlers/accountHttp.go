package handlers

import (
	"net/http"

	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/labstack/echo/v4"
)

type accountHttpHandler struct {
	accountUsecase usecases.AccountUsecase
}

// RegisterPersonAccount implements AccountHandler.
func (a *accountHttpHandler) SignupPerson(c echo.Context) error {
	var reqBody = new(models.AddPersonAccountData)

	if err := c.Bind(reqBody); err != nil {
		return response(c, http.StatusBadRequest, "error binding body")
	}

	if err := c.Validate(reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		// return response(c, http.StatusBadRequest, err.Error())
	}

	if err := a.accountUsecase.AddPersonAccount(reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		// return response(c, http.StatusBadRequest, err.Error())
	}

	return response(c, http.StatusCreated, "account successfully created")
}

func NewAccountHttpHandler(u usecases.AccountUsecase) AccountHandler {
	return &accountHttpHandler{
		accountUsecase: u,
	}
}
