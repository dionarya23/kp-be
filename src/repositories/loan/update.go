package loanrepository

import (
	"log"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sLoanRepository) Update(loanId *int64, p *entities.ParamsPaymentLoan) error {
	if p.Status != "" {
		query := "UPDATE Loans SET status = ? WHERE id = ?"
		_, err := i.DB.Exec(query, p.Status, loanId)

		if err != nil {
			log.Printf("Error getting last insert ID: %s", err)
			return err
		}
	} else {
		query := "UPDATE Loans SET remaining_balance = remaining_balance - ? WHERE id = ?"
		_, err := i.DB.Exec(query, p.Amount, loanId)

		if err != nil {
			log.Printf("Error getting last insert ID: %s", err)
			return err
		}
	}

	return nil
}
