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

func (i *V1Loan) Create(c echo.Context) (err error) {
	u := new(entities.ParamsCreateLoan)
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

	data, err := cu.Create(&entities.ParamsLoan{
		CustomerID: int64(idUser.ID),
		Amount:     u.Amount,
		Tenor:      u.Tenor,
		Purpose:    u.Purpose,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Loan application submitted successfully",
		Data:    data,
	})
}
