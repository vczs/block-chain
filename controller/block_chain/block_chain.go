package block_chain

import (
	"block-chain/common"
	"block-chain/controller/block"
	"log"
)

// 生成创世区块
func init() {
	if len(common.BlockChain) < 1 {
		block := block.CreationBlock()   // 生成创世区块
		_, err := common.AddBlock(block) // 将创世区块添加到区块链中
		if err != nil {
			log.Fatalln("创世区块生成失败:", err)
		}
		log.Println("创世区块已生成")
	}
}

// Show 查看区块处理器
func Show(c *common.Context) (int, interface{}) {
	data := ""
	// 根据请求类型获取data参数的值
	if c.R.Method == "GET" {
		data = c.R.URL.Query().Get("data")
	}
	if c.R.Method == "POST" {
		if c.R.ContentLength > 0 {
			get_data, err := c.Get("data")
			if err != nil {
				log.Println("Next Get err:", err)
				return common.PARAMETER_ERROR, nil
			}
			data = get_data
		}
	}
	if data != "" {
		// 查找目标区块并返回
		for _, value := range common.BlockChain {
			if value.Data == data {
				return common.SUCCESS, value // 返回目标区块
			}
		}
		return common.NOT_FOUND_BLOCK, nil
	}
	// 如果data值为空就返回所有区块(即返回区块链)
	return common.SUCCESS, common.BlockChain
}

// Next 添加区块处理器
func Next(c *common.Context) (int, interface{}) {
	bmp := ""
	data := ""
	if c.R.Method == "GET" {
		bmp = c.R.URL.Query().Get("bmp")
		data = c.R.URL.Query().Get("data")
	}
	if c.R.Method == "POST" {
		type Req struct {
			BMP  string `json:"bmp"`
			Data string `json:"data"`
		}
		req := Req{}
		err := c.Bind(&req) // 解析请求体内容
		if err != nil {
			log.Println("Next Bind err:", err)
			return common.PARAMETER_ERROR, nil
		}
		bmp = req.BMP
		data = req.BMP
	}
	code, block := block.CreateBlock(bmp, data) // 生成区块
	if code != common.SUCCESS {
		return code, nil
	}
	block_chain, err := common.AddBlock(block) // 将生成的区块添加到区块链中
	if err != nil {
		log.Println("Next AddBlock err:", err)
		return common.ADDCHAIN_ERROR, nil
	}
	return common.SUCCESS, block_chain // 返回区块链
}
