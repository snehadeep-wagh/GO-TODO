package router

import (
	"github.com/gorilla/mux"
	"github.com/snehadeep-wagh/go-todo/middleware"
)

func TaskRoutes(r *mux.Router) {
	r.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/createTask", middleware.CreateTask).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/undotask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTaskById).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")

}
