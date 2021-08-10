# Transaction

Projeto desenvolvido em [Golang 1.6](https://golang.org/doc/go1.6).
Utilizando libs como [gin](https://github.com/gin-gonic/gin), [gorm](https://gorm.io/) e também [swaggo](https://github.com/swaggo/swag) para gerar a documentação da API.

## Configurações

O projeto contém um `Dockerfile`, que gera a imagem do projeto, mas o mesmo precisa de uma base **PostgreSQL** para ser executado. O default é uma *config* local:
``` yaml
app:
  port: 8080
  database:
    host: localhost
    port: 5432
    user: postgres
    password: postgres
    name: transaction
    sslmode: disable

```
Para rodar a imagem com outra base, somente definir as variáveis de ambiente:
```
APP_DATABASE_HOST=127.0.0.1
APP_DATABASE_USER=postgres
....
```
## Server
A porta default é `8080`, mas como com as variáveis de base, pode ser definida com uma variável de ambiente:
```
APP_PORT=8181
```

## Paths

```go
// @Title saveTransaction
// @Summary Persist a transaction
// @Router /transactions/{id} [get]
func (repo OperationTypeRepositoryImpl) SaveOperation(operationModel OperationType)

// @Title createAccount
// @Summary Creates an Account
// @Router /accounts [post]
func (crtl AccountController) CreateAccount(c *gin.Context)

// @Title loadAccount
// @Summary Load an account
// @Router /accounts/{id} [get]
func (crtl AccountController) LoadAccount(c *gin.Context)
```

## Test
Casos de testes foram criados somente na camada do *controller*, visando um teste que valide toda a integração das camadas, somente a camada de base que tem seu comportamento `mocado`.
Dado o tempo acabei não criando testes para o `TransactionController`.
```go
type AccountSuite struct {
	suite.Suite
	ctx             context.Context
	repository      account.AccountRepository
	service         accountService.AccountService
	controller      AccountController
	router          *gin.Engine
	gormDB          *gorm.DB
	db              *sql.DB
	dbMock          sqlmock.Sqlmock
	helperMock      *db.MockDatabaseManager
	gdb             *gorm.DB
	accumulatedRows *sqlmock.Rows
	rows            *sqlmock.Rows
	accountID       string
	model           account.Account
}
```

## Observações
> Acabei criando um `docker-compose` para o projeto, acontece que estou usando um `Dockerfile` para carregar a imagem do projeto, por alguma falha de configuração minha, ele acaba instanciando a aplicação sem ainda ter terminado de criar a base, o que da problemas. No `docker-compose` coloquei um *script* que cria a base.

> Criei uma estrutura de Migrations, atualmente utilizo migrantios via *script*, resolvi por testar migration via  *model*. Com essa abordagem, caso a base esteja inacessível eu *jogo* um *panic*, o que acarretou no problema acima. A Migration se dá no trecho de código a seguir:
```go
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
```
