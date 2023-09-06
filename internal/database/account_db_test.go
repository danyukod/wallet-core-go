package database

import (
	"database/sql"
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	clientDB  *ClientDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance INT, created_at DATE)")

	s.accountDB = NewAccountDB(db)
	s.clientDB = NewClientDB(db)
	s.client, _ = entity.NewClient("Danilo", "d@email.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSaveAccount() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.SaveAccount(account)
	s.NoError(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.clientDB.SaveClient(s.client)
	account := entity.NewAccount(s.client)
	err := s.accountDB.SaveAccount(account)
	s.NoError(err)

	account, err = s.accountDB.FindByID(account.ID)
	s.NoError(err)
	s.Equal(account.ID, account.ID)
	s.Equal(account.Client.ID, account.Client.ID)
	s.Equal(account.Balance, account.Balance)
	s.Equal(account.CreatedAt, account.CreatedAt)
}
