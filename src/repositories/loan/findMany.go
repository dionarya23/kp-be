package loanrepository

import (
	"log"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sLoanRepository) FindMany(customerId int64) ([]*entities.Loan, error) {
	query := "SELECT id, customer_id, amount, tenor, purpose, interest_rate, monthly_installment, status, created_at, updated_at, remaining_balance FROM Loans WHERE customer_id = ?"

	rows, err := i.DB.Query(query, customerId)
	if err != nil {
		log.Printf("Error querying limits: %s", err)
		return nil, err
	}
	defer rows.Close()

	loans := make([]*entities.Loan, 0)

	for rows.Next() {
		loan := new(entities.Loan)
		err := rows.Scan(&loan.ID, &loan.CustomerID, &loan.Amount, &loan.Tenor, &loan.Purpose, &loan.InterestRate, &loan.MonthlyInstallment, &loan.Status, &loan.CreatedAt, &loan.UpdatedAt, &loan.RemainingBalance)
		if err != nil {
			log.Printf("Error scanning row: %s", err)
			return nil, err
		}
		loans = append(loans, loan)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %s", err)
		return nil, err
	}

	return loans, nil
}
