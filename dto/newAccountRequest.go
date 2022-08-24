package dto

import "github.com/sushilparajuli/go-banking/errors"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors.AppError {
	if r.Amount < 5000 {
		return errors.NewValidationError("To open a new account you need to deposit at least 5000")
	}
	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errors.NewValidationError("Account type should be either checking, saving")
	}
	return nil
}
