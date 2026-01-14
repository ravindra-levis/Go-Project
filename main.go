package main

import (
	"example/hello/task-api/handlers"
	"example/hello/task-api/storage"
	"net/http"
)

func main() {
	store := storage.NewMemoryStore()
	handlers := &handlers.TaskHandler{Store: store}

	http.HandleFunc("/tasks", handlers.Tasks)
	http.HandleFunc("/tasks/", handlers.TaskById)

	http.ListenAndServe(":8080", nil)
}
