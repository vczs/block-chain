package block_chain

import (
	"block-chain/common"
	"block-chain/utils"
	"net/http"
)

// ShowChain 查看区块处理器
func ShowChain(c *common.Context) (int, string, interface{}) {
	data := ""
	// 根据请求类型获取data参数的值
	if c.R.Method == http.MethodGet {
		data = c.R.URL.Query().Get("data")
	}
	if c.R.Method == http.MethodPost {
		if c.R.ContentLength > 0 {
			get_data, err := c.Get("data")
			if err != nil {
				utils.VczsLog("Next Get err:", err)
				return common.PARAMETER_ERROR, "", nil
			}
			data = get_data
		}
	}
	if data != "" {
		// 查找目标区块并返回
		for _, value := range common.BlockChain {
			if value.Data == data {
				return common.SUCCESS, "", value // 返回目标区块
			}
		}
		return common.NOT_FOUND_BLOCK, "", nil
	}
	// 如果data值为空就返回所有区块(即返回区块链)
	return common.SUCCESS, "", common.BlockChain
}
