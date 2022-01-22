package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJokeCategories(c *gin.Context)  {
	//s := models.JokeCategory{}
	///
	c.JSON(http.StatusOK, map[string]string{"demo": "demo"})
}