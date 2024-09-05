package limitrepository

import (
	"log"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sLimitRepository) FindMany(customerId int64, tenor int) ([]*entities.Limit, error) {
	query := "SELECT id, customer_id, tenor, limit_amount, created_at, updated_at FROM Limits WHERE customer_id = ? AND tenor = ?"

	rows, err := i.DB.Query(query, customerId, tenor)
	if err != nil {
		log.Printf("Error querying limits: %s", err)
		return nil, err
	}
	defer rows.Close()

	limits := make([]*entities.Limit, 0)

	for rows.Next() {
		limit := new(entities.Limit)
		err := rows.Scan(&limit.ID, &limit.CustomerID, &limit.Tenor, &limit.LimitAmount, &limit.CreatedAt, &limit.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %s", err)
			return nil, err
		}
		limits = append(limits, limit)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %s", err)
		return nil, err
	}

	return limits, nil
}
