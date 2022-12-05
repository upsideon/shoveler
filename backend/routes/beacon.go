package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	Id       string `json:"id"`
	OwnerId  string `json:"owner_id"`
	HelperId string `json:"helper_id"`
	Address  string `json:"address"`
}

type BeaconCreationRequest struct {
	Address string `form:"address" json:"address" binding:"required"`
}

type BeaconController struct {
	Database gocqlx.Session
}

func (c *BeaconController) Create(context *gin.Context) {
	var creationRequest BeaconCreationRequest

	ownerId, exists := context.Get("accountId")
	if !exists {
		context.String(http.StatusForbidden, "Forbidden")
		return
	}

	if err := context.BindJSON(&creationRequest); err != nil {
		log.Print(err)
		context.String(http.StatusBadRequest, "Bad Request")
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	beacon := Beacon{
		Id:      id.String(),
		Address: creationRequest.Address,
		OwnerId: ownerId.(string),
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
	ownerIdFromContext, exists := context.Get("accountId")
	if !exists {
		context.String(http.StatusForbidden, "Forbidden")
		return
	}

	ownerId := ownerIdFromContext.(string)

	includeOwnerStr := context.DefaultQuery("includeOwner", "true")
	includeOwner, err := strconv.ParseBool(includeOwnerStr)
	if err != nil {
		context.String(http.StatusBadRequest, "Bad Request")
		return
	}

	beacon := Beacon{
		OwnerId: ownerId,
	}

	selectBeaconsBuilder := qb.Select(
		"shoveler.beacons",
	).Columns(
		"id", "owner_id", "helper_id", "address",
	)

	if includeOwner {
		selectBeaconsBuilder = selectBeaconsBuilder.Where(qb.Eq("owner_id"))
	}

	selectBeacons := selectBeaconsBuilder.Query(c.Database).BindStruct(&beacon).Iter()

	allBeacons := []Beacon{}

	beacon = Beacon{}
	for selectBeacons.StructScan(&beacon) {
		if !includeOwner && beacon.OwnerId == ownerId {
			continue
		}
		allBeacons = append(allBeacons, beacon)
	}

	jsonBeacons, err := json.Marshal(allBeacons)
	if err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.JSON(http.StatusOK, string(jsonBeacons))
}

func (c *BeaconController) Delete(context *gin.Context) {
	ownerId, exists := context.Get("accountId")
	if !exists {
		context.String(http.StatusForbidden, "Forbidden")
		return
	}

	beacon := Beacon{
		Id:      context.Param("id"),
		OwnerId: ownerId.(string),
	}

	// TODO - It's troublesome that Cassandra doesn't allow us to include a
	// filter on owner_id as well for deletes. This means that, as of now,
	// if someone obtains an ID for a beacon that they can delete it without
	// being the owner. A solution for this must be determined.
	deleteBeacon := qb.Delete(
		"shoveler.beacons",
	).Where(
		qb.Eq("id"),
	).Query(c.Database).BindStruct(&beacon)

	if err := deleteBeacon.ExecRelease(); err != nil {
		log.Print(err)
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.String(http.StatusOK, "OK")
}
