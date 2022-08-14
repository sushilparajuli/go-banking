package domain

import "github.com/sushilparajuli/go-banking/errors"

type Customer struct {
	Id          string `json:"id"db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth"db:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errors.AppError)
	ById(string) (*Customer, *errors.AppError)
}
