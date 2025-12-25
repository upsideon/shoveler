package datastore

import (
	"errors"

	"shoveler/models"
)

var ErrNotFound = errors.New("not found")

type AccountsDatastore interface {
	CreateAccount(account *models.Account) error
	GetAccount(email string) (*models.Account, error)
}