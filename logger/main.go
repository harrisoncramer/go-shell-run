package logger

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "http: ", log.LstdFlags)

func Log(msg string) {
	logger.Println(msg)
}
