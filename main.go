package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/harrisoncramer/golang-webhook/handlers"
	"github.com/harrisoncramer/golang-webhook/middleware"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler)
	mux.HandleFunc("/restart", handlers.RestartHandler)

	wrappedMux := middleware.CheckToken(middleware.NewLogger(mux))

	port := flag.String("port", "3012", "Port of server")
	token := flag.String("token", "secret", "Secret required to run /restart and other routes")

	flag.Parse()

	os.Setenv("SECRET_TOKEN", *token)

	log.Println(fmt.Sprintf("Server is starting on port %s", *port))

	log.Fatal(http.ListenAndServe(":"+*port, wrappedMux))
}
