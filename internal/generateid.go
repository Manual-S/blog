// Package internal 用于生成文章的唯一id
package internal

import (
	"github.com/sony/sonyflake"
)

var (
	machineID uint16
)

// 获取 机器编码ID的 回调函数
func getMachineID() (uint16, error) {
	// machineID 返回nil, 则返回专用IP地址的低16位
	return machineID, nil
}

// InitID 初始化 sonyFlake 配置
func InitID(mID uint16) (*sonyflake.Sonyflake, error) {
	machineID = mID
	st := sonyflake.Settings{}
	sFlake := sonyflake.NewSonyflake(st)
	return sFlake, nil
}
