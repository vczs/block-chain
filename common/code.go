package common

const (
	SUCCESS = 0 // success

	NOT_FOUND = 404 // NotFound

	PARAMETER_ERROR = 1001 // 参数错误
	ADDCHAIN_ERROR  = 2001 // 区块加入区块链失败
	NOT_FOUND_BLOCK = 2002 // 未查找到目标区块
)

var message = map[int]string{
	SUCCESS: "success",

	NOT_FOUND: "NotFound",

	PARAMETER_ERROR: "参数错误",
	ADDCHAIN_ERROR:  "区块加入区块链失败",
	NOT_FOUND_BLOCK: "未查找到目标区块",
}

// GetMessage 获取message
func GetMessage(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务器发生未知错误~"
	}
}
