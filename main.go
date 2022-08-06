package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/harrisoncramer/golang-webhook/handlers"
	"github.com/harrisoncramer/golang-webhook/logger"
)

var port string

func main() {

	// mux := http.NewServeMux()
	// mux.Handle("/status", handlers.StatusHandler)

	http.HandleFunc("/restart", handlers.RestartHandler)
	http.HandleFunc("/status", handlers.StatusHandler)

	flag.StringVar(&port, "port", ":3012", "Port of server")
	logger.Log(fmt.Sprintf("Server is starting on port %s", port))
	log.Fatal(http.ListenAndServe(port, nil))
}
