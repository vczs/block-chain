package server

import (
	"block-chain/common"
	"block-chain/controller/block_chain"
)

// 路由
func (handler *Handler) router() {
	handler.handlerMap = make(map[string]func(*common.Context) (int, interface{}))
	blockChain(handler) // 调用区块链路由
}

// blockChain 区块链路由
func blockChain(handler *Handler) {
	handler.handlerMap["/show"] = block_chain.Show // 查看区块
	handler.handlerMap["/next"] = block_chain.Next // 添加区块
}
