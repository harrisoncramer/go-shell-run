package handlers

import (
	"io"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true }`)
}

func RestartHandler(w http.ResponseWriter, req *http.Request) {
}
