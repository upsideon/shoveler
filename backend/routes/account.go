package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"golang.org/x/crypto/bcrypt"
)

var accountsMetadata = table.Metadata{
	Name:    "accounts",
	Columns: []string{"id", "email", "password_hash"},
}

var accountsTable = table.New(accountsMetadata)

type Account struct {
	Id           string
	Email        string
	PasswordHash string
}

type AccountCreationRequest struct {
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	PasswordConfirm string `form:"password-confirm" json:"password_confirm" binding:"required"`
}

type AccountController struct {
	Database gocqlx.Session
}

func (c *AccountController) Create(context *gin.Context) {
	var creationRequest AccountCreationRequest

	if err := context.BindJSON(&creationRequest); err != nil {
		log.Print(err)
		context.String(http.StatusBadRequest, "Bad Request")
		return
	}

	account := Account{Email: creationRequest.Email}

	// Checking if account already exists.

	selectAccount := qb.Select(
		"shoveler.accounts",
	).CountAll().Where(qb.Eq("email")).Query(c.Database).BindStruct(
		account,
	)

	var count int

	if err := selectAccount.GetRelease(&count); err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if count != 0 {
		log.Print(account)
		// TODO - Add a reCaptcha to prevent scripts from determining e-mail addresses for accounts.
		context.String(
			http.StatusBadRequest,
			"Bad Request: An account already exists with that e-mail address.",
		)
		return
	}

	// Checking if password and password confirmation match.
	if creationRequest.Password != creationRequest.PasswordConfirm {
		context.String(
			http.StatusBadRequest,
			"Bad Request: Password and password confirmation do not match.",
		)
		return
	}

	// Hashing password.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(creationRequest.Password),
		bcrypt.MinCost,
	)
	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	account.Id = id.String()
	account.PasswordHash = string(hashedPassword)

	insertAccount := qb.Insert(
		"shoveler.accounts",
	).Columns(
		"id", "email", "password_hash",
	).Query(c.Database).BindStruct(account)

	if err := insertAccount.ExecRelease(); err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.String(http.StatusOK, "OK")
}
