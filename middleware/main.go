package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

var printer = log.New(os.Stdout, "http: ", log.LstdFlags)

type Logger struct {
	handler http.Handler
}

/* Define a ServeHTTP method on the Middleware struct so this method is called
first, which then calls the ServeHTTP method on the wrapped struct (l.handler.ServeHTTP) */
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	printer.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

type TokenChecker struct {
	handler    http.Handler
	headerName string
}

func (tc *TokenChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get(tc.headerName)
	if header != os.Getenv("SECRET_TOKEN") {
		http.Error(w, "That token is invalid.", http.StatusUnauthorized)
		return
	}

	tc.handler.ServeHTTP(w, r)
}

func CheckToken(handlerToWrap http.Handler) *TokenChecker {
	return &TokenChecker{handlerToWrap, "token"}
}
