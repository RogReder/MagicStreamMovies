package middleware

import (
	"net/http"

	"github.com/RogReder/MagicStreamMovies/Server/MagicStreamMoviesServer/utils"
	"github.com/gin-gonic/gin"
	"fmt"
)

func AuthMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context){
		token, err := utils.GetAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return 
		}

		if token == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"No token is provided"})
			c.Abort()
			return 
		}

		claims, err := utils.ValidationToken(token)

		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token"})
			c.Abort()
			return 
		}

		fmt.Println("ROLE", claims)
		c.Set("userId", claims.UserId)
		c.Set("role", claims.Role)

		c.Next()
	}
}