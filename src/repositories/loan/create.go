package loanrepository

import (
	"log"
	"time"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sLoanRepository) Create(p *entities.ParamsLoan) (*entities.Loan, error) {
	query := "INSERT INTO Loans (customer_id, amount, tenor, purpose, interest_rate, monthly_installment, status, remaining_balance) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := i.DB.Exec(query, p.CustomerID, p.Amount, p.Tenor, p.Purpose, p.InterestRate, p.MonthlyInstallment, p.Status, p.Amount)
	if err != nil {
		log.Printf("Error inserting loan: %s", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %s", err)
		return nil, err
	}

	loan := &entities.Loan{
		ID:                 id,
		CustomerID:         p.CustomerID,
		Amount:             p.Amount,
		Tenor:              p.Tenor,
		Purpose:            p.Purpose,
		InterestRate:       p.InterestRate,
		MonthlyInstallment: p.MonthlyInstallment,
		RemainingBalance:   p.Amount,
		Status:             p.Status,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	return loan, nil
}
