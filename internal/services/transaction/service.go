package transaction

import (
	"context"
	"github.com/PabloGilvan/transaction/commons"
	"github.com/PabloGilvan/transaction/internal/services/account"
	"github.com/PabloGilvan/transaction/pkg/domains/operation"
	"github.com/PabloGilvan/transaction/pkg/domains/transaction"
	"github.com/google/uuid"
	"time"
)

type TransactionService interface {
	SaveTransaction(ctx context.Context, persist TransactionPersist) (*TransactionResponse, error)
	PaymentsEvent(background context.Context, persist TransactionPersist) (*TransactionResponse, error)
}

type TransactionServiceImpl struct {
	Repository          transaction.TransactionRepository
	OperationRepository operation.OperationTypeRepository
	AccountService      account.AccountService
}

func NewTransactionService(repo transaction.TransactionRepository,
	operationRepo operation.OperationTypeRepository,
	accountService account.AccountService) TransactionService {
	return TransactionServiceImpl{
		Repository:          repo,
		OperationRepository: operationRepo,
		AccountService:      accountService,
	}
}

func (service TransactionServiceImpl) SaveTransaction(ctx context.Context, persist TransactionPersist) (*TransactionResponse, error) {
	accountDTO, err := service.AccountService.LoadAccount(ctx, persist.AccountId)
	if err != nil {
		return nil, err
	}

	op, err := service.OperationRepository.LoadOperation(persist.OperationTypeId)

	if err != nil {
		return nil, err
	}

	var transactionIdentifier = uuid.New().String()

	amount, err := calculateAndValidateCreditLimitForOperation(persist.Amount, accountDTO.AvailableCreditLimit, *op)
	if err != nil {
		return nil, err
	}

	_, err = service.Repository.SaveTransaction(transaction.Transaction{
		Id:              transactionIdentifier,
		AccountID:       persist.AccountId,
		OperationTypeID: persist.OperationTypeId,
		Amount:          amount,
		Balance:         amount,
		EventDate:       time.Now(),
		Approved:        true,
	})

	if err != nil {
		return nil, err
	}

	if err = service.AccountService.UpdateAccountLimit(ctx, accountDTO.ID, accountDTO.AvailableCreditLimit+amount); err != nil {
		return nil, err
	}

	return &TransactionResponse{TransactionIdentifier: transactionIdentifier}, nil
}

func (service TransactionServiceImpl) PaymentsEvent(ctx context.Context, transactionRequest TransactionPersist) (*TransactionResponse, error) {
	accountDTO, err := service.AccountService.LoadAccount(ctx, transactionRequest.AccountId)
	if err != nil {
		return nil, err
	}

	pendingTransactions, err := service.Repository.LoadPendingTransactionsSorted(accountDTO.ID)
	if err != nil {
		return nil, err
	}

	pendingTransactions = filterTransactionsToPay(pendingTransactions, transactionRequest.Amount)

	var transactionIdentifier = uuid.New().String()

}

func filterTransactionsToPay(trans transaction.Transaction, creditValue float64) []transaction.Transaction {
	for pos {
		trans.Balance = trans.SubtractValue(creditValue)
		creditValue := creditValue - trans.Amount
	}
	return nil
}

func calculateAndValidateCreditLimitForOperation(amount float64, balance float64, op operation.OperationType) (float64, error) {
	if op.ShouldUseMultiplicationFactor && amount > 0 {
		amount = amount * float64(op.MultiplicationFactor)
	}

	if amount < 0 && (balance+amount) < 0 {
		return 0, commons.ErrLimitExceeded
	}
	return amount, nil
}
