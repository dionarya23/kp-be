package customerrepository

import (
	"database/sql"

	"github.com/dionarya23/kredit-plus/src/entities"
)

type sCustomerRepository struct {
	DB *sql.DB
}

type CustomerRepository interface {
	Create(*entities.ParamsCreateCustomer) (*entities.Customer, error)
	IsExists(*entities.ParamsCustomer) (bool, error)
	FindOne(*entities.ParamsCustomer) (*entities.Customer, error)
}

func New(db *sql.DB) CustomerRepository {
	return &sCustomerRepository{DB: db}
}
