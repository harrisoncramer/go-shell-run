package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/harrisoncramer/golang-webhook/handlers"
	"github.com/harrisoncramer/golang-webhook/logger"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler)
	mux.HandleFunc("/run", handlers.RunHandler)

	wrappedMux := logger.NewLogger(mux)

	port := flag.String("port", "3012", "Port of server")

	flag.Parse()

	log.Println("port", *port)

	log.Println(fmt.Sprintf("Server is starting on port %s", *port))

	log.Fatal(http.ListenAndServe(":"+*port, wrappedMux))
}
