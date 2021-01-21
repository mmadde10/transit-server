package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func main() {

	var dbURI = "mongodb+srv://mmadden:1sx1dCeHR17yv2Tf@cluster0.8h0op.gcp.mongodb.net/transit?retryWrites=true&w=majority"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(dbURI)
	mongoClient, _ = mongo.Connect(ctx, clientOptions)

	r := router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
