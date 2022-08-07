package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true }`)
}

func runJob(command string, args []string) {

	cmd := exec.Command(command, args...)

	stdout, _ := cmd.StdoutPipe()

	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	outputScanner := bufio.NewScanner(stdout)
	outputScanner.Split(bufio.ScanLines)
	for outputScanner.Scan() {
		m := outputScanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()
}

type Jobs struct {
	Jobs []string
}

func RestartHandler(w http.ResponseWriter, req *http.Request) {

	var jobs Jobs

	err := json.NewDecoder(req.Body).Decode(&jobs)
	if err != nil {
		http.Error(w, "Must provide jobs.", http.StatusUnauthorized)
		return
	}

	log.Println(jobs.Jobs)
	for i := 0; i < len(jobs.Jobs); i++ {
		values := strings.Split(jobs.Jobs[i], " ")
		runJob(values[0], values[1:])
	}
}
