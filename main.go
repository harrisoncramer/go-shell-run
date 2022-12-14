package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/harrisoncramer/go-shell-run/handlers"
	"github.com/harrisoncramer/go-shell-run/middleware"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.StatusHandler)
	mux.HandleFunc("/jobs", handlers.JobHandler)

	wrappedMux := middleware.CheckToken(middleware.NewLogger(mux))

	port := flag.String("port", "3012", "Port of server")
	token := flag.String("token", "secret", "Secret required to run /jobs and other routes")

	flag.Parse()

	os.Setenv("SECRET_TOKEN", *token)

	log.Println(fmt.Sprintf("Server is starting on port %s", *port))

	log.Fatal(http.ListenAndServe(":"+*port, wrappedMux))
}
