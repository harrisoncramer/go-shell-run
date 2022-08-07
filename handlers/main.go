package handlers

import (
	"bufio"
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

func runJob(command string, argString string) {

	args := strings.Split(argString, " ")
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

func RestartHandler(w http.ResponseWriter, req *http.Request) {

	log.Println("Restarting production server...")
	runJob("docker-compose", "-f docker-compose.yml down --volumes")
	runJob("docker-compose", "-f docker-compose.yml pull")
	runJob("docker-compose", "-f docker-compose.yml up -d")
	runJob("docker", "system prune -a -f")
	log.Println("Production server rebooted!")
}
