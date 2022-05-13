package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// 解析请求体内容
func (c *Context) Bind(model interface{}) error {
	if reflect.TypeOf(model).Kind() != reflect.Pointer {
		// 如果参数不是结构体类型就返回err
		return fmt.Errorf("%v not kind struct pointer", model)
	}
	length := c.R.ContentLength
	if length < 1 {
		return nil
	}
	body := make([]byte, length)
	_, err := c.R.Body.Read(body) // 读取请求体内容到byte切片
	if err != nil && err.Error() != "EOF" {
		return err
	}
	err = json.Unmarshal(body, model) // 反序列化读取到的内容到结构体model
	return err
}

// 获取请求体中的某个参数的值
func (c *Context) Get(name string) (string, error) {
	v := ""
	if name == "" {
		return v, errors.New("name is null")
	}
	res_body := make([]byte, c.R.ContentLength)
	_, err := c.R.Body.Read(res_body) // 读取请求体内容到byte切片
	if err != nil && err.Error() != "EOF" {
		return v, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(res_body, &m) // 反序列化读取到的内容到map
	if err != nil {
		return v, err
	}
	for key, value := range m {
		// 遍历m找到目标参数
		if key == name {
			v = value.(string) // 将目标参数的值赋值给v
			break
		}
	}
	return v, nil
}
