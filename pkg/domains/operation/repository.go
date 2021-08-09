package operation

import (
	"github.com/PabloGilvan/transaction/commons"
	"github.com/PabloGilvan/transaction/internal/db"
)

type OperationTypeRepository interface {
	SaveOperation(operationModel OperationType) (*OperationType, error)
	LoadOperation(id int) (*OperationType, error)
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

func (repo OperationTypeRepositoryImpl) LoadOperation(id int) (*OperationType, error) {
	conn, err := repo.DatabaseManager.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	var operation OperationType
	conn.First(&operation, " id = ? ", id)
	if operation.ID == 0 {
		return nil, commons.ErrOperationNotFound
	}

	return &operation, nil
}
