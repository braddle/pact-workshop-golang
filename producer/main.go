package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
)

type HealthCheck struct {
	Status  string  `json:"Status"`
	Message string  `json:"message"`
	Integer int64   `json:"integer"`
	Float   float64 `json:"float"`
	Boolean bool    `json:"boolean"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		hc := HealthCheck{
			Status:  "OK",
			Message: "Testing Testing 123",
			Integer: 42,
			Float:   math.Pi,
			Boolean: true,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hc)
	})

	log.Info("Server Stating")
	log.Panic(http.ListenAndServe(os.Getenv("HOST"), r).Error())
}
