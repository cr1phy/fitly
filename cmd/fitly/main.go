package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func service() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		message, _ := json.Marshal(map[string]any{"status": "OK"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(message)
	})

	return router
}

func main() {
	server := &http.Server{Addr: ":8080", Handler: service()}
	serverCtx, serverCancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		shutdownCtx, shutdownCancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer shutdownCancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatalln("graceful shutdown timed out.")
			}
		}()

		err := server.Shutdown(serverCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverCancel()
	}()

	log.Fatalln(server.ListenAndServe())
}
