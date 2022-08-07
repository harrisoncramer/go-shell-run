package handlers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true }`)
}

func RunHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ls", "-la")

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()

}
