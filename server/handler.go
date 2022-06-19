package server

import (
	"block-chain/common"
	"block-chain/model"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	handlerMap map[string]func(*common.Context) (int, string, interface{}) // 路由对应处理器函数
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 打印api调用日志
	log.Println(r.Method, r.Proto, r.Host, r.URL.Path)
	if router, ok := handler.handlerMap[r.URL.Path]; ok {
		// 调用路由并传入Context
		code, message, result := router(&common.Context{W: w, R: r})
		// 如果处理器返回code码为404 就返回404页面
		if code == common.NOT_FOUND {
			notFound(w)
			return
		}
		// 如果message为空就获取定义好的message
		if message == "" {
			message = common.GetMessage(code)
		}
		res, err := json.Marshal(model.Result{Code: code, Message: message, Result: result})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("服务器数据序列化异常:" + err.Error()))
		} else {
			// 编码格式设置为json类型
			w.Header().Set("content-type", "text/json")
			w.Write(res)
		}
	} else {
		notFound(w)
	}
}

// notFound 返回404页面
func notFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("<h1 align=\"center\">404 Not Found</h1>"))
}
