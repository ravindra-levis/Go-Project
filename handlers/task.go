package handlers

import (
	"encoding/json"
	"example/hello/task-api/models"
	"example/hello/task-api/storage"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler struct {
	Store storage.TaskStore
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTasks(w, r)
	case http.MethodPost:
		h.CreateTask(w, r)
	default:
		http.Error(w, " Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
		return
	}
	if strings.TrimSpace(task.Title) == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Title is Required",
		})
		return
	}
	created, _ := h.Store.Create(task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, _ := h.Store.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) TaskById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid Id",
		})
		return
	}
	switch r.Method {
	case http.MethodPut:
		var payload struct {
			Done *bool `json:"done"` //Why pointer is added because to distinguish between false and “not provided”.
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.Done == nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid payload",
			})
		}
		err = h.Store.Update(id, *payload.Done)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"status": "updated",
		})
	case http.MethodDelete:
		err := h.Store.Delete(id)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{
				"error": "Bad Request",
			})
		}
		json.NewEncoder(w).Encode(map[string]string{
			"status": "deleted",
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
