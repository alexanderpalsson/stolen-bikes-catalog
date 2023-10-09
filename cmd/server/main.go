package main

import (
	"context"
	"fmt"
	"github.com/counter/officers/officershandler"
	"github.com/counter/officers/officerstorage"
	"github.com/counter/reports/reportshandler"
	"github.com/counter/reports/reportsstorage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName                 = "stolen-bikes-catalog"
	dbUri                  = "mongodb://mongodb:27017"
	officersCollectionName = "officers"
	reportsCollectionName  = "reports"
)

func main() {
	router := gin.Default()

	opts := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		fmt.Println("could not establish connection to mongoDB", err)
		panic(err)
	}

	officersStorage := officerstorage.NewStorage(client.Database(dbName).Collection(officersCollectionName))
	reportsStorage := reportsstorage.NewStorage(client.Database(dbName).Collection(reportsCollectionName))

	officerHandler := officershandler.NewHandler(officersStorage, reportsStorage)

	reportHandler := reportshandler.NewHandler(reportsStorage, officersStorage)

	// Endpoints for police officers
	router.POST("/officers", officerHandler.Create)
	router.PUT("/officers/:id", officerHandler.Update)
	router.DELETE("/officers/:id", officerHandler.Delete)
	router.GET("/officers", officerHandler.FindAll)

	// Endpoints for reports
	router.POST("/reports", reportHandler.Create)
	router.GET("/reports", reportHandler.FindAll)
	router.GET("/reports/:id", reportHandler.FindByID)
	router.PUT("/reports/:id/solved", reportHandler.ReportCaseSolved)

	// Start the server
	fmt.Println("Police Bike Recovery API is running on port 8080...")
	router.Run(":8080")

	if err := router.Run(); err != nil {
		fmt.Println("error running http handler", err)
	}
}
