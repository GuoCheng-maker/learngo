package main

import (
	"fmt"
	"time"
)

// SnowFlake 定义一个 SnowFlake 结构体
type SnowFlake struct {
	workerID      int64
	datacenterID  int64
	sequence      int64
	lastTimestamp int64
}

// 定义常量
const (
	twepoch            = int64(1288834974657) // 起始时间戳，毫秒为单位
	workerIDBits       = uint(5)              // 工作机器 ID 位数
	datacenterIDBits   = uint(5)              // 数据中心 ID 位数
	maxWorkerID        = -1 ^ (-1 << workerIDBits)
	maxDatacenterID    = -1 ^ (-1 << datacenterIDBits)
	sequenceBits       = uint(12)     // 序列号位数
	workerIDShift      = sequenceBits // 工作机器 ID 左移位数
	datacenterIDShift  = sequenceBits + workerIDBits
	timestampLeftShift = sequenceBits + workerIDBits + datacenterIDBits
	sequenceMask       = -1 ^ (-1 << sequenceBits)
)

// NewSnowFlake 函数用于创建一个 SnowFlake 实例
func NewSnowFlake(workerID int64, datacenterID int64, sequence int64) *SnowFlake {
	if workerID > maxWorkerID || workerID < 0 {
		panic(fmt.Sprintf("worker ID can't be greater than %d or less than 0", maxWorkerID))
	}
	if datacenterID > maxDatacenterID || datacenterID < 0 {
		panic(fmt.Sprintf("datacenter ID can't be greater than %d or less than 0", maxDatacenterID))
	}
	return &SnowFlake{
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      sequence,
		lastTimestamp: -1,
	}
}

// NextID 函数用于生成下一个 ID
func (sf *SnowFlake) NextID() int64 {
	timestamp := time.Now().UnixNano() / 1000000
	if timestamp < sf.lastTimestamp {
		panic(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", sf.lastTimestamp-timestamp))
	}
	if timestamp == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			timestamp = sf.tilNextMillis()
		}
	} else {
		sf.sequence = 0
	}
	sf.lastTimestamp = timestamp
	return (timestamp-twepoch)<<timestampLeftShift |
		(sf.datacenterID << datacenterIDShift) |
		(sf.workerID << workerIDShift) |
		sf.sequence
}

// tilNextMillis 函数用于等待下一毫秒
func (sf *SnowFlake) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano() / 1000000
	for timestamp <= sf.lastTimestamp {
		timestamp = time.Now().UnixNano() / 1000000
	}
	return timestamp
}

func main() {
	// 创建一个 SnowFlake 实例，传入 workerID 和 datacenterID
	sf := NewSnowFlake(1, 1, 1)
	fmt.Println(sf)
	// 循环10次 range只能循环数组、切片、字符串、map、通道
	for _ = range make([]int, 10) {
		fmt.Println(sf.NextID())
	}

	for i := range "123456" {
		fmt.Println(i)
	}

	// 死循环
	for {
		fmt.Println(sf.NextID())
	}
}
