package container

import (
	"github.com/PabloGilvan/transaction/internal/db"
	"github.com/PabloGilvan/transaction/internal/services/account"
	"github.com/PabloGilvan/transaction/pkg/domains"
	account2 "github.com/PabloGilvan/transaction/pkg/domains/account"
)

type components struct {
	DatabaseManager db.DatabaseManager
}

type services struct {
	AccountService account.AccountService
}

type Dependency struct {
	Components components
	Services   services
}

func Injector() Dependency {
	var dbm = db.NewDatabaseManager()

	domains.StartMigrationPlan(dbm)

	/*var operationRepository = operation.NewOperationTypeRepository(dbm)
	var transactionRepository = transaction.NewTransactionRepository(dbm)*/
	var accountRepository = account2.NewAccountRepository(dbm)

	return Dependency{
		Components: components{
			DatabaseManager: dbm,
		},
		Services: services{
			AccountService: account.NewAccountService(accountRepository),
		},
	}
}
