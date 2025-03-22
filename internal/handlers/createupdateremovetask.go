package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/andyshapirov/todolist/internal/database"
)

func (h *TaskHandler) CreateUpdateRemoveTask(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == "GET":
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

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(task)
	case r.Method == "POST" || r.Method == "PUT":
		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		var task database.Task
		if err := json.Unmarshal(body, &task); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		if len(task.Title) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "empty title"})
			return
		}

		now, _ := time.Parse("20060102", time.Now().UTC().Format("20060102"))

		nextDate := ""
		if len(task.Repeat) > 0 {
			nextDate, err = NextDate(now, now.Format("20060102"), task.Repeat)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
				return
			}
		}

		if len(task.Date) == 0 {
			task.Date = now.Format("20060102")
		}

		date, err := time.Parse("20060102", task.Date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid date format"})
			return
		}

		if len(task.Repeat) > 0 {
			if date.Before(now) {
				task.Date = nextDate
			}
		} else {
			if date.Before(now) {
				task.Date = now.Format("20060102")
			}

		}

		if r.Method == "POST" {
			id, err := h.Service.InsertOne(&task)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
				return
			}

			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(IDResponse{ID: strconv.Itoa(id)})
			return
		}

		if err := h.Service.UpdateOne(&task); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(struct{}{})
	case r.Method == "DELETE":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
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
