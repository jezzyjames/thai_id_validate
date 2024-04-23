package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please specify the HTTP port as environment variable, e.g. env PORT=8080 go run http-server.go")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	idleConnsClosed := make(chan struct{})
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Hi\n")
		if err != nil {
			slog.Error(err.Error())
		}
	})

	srv := http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		<-ctx.Done()
		fmt.Println("shutting down...")

		context, contextCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer contextCancel()

		if err := srv.Shutdown(context); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	slog.Debug("Listening and serving HTTP on :" + port)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Panicf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	fmt.Println("bye")
}

// ps -ef|grep go, see process id
// kill -SIGTERM {processID}
// kill -SIGINT {processID}
