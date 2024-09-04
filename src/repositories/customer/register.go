package customerrepository

import (
	"log"
	"time"

	"github.com/dionarya23/kredit-plus/src/entities"
)

func (i *sCustomerRepository) Create(p *entities.ParamsCreateCustomer) (*entities.Customer, error) {
	birthDate, err := time.Parse("2006-01-02", p.BirthDate)
	if err != nil {
		log.Printf("Error parsing birth date: %s", err)
		return nil, err
	}

	query := "INSERT INTO customers (full_name, legal_name, nik, birth_place, birth_date, salary, ktp_photo, selfie_photo) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := i.DB.Exec(query, p.FullName, p.LegalName, p.NIK, p.BirthPlace, birthDate, p.Salary, p.KTPPhoto, p.SelfiePhoto)
	if err != nil {
		log.Printf("Error inserting customer: %s", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %s", err)
		return nil, err
	}

	customer := &entities.Customer{
		ID:          id,
		FullName:    p.FullName,
		LegalName:   p.LegalName,
		NIK:         p.NIK,
		BirthPlace:  p.BirthPlace,
		BirthDate:   birthDate,
		Salary:      p.Salary,
		KTPPhoto:    p.KTPPhoto,
		SelfiePhoto: p.SelfiePhoto,
	}

	return customer, nil
}
