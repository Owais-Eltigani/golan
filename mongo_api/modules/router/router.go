package router

import (
	"mongo/modules/controllers"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movies", controllers.GetAllDbMovies).Methods("GET")
	router.HandleFunc("/movies", controllers.DeleteAll).Methods("DELETE")
	router.HandleFunc("/movies", controllers.InsertOneMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", controllers.MarkOneMovieWatched).Methods("PUT")

	return router
}
