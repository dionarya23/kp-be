package customerrepository

import (
	"bytes"
	"database/sql"
	"log"
	"strings"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sCustomerRepository) FindOne(filters *entities.ParamsCustomer) (*entities.Customer, error) {
	var query bytes.Buffer
	query.WriteString("SELECT id, full_name, phone_number, password FROM Customers WHERE ")

	params := []interface{}{}
	conditions := []string{}

	addCondition := func(condition string, param interface{}) {
		conditions = append(conditions, condition+" = ?")
		params = append(params, param)
	}

	if filters.ID != 0 {
		addCondition("id", filters.ID)
	}
	if filters.PhoneNumber != "" {
		addCondition("phone_number", filters.PhoneNumber)
	}

	query.WriteString(strings.Join(conditions, " AND "))
	query.WriteString(" LIMIT 1")

	row := i.DB.QueryRow(query.String(), params...)

	var customer entities.Customer
	err := row.Scan(&customer.ID, &customer.FullName, &customer.PhoneNumber, &customer.Password)

	if err != nil {
		log.Printf("Error finding customer: %s", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}
