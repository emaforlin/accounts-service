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
	reqBody := &models.AddPersonAccountData{}

	if err := c.Bind(reqBody); err != nil {
		return response(c, http.StatusBadRequest, "error binding body")
	}

	if err := a.accountUsecase.AddPersonAccount(reqBody); err != nil {
		return response(c, http.StatusBadRequest, err.Error())
	}

	return nil
}

func NewAccountHttpHandler(u usecases.AccountUsecase) AccountHandler {
	return &accountHttpHandler{
		accountUsecase: u,
	}
}
