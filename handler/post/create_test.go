package post

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// 初始化 router 和 handler
	router := gin.New()
	h := &Post{}
	router.POST("/create", h.create)

	// 準備測試 payload
	//jsonBody := []byte(`{"title":"Alice","content":"123"}`)
	jsonBody := []byte(`{"title":"Alice"}`)
	req, _ := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// 使用 httptest recorder 接收 response
	resp := httptest.NewRecorder()

	// 執行請求
	router.ServeHTTP(resp, req)

	// 驗證
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	expected := `{"error":"Key: 'PostCreate.Content' Error:Field validation for 'Content' failed on the 'required' tag"}`
	assert.JSONEq(t, expected, resp.Body.String())
}
