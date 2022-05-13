package utils

import (
	"block-chain/model"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// CalculateHash 计算区块hash值 sha256算法
func CalculateHash(block *model.Block) string {
	data := strconv.Itoa(block.Height) + strconv.FormatInt(block.TimeStamp, 10) + block.PreHash + block.Data + strconv.Itoa(block.Nonce)
	blockInBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(blockInBytes[:])
}
