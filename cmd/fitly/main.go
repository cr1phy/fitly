package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cr1phy/fitly/internal/database"
	"github.com/cr1phy/fitly/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func InitRouter(db *sql.DB) chi.Router {
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
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		filter := r.URL.Query().Get("filter")

		w.Header().Set("Content-Type", "application/json")

		if filter == "" {
			http.Error(w, "Filter is empty", http.StatusBadRequest)
			return
		}

		products, err := database.GetAllProductsFromFilter(db, filter)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		message, _ := json.Marshal(map[string]any{"products": products})
		w.Write(message)
	})
	router.Post("/addProduct", func(w http.ResponseWriter, r *http.Request) {
		body, err := r.GetBody()
		if err != nil {
			log.Println(err)
			return
		}
		defer body.Close()

		var data []byte
		_, err = body.Read(data)
		if err != nil {
			log.Println(err)
			return
		}

		var product entity.Product
		if err := json.Unmarshal(data, &product); err != nil {
			log.Println(err)
			return
		}
		w.Write(data)
	})

	return router
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("database error:", err)
	}
	db := stdlib.OpenDBFromPool(pool)

	server := &http.Server{Addr: ":8080", Handler: InitRouter(db)}
	serverCtx, serverCancel := context.WithCancel(ctx)

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
