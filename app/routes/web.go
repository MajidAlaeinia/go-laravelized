package routes

import (
	"github.com/gorilla/mux"
	"github.com/playground/app/controllers"
)

func WebRoutes() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/movies/{category}/{id:[0-9]}", controllers.MovieHandler)
	router.HandleFunc("/articles/{category}/{id:[0-9]}", controllers.ArticleHandler)

	return router
}
