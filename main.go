package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jezzyjames/thai_id_validate/programming"
	"github.com/jezzyjames/thai_id_validate/thai_id"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please specify the HTTP port as environment variable, e.g. env PORT=8080 go run http-server.go")
	}

	// DB
	db, err := sql.Open("sqlite3", "./languages.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	idleConnsClosed := make(chan struct{})

	router := gin.Default()
	router.Use(SimpleMiddleware)
	languageHandler := programming.NewLanguageHandler(db)
	router.GET("/languages", gin.WrapH(languageHandler))

	thaiIdHandler := thai_id.NewThaiIDHandler(db)
	router.POST("/thai/ids/verify", thaiIdHandler.ThaiIdValidateHandler)

	// Mux version
	// mux := http.NewServeMux()
	// languageHandler2 := programming.NewLanguageHandler(db)
	// mux.HandleFunc("/hello", languageHandler2.ServeHTTP)

	srv := http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           router,
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

func SimpleMiddleware(c *gin.Context) {
	// if c.Request.Header.Get("RequestID") == "" {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	c.Next()
	log.Print("Success")
}
