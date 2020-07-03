package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var key = ""

func getInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	CurrentAppInfo := AppInfo{
		Name:    "transitserver",
		Version: "0.1.0",
	}
	json.NewEncoder(w).Encode(CurrentAppInfo)
}

func getAllLines(w http.ResponseWriter, r *http.Request) {

	var results []*Line
	db := mongoClient.Database("transit")
	collection := db.Collection("lines")
	findOptions := options.Find()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, decodeError := collection.Find(ctx, bson.D{{}}, findOptions)

	if decodeError != nil {
		log.Fatal(decodeError)
	}

	for cur.Next(ctx) {
		var elem Line
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(results)

}

func getLineByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	db := mongoClient.Database("transit")
	collection := db.Collection("lines")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.D{{
		"name",
		bson.D{{
			"$in",
			bson.A{name},
		}},
	}}

	var line Line

	decodeError := collection.FindOne(ctx, filter).Decode(&line)

	if decodeError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(line)
}

// Get Stops

func getStopByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db := mongoClient.Database("transit")
	collection := db.Collection("stops")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.D{{
		"STOP_ID",
		bson.D{{
			"$in",
			bson.A{id},
		}},
	}}

	var stop Stop

	decodeError := collection.FindOne(ctx, filter).Decode(&stop)

	if decodeError != nil {
		log.Fatal(decodeError)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(stop)
}

func getStopsByLineName(w http.ResponseWriter, r *http.Request) {

	print("we here")
	lineName := mux.Vars(r)["line"]
	db := mongoClient.Database("transit")
	collection := db.Collection("stops")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	var results []*Stop

	filter := bson.D{{
		lineName,
		bson.D{{
			"$in",
			bson.A{true},
		}},
	}}

	findOptions := options.Find()
	cur, decodeError := collection.Find(ctx, filter, findOptions)

	if decodeError != nil {
		log.Fatal(decodeError)
	}

	for cur.Next(ctx) {
		var elem Stop
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(results)
}

func getStopArrivals(w http.ResponseWriter, r *http.Request) {
	stopID := mux.Vars(r)["stopid"]
	url := "http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx?key=" + key + "&stpid=" + stopID + "&max=5&outputType=JSON"

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)

	var arrivals Ctatt
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	error := json.Unmarshal(body, &arrivals)

	if error != nil {
		log.Fatal(error)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(arrivals)

}
