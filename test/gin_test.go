package test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"rmall/router"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := router.PingRouter()
	// 创建一个响应记录器
	w := httptest.NewRecorder()
	// 创建一个请求
	request := httptest.NewRequest("GET", "/ping", nil)
	// 请求
	r.ServeHTTP(w, request)
	fmt.Println(w)
	// 断言
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
