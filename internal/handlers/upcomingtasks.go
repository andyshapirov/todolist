package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *TaskHandler) UpcommingTasks(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == "GET":
		search := r.FormValue("search")
		date, err := time.Parse("02.01.2006", search)
		if err == nil {
			search = date.Format("20060102")
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		limit := 10
		tasks, err := h.Service.SelectN(search, limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(tasks)
	}

}
