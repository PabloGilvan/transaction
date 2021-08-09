package transaction

import (
	"context"
	"github.com/PabloGilvan/transaction/pkg/domains/operation"
	"github.com/PabloGilvan/transaction/pkg/domains/transaction"
	"github.com/google/uuid"
	"time"
)

type TransactionService interface {
	SaveTransaction(ctx context.Context, persist TransactionPersist) (*TransactionResponse, error)
}

type TransactionServiceImpl struct {
	Repository          transaction.TransactionRepository
	OperationRepository operation.OperationTypeRepository
}

func NewTransactionService(repo transaction.TransactionRepository, operationRepo operation.OperationTypeRepository) TransactionService {
	return TransactionServiceImpl{
		Repository:          repo,
		OperationRepository: operationRepo,
	}
}

func (service TransactionServiceImpl) SaveTransaction(ctx context.Context, persist TransactionPersist) (*TransactionResponse, error) {

	var op, err = service.OperationRepository.LoadOperation(persist.OperationTypeId)

	if err != nil {
		return nil, err
	}

	var transactionIdentifier = uuid.New().String()
	var amount = persist.Amount
	if op.ShouldUseMultiplicationFactor && amount > 0 {
		amount = amount * float64(op.MultiplicationFactor)
	}

	_, err = service.Repository.SaveTransaction(transaction.Transaction{
		Id:              transactionIdentifier,
		AccountID:       persist.AccountId,
		OperationTypeID: persist.OperationTypeId,
		Amount:          amount,
		EventDate:       time.Now(),
		Approved:        true,
	})

	if err != nil {
		return nil, err
	}

	return &TransactionResponse{TransactionIdentifier: transactionIdentifier}, nil
}
