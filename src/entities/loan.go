package entities

import "time"

type Loan struct {
	ID                 int64     `json:"id"`
	CustomerID         int64     `json:"customer_id"`
	Amount             float64   `json:"amount"`
	Tenor              int       `json:"tenor"`
	Purpose            string    `json:"purpose,omitempty"`
	InterestRate       float64   `json:"interest_rate"`
	MonthlyInstallment float64   `json:"monthly_installment"`
	RemainingBalance   float64   `json:"remaining_balance"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ParamsCreateLoan struct {
	Amount  float64 `json:"amount" validate:"required,gt=0"`
	Tenor   int     `json:"tenor" validate:"required,min=1"`
	Purpose string  `json:"purpose,omitempty"`
}

type ParamsCreatePaymentLoan struct {
	LoanId float64 `json:"loan_id"`
	Amount int     `json:"amount"`
}

type ParamsLoan struct {
	ID                 int64
	CustomerID         int64
	Amount             float64
	Tenor              int
	Purpose            string
	InterestRate       float64
	MonthlyInstallment float64
	RemainingBalance   float64
	Status             string
}

type ParamsPaymentLoan struct {
	CustomerId       float64
	LoanId           float64
	Amount           int
	RemainingBalance float64
	Status           string
}

type ResultCreateLoan struct {
	ID                 int64   `json:"loan_id"`
	Amount             float64 `json:"amount"`
	Tenor              int     `json:"tenor"`
	Interest           float64 `json:"interest"`
	MonthlyInstallment float64 `json:"monthly_installment"`
}

type MeValidator struct {
	ID int `mapstructure:"user_id" validate:"required"`
}
