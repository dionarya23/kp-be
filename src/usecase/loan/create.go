package loanusecase

import (
	"github.com/dionarya23/kredit-plus/src/entities"
	customerusecase "github.com/dionarya23/kredit-plus/src/usecase/customer"
)

func (i *sLoanUsecase) Create(p *entities.ParamsLoan) (*entities.ResultCreateLoan, error) {
	customer, err := i.customerRepository.FindOne(&entities.ParamsCustomer{
		ID: p.CustomerID,
	})

	if customer == nil {
		return nil, customerusecase.ErrUserNotFound
	}

	if err != nil {
		return nil, customerusecase.ErrApps
	}

	customerLimit, err := i.limitRepository.FindMany(p.CustomerID, p.Tenor)
	if err != nil {
		return nil, customerusecase.ErrApps
	}

	if len(customerLimit) == 0 || p.Amount > customerLimit[0].LimitAmount {
		return nil, customerusecase.ErrLimit
	}

	var interest = calculateInterest(p.Amount, p.Tenor)
	var installment = calculateInstallment(p.Amount, interest, p.Tenor)
	newLoan, err := i.loanRepository.Create(&entities.ParamsLoan{
		CustomerID:         p.CustomerID,
		Amount:             p.Amount,
		Tenor:              p.Tenor,
		Purpose:            p.Purpose,
		InterestRate:       interest,
		MonthlyInstallment: installment,
		Status:             "pending",
	})

	if err != nil {
		return nil, customerusecase.ErrApps
	}

	return &entities.ResultCreateLoan{
		ID:                 newLoan.ID,
		Amount:             newLoan.Amount,
		Tenor:              newLoan.Tenor,
		Interest:           newLoan.InterestRate,
		MonthlyInstallment: newLoan.MonthlyInstallment,
	}, nil
}

func calculateInterest(amount float64, tenor int) float64 {
	var interestRate = 0.05
	var tenorCalculate = float64(tenor / 12)
	var result = amount * interestRate * tenorCalculate

	return result
}

func calculateInstallment(amount float64, interest float64, tenor int) float64 {
	var total = amount + interest
	return total / float64(tenor)
}
