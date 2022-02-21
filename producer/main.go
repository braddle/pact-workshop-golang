package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
			Integer: 36,
			Float:   12.34,
			Boolean: true,
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hc)
	})

	log.Info("Server Stating")
	log.Panic(http.ListenAndServe(os.Getenv("HOST"), r).Error())
}
