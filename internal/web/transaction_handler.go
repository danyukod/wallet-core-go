package web

import (
	"encoding/json"
	usecase "github.com/danyukod/wallet-core-go/internal/usecase/create_transaction"
	"net/http"
)

type WebTransactionHandler struct {
	usecase.CreateTransactionUseCase
}

func NewWebTransactionHandler(createTransactionUseCase usecase.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{CreateTransactionUseCase: createTransactionUseCase}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateTransactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	output, err := h.CreateTransactionUseCase.Execute(ctx, &dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
