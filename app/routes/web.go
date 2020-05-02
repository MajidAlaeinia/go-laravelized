package routes

import (
	"github.com/go-laravelized/app/controllers"
	"github.com/gorilla/mux"
)

func WebRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies/{category}/{id:[0-9]}", controllers.MovieHandler).Name("movies.show")
	router.HandleFunc("/articles/{category}/{id:[0-9]}", controllers.ArticleHandler).Name("articles.show")

	return router
}
