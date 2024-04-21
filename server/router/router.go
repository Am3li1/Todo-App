package router

import(
	"project/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("api/undoTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("api/deleteTask/{id}", middleware.deleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}