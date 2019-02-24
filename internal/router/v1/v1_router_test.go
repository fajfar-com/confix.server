package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := gin.Default()

	GetRouters(router)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}