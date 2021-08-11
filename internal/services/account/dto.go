package account

import (
	"github.com/PabloGilvan/transaction/pkg/domains/account"
	"time"
)

type AccountPersist struct {
	DocumentNumber       string  `json:"document_number" binding:"required"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}

type AccountResponse struct {
	ID                   string  `json:"id"`
	Number               string  `json:"number"`
	DocumentNumber       string  `json:"document_number"`
	AvailableCreditLimit float64 `json:"available-credit-limit"`
	Active               bool    `json:"active"`
	CreateDate           time.Time
	UpdateDate           time.Time
}

func ConvertModelToResponse(account account.Account) AccountResponse {
	return AccountResponse{
		ID:                   account.ID,
		Number:               account.Number,
		DocumentNumber:       account.DocumentNumber,
		AvailableCreditLimit: account.AvailableCreditLimit,
		Active:               account.Active,
		CreateDate:           account.CreateDate,
		UpdateDate:           account.UpdateDate,
	}
}
