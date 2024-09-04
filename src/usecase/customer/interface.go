package customerusecase

import (
	"github.com/dionarya23/kredit-plus/src/entities"
	customer "github.com/dionarya23/kredit-plus/src/repositories/customer"
)

type sCustomerUsecase struct {
	customerRepository customer.CustomerRepository
}

type CustomerUsecase interface {
	Register(*entities.ParamsCreateCustomer) (*entities.ResultAuth, error)
	Login(p *ParamsLogin) (*ResultLogin, error)
}

func New(
	customerRepository customer.CustomerRepository,
) CustomerUsecase {
	return &sCustomerUsecase{
		customerRepository: customerRepository,
	}
}
