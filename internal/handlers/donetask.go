package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (h *TaskHandler) DoneTask(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == "POST":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		task, err := h.Service.SelectOne(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		if task == nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "internal server error"})
			return
		}

		if len(task.Repeat) > 0 {
			now := time.Now().UTC()
			nextDate, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
				return
			}

			if err := h.Service.UpdateDateOne(id, nextDate); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
				return
			}

			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(struct{}{})
			return
		}

		if err := h.Service.DeleteOne(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(struct{}{})

	}
}
