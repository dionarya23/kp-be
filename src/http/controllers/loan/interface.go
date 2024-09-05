package loanv1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Loan struct {
	DB *sql.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Loan interface {
	Create(c echo.Context) error
}

func New(v1Loan *V1Loan) iV1Loan {
	return v1Loan
}
