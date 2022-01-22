package tests

import (
	"appslab-ke/kipchoge-go/pkg/routes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func InitSetup()  {
	gin.SetMode(gin.TestMode)
}

func TestJokeRoutes(t *testing.T)  {
	InitSetup()
	r := routes.Router()
	w := httptest.NewRecorder()
	t.Run("get jokes categories", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/jokes/categories", nil)
		r.ServeHTTP(w, req)

		bdy, _ := io.ReadAll(w.Result().Body)
		var f interface{}
		json.Unmarshal(bdy, &f)
		cat := f.(map[string]interface{})
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, cat["demo"], "demo")
	})
}
