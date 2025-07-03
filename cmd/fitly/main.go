package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cr1phy/fitly/internal/router"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func GetMigration() string {
	migration, _ := os.ReadFile("migration.sql")
	return string(migration)
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("database error:", err)
	}
	db := stdlib.OpenDBFromPool(pool)
	_, err = db.Exec(GetMigration())
	if err != nil {
		log.Fatalln("migration error:", err)
	}
	log.Println("migration finished successfully")

	server := &http.Server{Addr: ":8080", Handler: router.InitHandler(db)}
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

	log.Println("server started")
	log.Fatalln(server.ListenAndServe())
}
