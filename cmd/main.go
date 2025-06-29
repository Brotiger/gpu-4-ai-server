package main

import (
	"encoding/json"
	"net/http"
)

type TaskRequest struct {
	Input string `json:"input"`
}

type TaskResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	// TODO: отправить задачу воркеру через gRPC
	resp := TaskResponse{Output: "stub output"}
	json.NewEncoder(w).Encode(resp)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	// TODO: собрать статусы воркеров
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	http.HandleFunc("/task", handleTask)
	http.HandleFunc("/health", handleHealth)
	http.ListenAndServe(":8080", nil)
} 