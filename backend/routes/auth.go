package routes

import (
	"log"
	"net/http"
	"shoveler/auth"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginController struct {
	Database gocqlx.Session
}

func (c *LoginController) Login(context *gin.Context) {
	var loginRequest LoginRequest

	if err := context.BindJSON(&loginRequest); err != nil {
		log.Print(err)
		context.String(http.StatusBadRequest, "Bad Request")
		return
	}

	account := Account{
		Email: loginRequest.Email,
	}

	selectAccount := qb.Select(
		"shoveler.accounts",
	).Columns("password_hash").Where(qb.Eq("email")).Query(c.Database).BindStruct(
		account,
	)

	if err := selectAccount.GetRelease(&account); err != nil {
		if err == gocql.ErrNotFound {
			context.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(account.PasswordHash),
		[]byte(loginRequest.Password),
	); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			context.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	token, err := auth.GenerateToken(account.Id)

	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
