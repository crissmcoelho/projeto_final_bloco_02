package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/your/estoque/internal/handlers"
	"github.com/your/estoque/internal/repository"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)

	repo := repository.NewProductRepo(db)
	h := handlers.NewProductHandler(repo)

	r := chi.NewRouter()
	r.Post("/products", h.CreateProduct)
	r.Get("/products", h.ListProducts)
	r.Get("/products/{codigo}", h.GetByCodigo)
	r.Post("/products/{codigo}/adjust", h.AdjustProduct) // adjust with idempotency

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Estoque listening on :8081")
	log.Fatal(srv.ListenAndServe())
}
