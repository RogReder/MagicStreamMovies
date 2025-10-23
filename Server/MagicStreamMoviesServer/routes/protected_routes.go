package routes

import (
	controller "github.com/RogReder/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/RogReder/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.GET("/recommendedmovies", controller.GetRecommendedMovies())
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate())
}
