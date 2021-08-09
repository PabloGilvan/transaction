package domains

import (
	"github.com/PabloGilvan/transaction/internal/db"
	"github.com/PabloGilvan/transaction/pkg/domains/account"
	"github.com/PabloGilvan/transaction/pkg/domains/operation"
	"github.com/PabloGilvan/transaction/pkg/domains/transaction"
	"time"
)

func StartMigrationPlan(dbm db.DatabaseManager) {
	err := dbm.Migrate(&account.Account{})
	if err != nil {
		panic(err)
	}

	err = dbm.Migrate(&transaction.Transaction{})
	if err != nil {
		panic(err)
	}

	migrateOperationType(dbm)
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
		ID:                            1,
		Name:                          "CREDIT",
		Description:                   "",
		MultiplicationFactor:          1,
		ShouldUseMultiplicationFactor: false,
		Active:                        true,
		CreateDate:                    time.Now(),
		UpdateDate:                    time.Now(),
	})

	con.Save(&operation.OperationType{
		ID:                            2,
		Name:                          "DEBIT",
		Description:                   "",
		MultiplicationFactor:          -1,
		ShouldUseMultiplicationFactor: true,
		Active:                        true,
		CreateDate:                    time.Now(),
		UpdateDate:                    time.Now(),
	})
}
