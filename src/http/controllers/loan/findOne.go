package loanv1controller

import (
	"net/http"
	"strconv"

	"github.com/dionarya23/kredit-plus/src/entities"
	customerrepository "github.com/dionarya23/kredit-plus/src/repositories/customer"
	limitrepository "github.com/dionarya23/kredit-plus/src/repositories/limit"
	loanrepository "github.com/dionarya23/kredit-plus/src/repositories/loan"
	loanUsecase "github.com/dionarya23/kredit-plus/src/usecase/loan"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func (i *V1Loan) FindOne(c echo.Context) (err error) {
	filters := &entities.ParamsLoan{}

	idUser := new(entities.MeValidator)
	mapstructure.Decode(c.Get("user"), &idUser)

	if id := c.QueryParam("id"); id != "" {
		result, _ := strconv.ParseInt(id, 10, 64)
		filters.ID = result
	}
	filters.CustomerID = int64(idUser.ID)

	cu := loanUsecase.New(
		loanrepository.New(i.DB),
		customerrepository.New(i.DB),
		limitrepository.New(i.DB),
	)

	data, err := cu.FindOne(filters.ID, filters.CustomerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Loan found successfully",
		Data:    data,
	})
}
