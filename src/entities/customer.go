package entities

import "time"

type Customer struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	NIK         string    `json:"nik"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      float64   `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ParamsCreateCustomer struct {
	FullName    string  `json:"full_name" validate:"required,min=5,max=255"`
	LegalName   string  `json:"legal_name" validate:"required,min=5,max=255"`
	NIK         string  `json:"nik" validate:"required,len=20"`
	BirthPlace  string  `json:"birth_place" validate:"required,min=3,max=100"`
	BirthDate   string  `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Salary      float64 `json:"salary" validate:"required,gt=0"`
	KTPPhoto    string  `json:"ktp_photo" validate:"required,url"`
	SelfiePhoto string  `json:"selfie_photo" validate:"required,url"`
}

type ParamsCustomer struct {
	ID          int64
	FullName    string
	LegalName   string
	NIK         string
	BirthPlace  string
	BirthDate   string
	Salary      float64
	KTPPhoto    string
	SelfiePhoto string
}

type ResultRegsiterCustomer struct {
	ID int64 `json:"id"`
}
