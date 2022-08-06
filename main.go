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

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler)
	mux.HandleFunc("/restart", handlers.RestartHandler)

	wrappedMux := logger.NewLogger(mux)

	flag.StringVar(&port, "port", ":3012", "Port of server")
	log.Println(fmt.Sprintf("Server is starting on port %s", port))

	log.Fatal(http.ListenAndServe(port, wrappedMux))
}
