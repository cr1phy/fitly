package main

import (
	"log"

	"github.com/cr1phy/fitly/internal/models"
	"github.com/cr1phy/fitly/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")

	models.InitDB()
	r := router.InitRouter()

	log.Println("server is running...")
	log.Fatalln(r.Run())
}
