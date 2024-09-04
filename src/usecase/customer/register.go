package customerusecase

import (
	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sCustomerUsecase) Register(p *entities.ParamsCreateCustomer) (*entities.ResultRegsiterCustomer, error) {
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

	return result, nil
}
