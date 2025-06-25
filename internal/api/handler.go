package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"appsconcept/internal/domain"
	"appsconcept/internal/service"
)

type Handler struct {
	service *service.FizzBuzzService
}

func NewHandler(svc *service.FizzBuzzService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) FizzBuzz(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	q := r.URL.Query()
	int1, err1 := strconv.Atoi(q.Get("int1"))
	int2, err2 := strconv.Atoi(q.Get("int2"))
	limit, err3 := strconv.Atoi(q.Get("limit"))
	str1 := q.Get("str1")
	str2 := q.Get("str2")

	if err1 != nil || err2 != nil || err3 != nil || str1 == "" || str2 == "" {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	params := domain.FizzBuzzParams{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	result, err := h.service.GenerateFizzBuzz(ctx, params)
	if err != nil {
		errorMessage := fmt.Sprintf("Service error: %+v", err)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	stats, err := h.service.GetStats(ctx)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Stats unavailable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
