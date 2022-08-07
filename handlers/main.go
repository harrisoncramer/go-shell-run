package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true }`)
}

func RunHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ls")

	/* Send output to buffer */
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(out.String())
}
