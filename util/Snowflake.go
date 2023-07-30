package util

import (
	"fmt"
	"sync"
)

const (
	// 基准时间戳，可以根据实际需要调整
	twepoch int64 = 946656000000 // Sat Jan 01 2000 00:00:00 GMT+0800 (中国标准时间)
	// 机器ID占用的位数
	workerIDBits = 10
	// 序列号占用的位数
	sequenceBits = 12
	// 最大机器ID和最大序列号
	maxWorkerID   = -1 ^ (-1 << workerIDBits)
	maxSequenceID = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
	mu            sync.Mutex // 互斥锁
	lastTimestamp int64      // 上一次生成ID的时间戳
	workerID      int64      // 机器ID
	sequence      int64      // 序列号
}

var SF = Snowflake{}

// NewSnowflake 根据机器号号初始化Snowflake实例
func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID 超出范围")
	}
	return &Snowflake{
		lastTimestamp: -1,
		workerID:      workerID,
		sequence:      0,
	}, nil
}

// NextID 生成唯一ID
func (sf *Snowflake) NextID() int64 {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	timestamp := NowTimestamp()

	// 其他的生成ID可能已经将及时更新了 所以要等到下一个时间戳
	if timestamp < sf.lastTimestamp {
		timestamp = sf.waitNextTimestamp(sf.lastTimestamp)
	}

	// 相等时判断还有无sequence空缺 无时等到下一个时间戳生成
	if timestamp == sf.lastTimestamp {
		// 计算sequence是否到2^12了(+1之后与111111111111按位 等同于满2^12进一然后再减去那个一)
		sf.sequence = (sf.sequence + 1) & maxSequenceID
		if sf.sequence == 0 {
			timestamp = sf.waitNextTimestamp(sf.lastTimestamp)
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTimestamp = timestamp
	ID := ((timestamp - twepoch) << (workerIDBits + sequenceBits)) | (sf.workerID << sequenceBits) | sf.sequence
	return ID
}

// waitNextTimestamp 等待下一个时间戳
// 如果一秒钟内生成超过2^12个时等到下一个时间戳生成
func (sf *Snowflake) waitNextTimestamp(lastTimestamp int64) int64 {
	timestamp := NowTimestamp()
	for timestamp <= lastTimestamp {
		timestamp = NowTimestamp()
	}
	return timestamp
}
