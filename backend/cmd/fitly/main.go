package main

import (
	"log"

	"github.com/cr1phy/fitly/internal/models"
	"github.com/cr1phy/fitly/internal/router"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	ctx := context.Background()

// 	server := &http.Server{Addr: ":8080", Handler: router.InitHandler(db)}
// 	serverCtx, serverCancel := context.WithCancel(ctx)

// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

// 	go func() {
// 		<-sig
// 		shutdownCtx, shutdownCancel := context.WithTimeout(serverCtx, 30*time.Second)
// 		defer shutdownCancel()

// 		go func() {
// 			<-shutdownCtx.Done()
// 			if shutdownCtx.Err() == context.DeadlineExceeded {
// 				log.Fatalln("graceful shutdown timed out.")
// 			}
// 		}()

// 		err := server.Shutdown(serverCtx)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		serverCancel()
// 	}()

// 	log.Println("server started")
// 	log.Fatalln(server.ListenAndServe())
// }

func main() {
	gin.SetMode("release")

	models.InitDB()
	r := router.InitRouter()

	log.Println("server is running...")
	log.Fatalln(r.Run())
}
