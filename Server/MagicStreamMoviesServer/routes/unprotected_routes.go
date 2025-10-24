package routes

import(
	controller "github.com/RogReder/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetUpUnProtectedRoutes(router *gin.Engine, client*mongo.Client){
	

	router.GET("/movies", controller.GetMovies(client))
	router.POST("/register", controller.RegisterUser(client))
	router.POST("/login", controller.LoginUser(client))
}