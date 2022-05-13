package block

import (
	"block-chain/common"
	"block-chain/model"
	"block-chain/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// CreationBlock 生成创世区块
func CreationBlock() *model.Block {
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

// CreateBlock 生成区块
func CreateBlock(bmp, data string) (int, *model.Block) {
	int_bmp := 0
	if bmp != "" {
		b, err := strconv.Atoi(bmp)
		if err != nil {
			log.Println("Next Get err:", err)
			return common.PARAMETER_ERROR, nil
		}
		int_bmp = b
	}
	next_block := model.Block{
		Height:    len(common.BlockChain),
		BMP:       int_bmp,
		Nonce:     common.BlockChain[len(common.BlockChain)-1].Nonce,
		PreHash:   common.BlockChain[len(common.BlockChain)-1].Hash,
		TimeStamp: time.Now().Unix(),
		Difficult: common.Difficult,
		Data:      data,
	}
	start := time.Now() // 开始挖矿时间
	log.Println("开始挖矿,当前挖矿难度:", common.Difficult)
	for {
		next_block.Nonce++                       // 每计算一次随机数自增1 挖矿的唯一变量
		hash := utils.CalculateHash(&next_block) // 计算区块hash
		if strings.HasPrefix(hash, strings.Repeat("0", next_block.Difficult)) {
			// hash值满足挖矿条件代表挖矿成功
			log.Println("->", next_block.Nonce, ":", hash, "<挖矿成功>+++++++")
			next_block.Hash = hash
			break
		}
		log.Println("->", next_block.Nonce, ":", hash, "<无效>")
	}
	use_time := time.Since(start) // 计算挖矿耗时
	log.Println("耗时:", use_time)
	fmt.Println()
	return common.SUCCESS, &next_block
}
