package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}
	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(out)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "highlander",
		ReleaseDate: rd,
		MPAARating:  "G",
		RunTime:     116,
		Description: "良い映画",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	rd, _ = time.Parse("2006-01-02", "1981-03-07")

	riders := models.Movie{
		ID:          2,
		Title:       "ジュラシックパーク",
		ReleaseDate: rd,
		MPAARating:  "PG-12",
		RunTime:     115,
		Description: "迫力満点",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, riders)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(out)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
