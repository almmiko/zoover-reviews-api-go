package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zoover-reviews-api-go/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := gin.New()

	router := routes.NewRouter(app)

	router = router.RegisterAll()

	server := &http.Server{
		Addr:              "localhost:8081",
		Handler:           router.App,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed  {
			log.Fatal(errors.Wrap(err, "Error starting server"))
		}
	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-shutdown

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "Error shutdown server"))
	}

	select {
	case <-ctx.Done():
		log.Println("Server closed")
	}
}
