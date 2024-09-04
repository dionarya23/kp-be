package customerv1controller

import (
	"net/http"

	customerUsecase "github.com/dionarya23/kredit-plus/src/usecase/customer"

	customerRepository "github.com/dionarya23/kredit-plus/src/repositories/customer"
	"github.com/labstack/echo/v4"
)

func (i *V1Customer) Login(c echo.Context) (err error) {
	u := new(loginRequest)

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

	uu := customerUsecase.New(
		customerRepository.New(i.DB),
	)

	data, err := uu.Login(&customerUsecase.ParamsLogin{
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
	})

	if err != nil {
		if err == customerUsecase.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}

type (
	loginRequest struct {
		Password    string `json:"password" validate:"required"`
		PhoneNumber string `json:"phoneNumber" validate:"required"`
	}
)
