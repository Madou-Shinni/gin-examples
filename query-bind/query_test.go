package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldBindQuery(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/?name=Shrek", nil)

	var req SearchDTO[CharactersFilter]
	assert.Equal(t, c.ShouldBindQuery(&req), nil)

	log.Printf("name: %v\n", *req.Filter.Name)
}
