package account

import (
	"github.com/PabloGilvan/transaction/commons"
	"github.com/PabloGilvan/transaction/internal/db"
)

type AccountRepository interface {
	SaveAccount(account Account) (*Account, error)
	LoadAccount(id string) (*Account, error)
	UpdateAccountLimit(id string, limit float64) error
}

type AccountRepositoryImpl struct {
	DatabaseManager db.DatabaseManager
}

func NewAccountRepository(dbm db.DatabaseManager) AccountRepository {
	return AccountRepositoryImpl{
		DatabaseManager: dbm,
	}
}

func (repo AccountRepositoryImpl) LoadAccount(id string) (*Account, error) {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	var account Account
	conn.First(&account, " id = ? ", id)
	if len(account.ID) == 0 {
		return nil, commons.ErrAccountNotFound
	}

	if !account.Active {
		return nil, commons.ErrAccountInactive
	}

	return &account, nil
}

func (repo AccountRepositoryImpl) SaveAccount(accountModel Account) (*Account, error) {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	err = conn.Save(accountModel).Error
	if err != nil {
		return nil, err
	}

	return &accountModel, nil
}

func (repo AccountRepositoryImpl) UpdateAccountLimit(id string, limit float64) error {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return err
	}

	var account = Account{ID: id}
	return conn.Model(&account).Update("available_credit_limit", limit).Error
}
