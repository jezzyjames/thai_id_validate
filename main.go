package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please specify the HTTP port as environment variable, e.g. env PORT=8080 go run http-server.go")
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, TLS!\n")
	})
	slog.Debug("Listening and serving HTTP on :" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
