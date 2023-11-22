package web

import (
	"encoding/json"
	usecase "github.com/danyukod/wallet-core-go/internal/usecase/create_client"
	"net/http"
)

type ClientHandler struct {
	CreateClientUseCase usecase.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase usecase.CreateClientUseCase) *ClientHandler {
	return &ClientHandler{CreateClientUseCase: createClientUseCase}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.CreateClientUseCase.Execute(dto)
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
