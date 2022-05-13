package common

import (
	"block-chain/model"
	"errors"
	"sync"
)

var BlockChain []*model.Block // 区块链

var mutex sync.Mutex

// AddBlock 添加区块到区块链
func AddBlock(block *model.Block) ([]*model.Block, error) {
	if len(BlockChain) > 0 && block.PreHash != BlockChain[len(BlockChain)-1].Hash {
		// 在创世区块已生成的情况下，如果当前要添加的区块的PreHash值不等于区块链中最后一个区块的hash值就返回err
		return nil, errors.New("this block preHash not equal to previous block hash")
	}
	mutex.Lock()                           // 加锁
	BlockChain = append(BlockChain, block) // 添加该区块到区块链中
	mutex.Unlock()                         // 释放锁
	return BlockChain, nil
}
