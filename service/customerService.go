package service

import (
	"github.com/sushilparajuli/go-banking/domain"
	"github.com/sushilparajuli/go-banking/errors"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errors.AppError)
	GetCustomer(string) (*domain.Customer, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errors.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
