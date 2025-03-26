package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andyshapirov/todolist/internal/database"
)

func (h *TaskHandler) GetUpcommingTasks(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	date, err := time.Parse("02.01.2006", search)
	if err == nil {
		search = date.Format(Layout)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	taskList, err := h.Service.GetTasks(search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	tasks := map[string]*[]database.Task{
		"tasks": taskList,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)

}
