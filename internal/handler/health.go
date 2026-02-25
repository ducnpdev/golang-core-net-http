package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

var startTime = time.Now()

type HealthResponse struct {
	Status     string        `json:"status"`
	Uptime     time.Duration `json:"uptime"`
	Goroutines int           `json:"goroutines"`
	Timestamp  time.Time     `json:"timestamp"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:     "ok",
		Uptime:     time.Since(startTime),
		Goroutines: runtime.NumGoroutine(),
		Timestamp:  time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
