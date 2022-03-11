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

type JsonError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Thing struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Integer       int64          `json:"integer"`
	Float         float64        `json:"float"`
	Boolean       bool           `json:"boolean"`
	SmallerThing  SmallerThing   `json:"smaller_thing"`
	Strings       []string       `json:"strings"`
	Integers      []int64        `json:"integers"`
	floats        []float64      `json:"floats"`
	SmallerThings []SmallerThing `json:"smaller_things"`
}

type SmallerThing struct {
	Title string `json:"title"`
	Age   int64  `json:"age"`
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
	r.HandleFunc("/thing/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "123456789" {
			e := JsonError{
				Code:    http.StatusNotFound,
				Status:  "NotFound",
				Message: "Could not find thing with id: 123456789",
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(e)
			return
		}

		t := Thing{
			ID:      "987654321",
			Name:    "Testing",
			Integer: 1357,
			Float:   4.56,
			Boolean: false,
			SmallerThing: SmallerThing{
				Title: "",
				Age:   0,
			},
			Strings:  []string{"one", "two", "three"},
			Integers: []int64{1, 2, 3, 4},
			floats:   []float64{1.23, 3.45, 5.67},
			SmallerThings: []SmallerThing{
				{
					Title: "First",
					Age:   3,
				},
				{
					Title: "Second",
					Age:   2,
				},
				{
					Title: "Third",
					Age:   1,
				},
			},
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(t)
	})

	log.Info("Server Stating")
	log.Panic(http.ListenAndServe(os.Getenv("HOST"), r).Error())
}
