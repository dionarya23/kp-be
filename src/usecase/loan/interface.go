package loanusecase

import (
	"github.com/dionarya23/kredit-plus/src/entities"
	customer "github.com/dionarya23/kredit-plus/src/repositories/customer"
	limit "github.com/dionarya23/kredit-plus/src/repositories/limit"
	loan "github.com/dionarya23/kredit-plus/src/repositories/loan"
)

type sLoanUsecase struct {
	loanRepository     loan.LoanRepository
	customerRepository customer.CustomerRepository
	limitRepository    limit.LimitRepository
}

type LoanUsecase interface {
	Create(p *entities.ParamsLoan) (*entities.ResultCreateLoan, error)
	FindOne(loanId int64, customerId int64) (*entities.Loan, error)
}

func New(
	loanRepository loan.LoanRepository,
	customerRepository customer.CustomerRepository,
	limitRepository limit.LimitRepository,
) LoanUsecase {
	return &sLoanUsecase{
		loanRepository:     loanRepository,
		customerRepository: customerRepository,
		limitRepository:    limitRepository,
	}
}
