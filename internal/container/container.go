package container

import (
	"github.com/PabloGilvan/transaction/internal/db"
	accountService "github.com/PabloGilvan/transaction/internal/services/account"
	transactionService "github.com/PabloGilvan/transaction/internal/services/transaction"
	"github.com/PabloGilvan/transaction/pkg/domains"
	"github.com/PabloGilvan/transaction/pkg/domains/account"
	"github.com/PabloGilvan/transaction/pkg/domains/operation"
	"github.com/PabloGilvan/transaction/pkg/domains/transaction"
)

type components struct {
	DatabaseManager db.DatabaseManager
}

type services struct {
	AccountService     accountService.AccountService
	TransactionService transactionService.TransactionService
}

type Dependency struct {
	Components components
	Services   services
}

func Injector() Dependency {
	var dbm = db.NewDatabaseManager()

	domains.StartMigrationPlan(dbm)

	var operationRepository = operation.NewOperationTypeRepository(dbm)
	var transactionRepository = transaction.NewTransactionRepository(dbm)

	var accountService = accountService.NewAccountService(account.NewAccountRepository(dbm))

	return Dependency{
		Components: components{
			DatabaseManager: dbm,
		},
		Services: services{
			AccountService:     accountService,
			TransactionService: transactionService.NewTransactionService(transactionRepository, operationRepository, accountService),
		},
	}
}
