package customerusecase

import (
	"os"
	"strconv"

	"github.com/dionarya23/kredit-plus/src/entities"
	"github.com/dionarya23/kredit-plus/src/helpers"
)

type (
	ParamsLogin struct {
		PhoneNumber string
		Password    string
	}
	GeneratedToken struct {
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expired_at"`
	}
	ResultLogin struct {
		ID          string `json:"userId"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		AccessToken string `json:"accessToken"`
	}
)

func (i *sCustomerUsecase) Login(p *ParamsLogin) (*ResultLogin, error) {
	filters := entities.ParamsCustomer{
		PhoneNumber: p.PhoneNumber,
	}

	user, _ := i.customerRepository.FindOne(&filters)

	if user == nil {
		return nil, ErrUserNotFound
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		UserId:          user.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, user.Password)
	if !isValidPassword {
		return nil, ErrInvalidUser
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTLogin)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	return &ResultLogin{
		ID:          strconv.FormatInt(user.ID, 10),
		PhoneNumber: p.PhoneNumber,
		AccessToken: accessToken,
	}, nil
}
