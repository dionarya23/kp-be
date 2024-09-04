package customerrepository

import (
	"log"
	"strings"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sCustomerRepository) IsExists(filters *entities.ParamsCustomer) (bool, error) {
	var sb strings.Builder
	var params []interface{}
	var conditions []string

	sb.WriteString("SELECT EXISTS (SELECT 1 FROM Customers WHERE ")

	if filters.ID != 0 {
		params = append(params, filters.ID)
		conditions = append(conditions, "id = ?")
	}

	if filters.NIK != "" {
		params = append(params, filters.NIK)
		conditions = append(conditions, "nik = ?")
	}

	if filters.PhoneNumber != "" {
		params = append(params, filters.PhoneNumber)
		conditions = append(conditions, "phone_number = ?")
	}

	sb.WriteString(strings.Join(conditions, " AND "))
	sb.WriteString(")")

	var exists bool
	err := i.DB.QueryRow(sb.String(), params...).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if customer exists: %s", err)
		return false, err
	}

	return exists, nil
}
