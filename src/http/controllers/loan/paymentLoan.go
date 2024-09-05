package loanv1controller

import (
	"net/http"

	"github.com/dionarya23/kredit-plus/src/entities"
	customerrepository "github.com/dionarya23/kredit-plus/src/repositories/customer"
	limitrepository "github.com/dionarya23/kredit-plus/src/repositories/limit"
	loanrepository "github.com/dionarya23/kredit-plus/src/repositories/loan"
	loanUsecase "github.com/dionarya23/kredit-plus/src/usecase/loan"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func (i *V1Loan) PaymentLoan(c echo.Context) (err error) {
	u := new(entities.ParamsCreatePaymentLoan)
	idUser := new(entities.MeValidator)
	mapstructure.Decode(c.Get("user"), &idUser)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	cu := loanUsecase.New(
		loanrepository.New(i.DB),
		customerrepository.New(i.DB),
		limitrepository.New(i.DB),
	)

	data, err := cu.Payment(&entities.ParamsPaymentLoan{
		Amount:     u.Amount,
		LoanId:     u.LoanId,
		CustomerId: float64(idUser.ID),
	})

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Payment successful",
		Data:    data,
	})
}
