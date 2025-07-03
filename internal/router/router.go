package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cr1phy/fitly/internal/database"
	"github.com/cr1phy/fitly/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func status(w http.ResponseWriter, r *http.Request) {
	message, _ := json.Marshal(map[string]any{"status": "OK"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func getProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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
}

func addProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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
	id, err := database.CreateProduct(db, product)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte{byte(id)})
}

func InitHandler(db *sql.DB) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Get("/", status)
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		getProducts(db, w, r)
	})
	router.Post("/addProduct", func(w http.ResponseWriter, r *http.Request) {
		addProduct(db, w, r)
	})

	return router
}
