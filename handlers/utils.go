package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/harrisoncramer/golang-webhook/state"
)

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

func RunJobs(jobs []string) {
	for i := 0; i < len(jobs); i++ {
		values := strings.Split(jobs[i], " ")
		runJob(values[0], values[1:])
	}

	state.Processing = false

}
