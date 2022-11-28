package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

var beaconsMetadata = table.Metadata{
	Name:    "beacons",
	Columns: []string{"id", "owner_id", "helper_id", "address"},
}

var beaconsTable = table.New(beaconsMetadata)

type Beacon struct {
	Id       string
	OwnerId  string
	HelperId string
	Address  string
}

type BeaconCreationRequest struct {
	Address string `form:"address" json:"address" binding:"required"`
}

type BeaconController struct {
	Database gocqlx.Session
}

func (c *BeaconController) Create(context *gin.Context) {
	var creationRequest BeaconCreationRequest

	if err := context.BindJSON(&creationRequest); err != nil {
		log.Print(err)
		context.String(http.StatusBadRequest, "Bad Request")
		return
	}

	token, exists := context.Get("token")
	if !exists {
		context.String(http.StatusForbidden, "Forbidden")
		return
	}

	claims, ok := token.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		log.Print(errors.New("Unable to access token claims."))
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	address := creationRequest.Address
	log.Print(claims)
	ownerId := claims["sub"].(string)

	beacon := Beacon{
		Id:      id.String(),
		Address: address,
		OwnerId: ownerId,
	}

	insertBeacon := qb.Insert(
		"shoveler.beacons",
	).Columns(
		"id", "owner_id", "helper_id", "address",
	).Query(c.Database).BindStruct(beacon)

	if err := insertBeacon.ExecRelease(); err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.String(http.StatusOK, "OK")
}

func (c *BeaconController) List(context *gin.Context) {
}
