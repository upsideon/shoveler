package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"shoveler/datastore"
	"shoveler/models"
)

type AccountCreationRequest struct {
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required"`
}

type AccountController struct {
	AccountsDatastore datastore.AccountsDatastore
}

func (c *AccountController) Create(context *gin.Context) {
	var creationRequest AccountCreationRequest

	if err := context.BindJSON(&creationRequest); err != nil {
		log.Print(err)
		context.String(http.StatusBadRequest, "Bad Request")
		return
	}

	// Check if account already exists with the provided email.
	// TODO - Consider adding a reCaptcha to prevent scripts from determining
	// e-mail addresses for accounts.
	existingAccount, err := c.AccountsDatastore.GetAccount(creationRequest.Email)
	if err != nil && !errors.Is(err, datastore.ErrNotFound) {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if existingAccount != nil {
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

	account := models.Account{
		Email:        creationRequest.Email,
		Id:           id.String(),
		PasswordHash: string(hashedPassword),
	}

	err = c.AccountsDatastore.CreateAccount(&account)
	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.String(http.StatusOK, "OK")
}
