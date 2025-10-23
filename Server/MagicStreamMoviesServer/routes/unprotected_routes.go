package routes

import(
	controller "github.com/RogReder/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpUnProtectedRoutes(router *gin.Engine){
	

	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
}