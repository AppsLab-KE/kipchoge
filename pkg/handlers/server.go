package handlers

import (
	"appslab-ke/kipchoge-go/internal/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db    storage.Storage
	cache *storage.Cache
}

func NewApp() *gin.Engine {
	r := gin.Default()
	//Apis
	jokes := r.Group("/jokes")
	jokes.GET("/categories", GetJokeCategories)

	return r
}
