package loanusecase

import (
	"github.com/dionarya23/kredit-plus/src/entities"
	customerusecase "github.com/dionarya23/kredit-plus/src/usecase/customer"
)

func (i *sLoanUsecase) Payment(p *entities.ParamsPaymentLoan) (int, error) {
	customerExist, err := i.customerRepository.FindOne(&entities.ParamsCustomer{
		ID: int64(p.CustomerId),
	})

	if err != nil {
		return 0, err
	}

	if customerExist == nil {
		return 0, customerusecase.ErrUserNotFound
	}

	loan, err := i.loanRepository.FindOne(int64(p.LoanId), int64(p.CustomerId))
	if err != nil {
		return 0, err
	}

	if loan == nil {
		return 0, err
	}

	newBalance := int64(loan.RemainingBalance) - int64(p.Amount)

	err = i.loanRepository.Update(&loan.ID, &entities.ParamsPaymentLoan{
		Amount: p.Amount,
	})

	if err != nil {
		return 0, err
	}

	loanId := int64(p.LoanId)

	newLoan, err := i.loanRepository.FindOne(loanId, loan.CustomerID)
	if err != nil {
		return 0, err
	}

	if newLoan.RemainingBalance <= 0 {
		err = i.loanRepository.Update(&loan.ID, &entities.ParamsPaymentLoan{
			Status: "completed",
		})

		if err != nil {
			return 0, err
		}
	}

	return int(newBalance), nil
}
