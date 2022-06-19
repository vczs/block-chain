package block

import (
	"block-chain/common"
	"block-chain/model"
	"block-chain/utils"
	"time"
)

// CreateMetaBlock 生成创世区块
func CreateMetaBlock() *model.Block {
	block := model.Block{
		Height:    0,
		BMP:       0,
		Nonce:     0,
		PreHash:   "Creation block",
		TimeStamp: time.Now().Unix(),
		Difficult: common.Difficult,
		Data:      "创世区块",
	}
	block.Hash = utils.CalculateHash(&block) // 计算区块hash
	return &block
}
