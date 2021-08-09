package domains

import (
	"github.com/PabloGilvan/transaction/internal/db"
	"github.com/PabloGilvan/transaction/pkg/domains/account"
	"github.com/PabloGilvan/transaction/pkg/domains/operation"
	"github.com/PabloGilvan/transaction/pkg/domains/transaction"
	"github.com/google/uuid"
	"time"
)

func StartMigrationPlan(dbm db.DatabaseManager) {
	migrateAccount(dbm)
	migrateOperationType(dbm)

	err := dbm.Migrate(&transaction.Transaction{})
	if err != nil {
		panic(err)
	}
}

func migrateAccount(dbm db.DatabaseManager) {
	err := dbm.Migrate(&account.Account{})
	if err != nil {
		panic(err)
	}

	con, err := dbm.GetDatabaseConnection()
	if err != nil {
		panic(err)
	}

	con.Save(&account.Account{
		ID:             uuid.New().String(),
		DocumentNumber: "00000001",
		Active:         true,
		CreateDate:     time.Now(),
		UpdateDate:     time.Now(),
	})
}

func migrateOperationType(dbm db.DatabaseManager) {
	err := dbm.Migrate(&operation.OperationType{})
	if err != nil {
		panic(err)
	}

	con, err := dbm.GetDatabaseConnection()
	if err != nil {
		panic(err)
	}

	con.Save(&operation.OperationType{
		ID:          uuid.New().String(),
		Name:        "Simple",
		Description: "Test case",
		Active:      true,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
	})
}
