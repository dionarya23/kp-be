package limitrepository

import (
	"database/sql"

	"github.com/dionarya23/kredit-plus/src/entities"
)

type sLimitRepository struct {
	DB *sql.DB
}

type LimitRepository interface {
	FindMany(customerId int64, tenor int) ([]*entities.Limit, error)
}

func New(db *sql.DB) LimitRepository {
	return &sLimitRepository{DB: db}
}
