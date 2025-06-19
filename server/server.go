package server

import (
	"Todo/handlers"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(struct{ Message string }{Message: "server is running"})
		if err != nil {
			logrus.Errorf("Error encoding response: %v", err)
		}
	}).Methods("GET")

	// User Auth Routes
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutUser).Methods("POST")

	// Profile
	r.HandleFunc("/profile", handlers.GetUserProfile).Methods("GET")

	// Task Routes (Protected)
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/update", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/update/status", handlers.UpdateTaskStatus).Methods("PATCH")
	//r.HandleFunc("/tasks", handlers.DeleteTask).Methods("DELETE")

	return r
}
