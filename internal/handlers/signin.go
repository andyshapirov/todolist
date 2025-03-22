package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/andyshapirov/todolist/tests"
)

func (h *TaskHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == "POST":
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

		pass := tests.Password
		if v, ok := os.LookupEnv("TODO_PASSWORD"); len(v) > 0 && ok {
			pass = v
		}

		if passwordRequest.Password != pass {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "wrong password"})
			return
		}

		token, err := CreateJWT(pass)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(TokenRequest{Token: token})
	}
}
