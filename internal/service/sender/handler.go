package sender

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NRKA/home-service/internal/usecase"
	"github.com/NRKA/home-service/pkg/postgres"
	"net/http"
	"strconv"
)

type Subscriber interface {
	Subscribe(ctx context.Context, houseID int, email usecase.Subscribe) error
}

type Handler struct {
	repo Subscriber
}

func NewHandler(db *postgres.Database) *Handler {
	return &Handler{NewRepo(db)}
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, ErrInvalidID.Error(), http.StatusBadRequest)
		return
	}

	var email usecase.Subscribe
	if err = json.NewDecoder(r.Body).Decode(&email); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err = email.Validate(); err != nil {
		http.Error(w, ErrInvalidEmail.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.Subscribe(r.Context(), id, email)
	if err != nil {
		if errors.Is(err, ErrHouseNotFound) {
			http.Error(w, ErrHouseNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to subscribe", http.StatusInternalServerError)
		return
	}
}
