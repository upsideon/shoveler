package main

import (
	"log"
	"shoveler/middleware"
	"shoveler/routes"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"DELETE", "GET", "OPTIONS", "POST", "PUT"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/accounts", accountController.Create)
	router.POST("/login", loginController.Login)
	router.POST("/beacons", authorization, beaconController.Create)
	router.GET("/beacons", authorization, beaconController.List)
	router.DELETE("/beacons/:id", authorization, beaconController.Delete)
	router.GET("/test", middleware.Authorization(), func(ctx *gin.Context) {})
	return router
}
