package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PabloGilvan/transaction/cmd/helpers"
	"github.com/PabloGilvan/transaction/commons"
	"github.com/PabloGilvan/transaction/internal/config/global"
	"github.com/PabloGilvan/transaction/internal/db"
	accountService "github.com/PabloGilvan/transaction/internal/services/account"
	"github.com/PabloGilvan/transaction/pkg/domains/account"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

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

func (s *AccountSuite) initTest() {
	global.ViperConfig()
	s.helperMock = new(db.MockDatabaseManager)
	s.repository = account.NewAccountRepository(s.helperMock)
	s.service = accountService.NewAccountService(s.repository)
	s.controller = NewAccountController(s.service)
	s.accountID = uuid.New().String()
	s.model = account.Account{
		ID:             s.accountID,
		Number:         "000010220234567",
		DocumentNumber: "00811100021",
		Active:         true,
		CreateDate:     time.Now(),
		UpdateDate:     time.Now(),
	}
	var err error
	s.db, s.dbMock, err = sqlmock.New()
	assert.Nil(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       s.db,
	})

	s.gdb, err = gorm.Open(dialector, &gorm.Config{})
	assert.Nil(s.T(), err)

	s.router = gin.Default()
	s.controller.Router(s.router.Group("/v1"))
}

func TestAccountSuiteSuite(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}

func (s *AccountSuite) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)
	return rr
}

func (s *AccountSuite) addLoadAccountRow() {
	s.rows = sqlmock.NewRows(
		[]string{
			"id",
			"number",
			"document_number",
			"active",
			"create_date",
			"update_date",
		}).AddRow(s.model.ID, s.model.Number, s.model.DocumentNumber, s.model.Active, s.model.CreateDate, s.model.UpdateDate)
}

func (s *AccountSuite) TestGetAccount() {
	s.T().Run("test get an account with success", func(t *testing.T) {
		s.initTest()
		s.addLoadAccountRow()

		s.helperMock.On("GetDatabaseConnection").Return(s.gdb, nil)

		const sqlSelect = `SELECT * FROM "accounts" WHERE id = $1`

		s.dbMock.ExpectQuery(regexp.QuoteMeta(sqlSelect)).WithArgs(s.accountID).WillReturnRows(s.rows)

		request := httptest.NewRequest("GET", fmt.Sprintf("/v1/accounts/%s", s.accountID), nil)

		responseRecorder := s.executeRequest(request)

		if err := s.dbMock.ExpectationsWereMet(); err != nil {
			s.T().Errorf("there were unfulfilled expectations: %s", err)
		}
		bodyResp, err := ioutil.ReadAll(responseRecorder.Body)
		assert.Nil(s.T(), err)

		var response accountService.AccountResponse
		err = json.Unmarshal(bodyResp, &response)
		assert.Nil(s.T(), err)

		assert.Equal(s.T(), http.StatusOK, responseRecorder.Result().StatusCode)
		assert.Equal(s.T(), s.model.ID, response.ID)
		assert.Equal(s.T(), s.model.Number, response.Number)
		assert.Equal(s.T(), s.model.DocumentNumber, response.DocumentNumber)
	})
	s.T().Run("test get an account with bad request when account is inactive", func(t *testing.T) {
		s.initTest()
		s.model.Active = false

		s.addLoadAccountRow()

		s.helperMock.On("GetDatabaseConnection").Return(s.gdb, nil)

		const sqlSelect = `SELECT * FROM "accounts" WHERE id = $1`

		s.dbMock.ExpectQuery(regexp.QuoteMeta(sqlSelect)).WithArgs(s.accountID).WillReturnRows(s.rows)

		request := httptest.NewRequest("GET", fmt.Sprintf("/v1/accounts/%s", s.accountID), nil)

		responseRecorder := s.executeRequest(request)

		if err := s.dbMock.ExpectationsWereMet(); err != nil {
			s.T().Errorf("there were unfulfilled expectations: %s", err)
		}
		bodyResp, err := ioutil.ReadAll(responseRecorder.Body)
		assert.Nil(s.T(), err)

		var response helpers.CustomResponseError
		err = json.Unmarshal(bodyResp, &response)
		assert.Nil(s.T(), err)

		assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Result().StatusCode)
		assert.Equal(s.T(), commons.ErrAccountInactive.Error(), response.ErrorMessage)
	})
	s.T().Run("test get an account with not found when account do not exist", func(t *testing.T) {
		s.initTest()
		s.model.Active = false

		s.addLoadAccountRow()

		s.helperMock.On("GetDatabaseConnection").Return(s.gdb, nil)

		const sqlSelect = `SELECT * FROM "accounts" WHERE id = $1`

		s.dbMock.ExpectQuery(regexp.QuoteMeta(sqlSelect)).WithArgs(s.accountID).WillReturnRows(sqlmock.NewRows(nil))

		request := httptest.NewRequest("GET", fmt.Sprintf("/v1/accounts/%s", s.accountID), nil)

		responseRecorder := s.executeRequest(request)

		if err := s.dbMock.ExpectationsWereMet(); err != nil {
			s.T().Errorf("there were unfulfilled expectations: %s", err)
		}
		bodyResp, err := ioutil.ReadAll(responseRecorder.Body)
		assert.Nil(s.T(), err)

		var response helpers.CustomResponseError
		err = json.Unmarshal(bodyResp, &response)
		assert.Nil(s.T(), err)

		assert.Equal(s.T(), http.StatusNotFound, responseRecorder.Result().StatusCode)
		assert.Equal(s.T(), commons.ErrAccountNotFound.Error(), response.ErrorMessage)
	})
}
