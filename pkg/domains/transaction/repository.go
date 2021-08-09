package transaction

import (
	"github.com/PabloGilvan/transaction/internal/db"
)

type TransactionRepository interface {
	SaveTransaction(transactionModel Transaction) (*Transaction, error)
}

type TransactionRepositoryImpl struct {
	DatabaseManager db.DatabaseManager
}

func NewTransactionRepository(dbm db.DatabaseManager) TransactionRepository {
	return TransactionRepositoryImpl{
		DatabaseManager: dbm,
	}
}

func (repo TransactionRepositoryImpl) SaveTransaction(transactionModel Transaction) (*Transaction, error) {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	conn.Save(transactionModel)

	return &transactionModel, nil
}
