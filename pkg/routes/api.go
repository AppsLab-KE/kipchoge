package routes

import (
	"appslab-ke/kipchoge-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//Apis
	jokes := r.Group("/jokes")

		jokes.GET("/categories", handlers.GetJokeCategories)

	return r
}
