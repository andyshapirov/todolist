package handlers

import (
	"net/http"
	"time"
)

func (h *TaskHandler) TaskDate(w http.ResponseWriter, r *http.Request) {
	now, err := time.Parse("20060102", r.FormValue("now"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	date, err := NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(date))
}
