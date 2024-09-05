package entities

import "time"

type Limit struct {
	ID          int64     `json:"id"`
	CustomerID  int64     `json:"customer_id"`
	Tenor       int       `json:"tenor"`
	LimitAmount float64   `json:"limit_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ParamsCreateLimit struct {
	CustomerID  int64   `json:"customer_id" validate:"required"`
	Tenor       int     `json:"tenor" validate:"required"`
	LimitAmount float64 `json:"limit_amount" validate:"required,gt=0"`
}

type ParamsLimit struct {
	ID          int64
	CustomerID  int64
	Tenor       int
	LimitAmount float64
}

type ResultCreateLimit struct {
	ID int64 `json:"id"`
}
