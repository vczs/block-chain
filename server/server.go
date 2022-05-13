package server

import (
	"net/http"
)

// 服务器
func Server() {
	handler := &Handler{}
	server := http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	handler.router() // 路由
	server.ListenAndServe()
}
