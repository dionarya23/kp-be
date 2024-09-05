package loanusecase

import (
	"github.com/dionarya23/kredit-plus/src/entities"
	customerusecase "github.com/dionarya23/kredit-plus/src/usecase/customer"
)

func (i *sLoanUsecase) FindMany(customerId int64) ([]*entities.Loan, error) {
	customerExist, err := i.customerRepository.FindOne(&entities.ParamsCustomer{
		ID: customerId,
	})

	if err != nil {
		return nil, err
	}

	if customerExist == nil {
		return nil, customerusecase.ErrUserNotFound
	}

	loan, err := i.loanRepository.FindMany(customerId)
	if err != nil {
		return nil, err
	}

	return loan, nil
}
