package database

import (
	"database/sql"
	"github.com/danyukod/wallet-core-go/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSaveClient() {
	client, _ := entity.NewClient("John", "j@d.com")
	s.Nil(s.clientDB.SaveClient(client))
}

func (s *ClientDBTestSuite) TestGetClient() {
	client, _ := entity.NewClient("John", "j@d.com")
	s.Nil(s.clientDB.SaveClient(client))
	c, err := s.clientDB.GetClient(client.ID)
	s.Nil(err)
	s.Equal(client.ID, c.ID)
	s.Equal(client.Name, c.Name)
	s.Equal(client.Email, c.Email)
}
