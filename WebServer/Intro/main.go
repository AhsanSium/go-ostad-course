package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func initLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

func handleMiddleware(next http.Handler) http.Handler {
	// Logic for middleware can be added here, such as logging, authentication, etc.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestId := time.Now().UnixNano()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-ID", fmt.Sprintf("%d", requestId))
		slog.Info("Incoming Request", 
			"Method", r.Method,
			"Path", r.URL.Path,
			"RequestID", requestId,
			"RemoteAddr", r.RemoteAddr,
	  )
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		slog.Info("Request Completed:", 
			"Duration", duration.Microseconds(),
			"RequestID", requestId,
		)
	})
}

func jsonHandler(w http.ResponseWriter,  r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	response := map [string]string{
		"title": "Hello from Go",
		"message": "This is a JSON response",
	}
	json.NewEncoder(w).Encode(response)
}

func main () {
	initLogger()
	http.Handle("/", handleMiddleware(http.HandlerFunc(jsonHandler)))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, From Go, You Requested: %s\n", r.URL.Path)
	})
	fmt.Println("Starting Server at 8010")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
