package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/crissmcoelho/estoque/internal/repository"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct{ repo *repository.ProductRepo }

func NewProductHandler(r *repository.ProductRepo) *ProductHandler { return &ProductHandler{repo: r} }

type AdjustRequest struct {
	RequestID string `json:"request_id"` // idempotency key
	Delta     int    `json:"delta"`      // negative to subtract
}

func (h *ProductHandler) AdjustProduct(w http.ResponseWriter, r *http.Request) {
	codigo := chi.URLParam(r, "codigo")
	var req AdjustRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := h.repo.AdjustByCodigoIdempotent(r.Context(), codigo, req.Delta, req.RequestID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
