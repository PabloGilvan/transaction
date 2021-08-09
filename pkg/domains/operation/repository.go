package operation

import (
	"github.com/PabloGilvan/transaction/internal/db"
	"gorm.io/gorm"
)

type OperationTypeRepository interface {
	SaveOperation(operationModel OperationType) (*OperationType, error)
	LoadAccount(id string) (*OperationType, error)
}

type OperationTypeRepositoryImpl struct {
	DatabaseManager db.DatabaseManager
}

func NewOperationTypeRepository(dbm db.DatabaseManager) OperationTypeRepository {
	return OperationTypeRepositoryImpl{
		DatabaseManager: dbm,
	}
}

func (repo OperationTypeRepositoryImpl) SaveOperation(operationModel OperationType) (*OperationType, error) {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	conn.Save(operationModel)

	return &operationModel, nil
}

func (repo OperationTypeRepositoryImpl) LoadAccount(id string) (*OperationType, error) {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	var operation OperationType
	conn.Find(operation, " id = ? and active == 1 ", id)
	if len(operation.ID) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &operation, nil
}
