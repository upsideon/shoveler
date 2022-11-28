package main

import (
	"log"
	"shoveler/middleware"
	"shoveler/routes"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func main() {
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "shoveler"

	databaseSession, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}
	defer databaseSession.Close()

	router := initializeRouter(databaseSession)
	router.Run(":8080")
}

func initializeRouter(db gocqlx.Session) *gin.Engine {
	accountController := routes.AccountController{
		Database: db,
	}

	beaconController := routes.BeaconController{
		Database: db,
	}

	loginController := routes.LoginController{
		Database: db,
	}

	authorization := middleware.Authorization()

	router := gin.Default()
	router.POST("/accounts", accountController.Create)
	router.POST("/login", loginController.Login)
	router.POST("/beacons", authorization, beaconController.Create)
	router.GET("/beacons", authorization, beaconController.List)
	router.GET("/test", middleware.Authorization(), func(ctx *gin.Context) {})
	return router
}
