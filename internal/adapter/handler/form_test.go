package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func testFormHandler(t *testing.T) {
	router := SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/forms", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
