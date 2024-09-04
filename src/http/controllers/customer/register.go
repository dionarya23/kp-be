package customerv1controller

import (
	"net/http"

	"github.com/dionarya23/kredit-plus/src/entities"
	customerrepository "github.com/dionarya23/kredit-plus/src/repositories/customer"
	customerUsecase "github.com/dionarya23/kredit-plus/src/usecase/customer"
	"github.com/labstack/echo/v4"
)

func (i *V1Customer) Register(c echo.Context) (err error) {
	u := new(entities.ParamsCreateCustomer)

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

	cu := customerUsecase.New(
		customerrepository.New(i.DB),
	)

	data, err := cu.Register(u)
	if err != nil {
		return c.JSON(http.StatusConflict, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Customer registered successfully",
		Data:    data,
	})
}
