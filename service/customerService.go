package service

import (
	"github.com/sushilparajuli/go-banking/domain"
	"github.com/sushilparajuli/go-banking/dto"
	"github.com/sushilparajuli/go-banking/errors"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customersDto := make([]dto.CustomerResponse, 0)
	all, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	for _, v := range all {
		response := v.ToDto()
		customersDto = append(customersDto, *response)
	}
	return customersDto, nil
}
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
