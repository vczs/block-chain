package block_chain

import (
	"block-chain/common"
	"block-chain/controller/block"
	"log"
)

// 生成创世区块
func init() {
	if len(common.BlockChain) < 1 {
		block := block.CreateMetaBlock() // 生成创世区块
		_, err := common.AddBlock(block) // 将创世区块添加到区块链中
		if err != nil {
			log.Fatalln("创世区块生成失败:", err)
		}
		log.Println("创世区块已生成")
	}
}
