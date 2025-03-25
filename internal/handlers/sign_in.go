package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *TaskHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	var passwordRequest PasswordRequest
	if err := json.Unmarshal(body, &passwordRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	pass := h.Password
	if passwordRequest.Password != h.Password {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "wrong password"})
		return
	}

	secret := h.Secret
	token, err := CreateJWT(pass, secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(TokenRequest{Token: token})
}
