package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

const LIMIT = 10

func (h *TaskHandler) GetUpcommingTasks(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	date, err := time.Parse("02.01.2006", search)
	if err == nil {
		search = date.Format(LAYOUT)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	tasks, err := h.Service.GetTasks(search, LIMIT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)

}
