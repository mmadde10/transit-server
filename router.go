package main

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func router() *mux.Router {
	router := mux.NewRouter()

	//app info route
	router.HandleFunc("/api/info", getInfo).Methods("GET", "OPTIONS")

	// Line Routes
	router.HandleFunc("/api/lines", getAllLines).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/lines/{name}", getLineByName).Methods("GET", "OPTIONS")

	// Stop routes
	router.HandleFunc("/api/stops/{id}", getStopByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stops", getStopsByLineName).Queries("line", "{line}").Methods("GET", "OPTIONS")

	// arrivals
	router.HandleFunc("/api/arrivals/{stopID}", getStopArrivals).Methods("GET", "OPTIONS")

	return router
}
