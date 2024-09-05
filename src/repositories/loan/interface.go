package loanrepository

import (
	"database/sql"

	"github.com/dionarya23/kredit-plus/src/entities"
)

type sLoanRepository struct {
	DB *sql.DB
}

type LoanRepository interface {
	Create(*entities.ParamsLoan) (*entities.Loan, error)
}

func New(db *sql.DB) LoanRepository {
	return &sLoanRepository{DB: db}
}
