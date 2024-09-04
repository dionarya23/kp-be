package customerusecase

import (
	"os"
	"strconv"

	"github.com/dionarya23/kredit-plus/src/entities"
	"github.com/dionarya23/kredit-plus/src/helpers"
)

func (i *sCustomerUsecase) Register(p *entities.ParamsCreateCustomer) (*entities.ResultAuth, error) {
	filters := entities.ParamsCustomer{
		NIK: p.NIK,
	}

	checkNik, _ := i.customerRepository.IsExists(&filters)

	if checkNik {
		return nil, ErrNikAlreadyUsed
	}

	customer, err := i.customerRepository.Create(p)
	if err != nil {
		return nil, err
	}

	result := &entities.ResultRegsiterCustomer{
		ID: customer.ID,
	}

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		UserId:          result.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if err != nil {
		return nil, err
	}

	return &entities.ResultAuth{
		ID:          strconv.FormatInt(customer.ID, 10),
		AccessToken: accessToken,
	}, nil
}
