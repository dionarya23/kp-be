package loanrepository

import (
	"bytes"
	"database/sql"
	"errors"
	"log"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sLoanRepository) FindOne(loanId int64, customerId int64) (*entities.Loan, error) {
	var query bytes.Buffer
	query.WriteString("SELECT id, customer_id, amount, tenor, purpose, interest_rate, monthly_installment, status, created_at, updated_at, remaining_balance FROM Loans WHERE id = ? AND customer_id = ?")

	row := i.DB.QueryRow(query.String(), loanId, customerId)
	var loan entities.Loan

	err := row.Scan(
		&loan.ID,
		&loan.CustomerID,
		&loan.Amount,
		&loan.Tenor,
		&loan.Purpose,
		&loan.InterestRate,
		&loan.MonthlyInstallment,
		&loan.Status,
		&loan.CreatedAt,
		&loan.UpdatedAt,
		&loan.RemainingBalance,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Error finding loan")
		}
		log.Printf("Error finding loan: %s", err)
		return nil, err
	}

	return &loan, nil
}
