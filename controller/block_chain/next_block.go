package block_chain

import (
	"block-chain/common"
	"block-chain/controller/block"
	"block-chain/utils"
	"net/http"
)

// NextBlock 添加区块处理器
func NextBlock(c *common.Context) (int, string, interface{}) {
	bmp := ""
	data := ""
	if c.R.Method == http.MethodGet {
		bmp = c.R.URL.Query().Get("bmp")
		data = c.R.URL.Query().Get("data")
	}
	if c.R.Method == http.MethodPost {
		type Req struct {
			BMP  string `json:"bmp"`
			Data string `json:"data"`
		}
		req := Req{}
		err := c.Bind(&req) // 解析请求体内容
		if err != nil {
			utils.VczsLog("Next Bind err:", err)
			return common.PARAMETER_ERROR, "", nil
		}
		bmp = req.BMP
		data = req.BMP
	}
	code, message, block := block.CreateNewBlock(bmp, data) // 生成区块
	if code != common.SUCCESS {
		return code, message, nil
	}
	block_chain, err := common.AddBlock(block) // 将生成的区块添加到区块链中
	if err != nil {
		utils.VczsLog("Next AddBlock err:", err)
		return common.ADDCHAIN_ERROR, "", nil
	}
	return common.SUCCESS, "", block_chain // 返回区块链
}
