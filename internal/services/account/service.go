package account

import (
	"context"
	"github.com/PabloGilvan/transaction/internal/services"
	"github.com/PabloGilvan/transaction/pkg/domains/account"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AccountService interface {
	SaveAccount(ctx context.Context, persist AccountPersist) (AccountResponse, error)
	LoadAccount(ctx context.Context, id string) (AccountResponse, error)
}

type AccountServiceImpl struct {
	Repository account.AccountRepository
}

func NewAccountService(repo account.AccountRepository) AccountService {
	return AccountServiceImpl{
		Repository: repo,
	}
}

func (service AccountServiceImpl) SaveAccount(ctx context.Context, persist AccountPersist) (AccountResponse, error) {

	accountNumber := services.GenerateAccountNumber(services.GenerateTraceNumber(), time.Now())

	model, err := service.Repository.SaveAccount(account.Account{
		ID:             uuid.New().String(),
		Number:         accountNumber,
		DocumentNumber: persist.DocumentNumber,
		Active:         true,
		CreateDate:     time.Now(),
		UpdateDate:     time.Now(),
	})

	if err != nil {
		return AccountResponse{}, err
	}

	return ConvertModelToResponse(*model), nil
}

func (service AccountServiceImpl) LoadAccount(ctx context.Context, uuid string) (AccountResponse, error) {
	model, err := service.Repository.LoadAccount(uuid)
	if err != nil {
		return AccountResponse{}, err
	}

	if model == nil {
		return AccountResponse{}, gorm.ErrRecordNotFound
	}
	return ConvertModelToResponse(*model), nil
}
