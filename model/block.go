package model

// 区块结构
type Block struct {
	Height    int    `json:"height"`     // 区块链高度
	BMP       int    `json:"bmp"`        // 交易信息
	Nonce     int    `json:"nonce"`      // 随机数
	Hash      string `json:"hash"`       // 当前区块随机值
	PreHash   string `json:"pre_hash"`   // 上一区块随机值
	TimeStamp int64  `json:"time_stamp"` // 时间戳
	Difficult int    `json:"difficult"`  // 挖矿难度
	Data      string `json:"data"`       // 区块信息
}
