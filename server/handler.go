package server

import (
	"block-chain/common"
	"block-chain/model"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	handlerMap map[string]func(*common.Context) (int, interface{}) // 路由对应处理器函数
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.Proto, r.Host, r.URL.Path) // 打印接口调用日志
	if router, ok := handler.handlerMap[r.URL.Path]; ok {
		code, result := router(&common.Context{W: w, R: r}) // 调用路由并传入Context
		if code == common.NOT_FOUND {
			// 如果处理器返回code码为404 就返回404页面
			notFound(w)
			return
		}
		message := common.GetMessage(code) // 获取message
		res, err := json.Marshal(model.Result{Code: code, Message: message, Result: result})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("服务器异常:" + err.Error()))
		} else {

			w.Header().Set("content-type", "text/json") // 编码格式设置为json类型
			w.Write(res)                                // 返回结果
		}
	} else {
		notFound(w) // 404页面
	}
}

// notFound 返回404页面
func notFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("<h1 align=\"center\">404 Not Found</h1>"))
}
