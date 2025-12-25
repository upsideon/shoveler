package datastore

import (
	"errors"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"

	"shoveler/models"
)

var (
	accountsMetadata = table.Metadata{
		Name:    "accounts",
		Columns: []string{"id", "email", "password_hash"},
	}

	accountsTable = table.New(accountsMetadata)
)

type CassandraAccountsDatastore struct {
	database gocqlx.Session
}

func NewCassandraAccountsDatastore(database gocqlx.Session) *CassandraAccountsDatastore {
	return &CassandraAccountsDatastore{database}
}

func (d *CassandraAccountsDatastore) CreateAccount(account *models.Account) error {
	insertAccount := qb.Insert(
		"shoveler.accounts",
	).Columns(
		"id", "email", "password_hash",
	).Query(d.database).BindStruct(account)

	return insertAccount.ExecRelease()
}

func (d *CassandraAccountsDatastore) GetAccount(email string) (*models.Account, error) {
	account := models.Account{Email: email}

	selectAccount := qb.Select(
		"shoveler.accounts",
	).Where(qb.Eq("email")).Query(d.database).BindStruct(
		account,
	)

	err := selectAccount.GetRelease(&account)
	if err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &account, err
}
