package library

/**
 *
 *
 * @author        Jessie Gui <guijiaxian@gmail.com>
 * @version       1.0.0
 * @copyright (c) 2022, Jessie Gui
 */
import (
	"sync"
	"time"
)

const (
	// 时间起点
	epoch int64 = 1609459200000 // 2021-01-01 00:00:00 UTC
	// 机器标识位数
	workerBits uint8 = 10
	// 序列号位数
	sequenceBits uint8 = 12
)

// Snowflake 雪花算法对象。
type Snowflake struct {
	mu            sync.Mutex // 互斥锁，保证并发安全
	lastTimestamp int64      // 上次生成ID的时间戳
	workerId      uint16     // 机器标识
	sequence      uint16     // 当前序列号
}

func NewSnowflake(workerId uint16) *Snowflake {
	if workerId < 0 || workerId >= (1<<workerBits) {
		panic("worker ID out of range")
	}
	return &Snowflake{
		lastTimestamp: 0,
		workerId:      workerId,
		sequence:      0,
	}
}

func (s *Snowflake) NextId() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 获取当前时间戳。
	now := time.Now().UnixNano() / int64(time.Millisecond)

	// 如果当前时间小于上次生成ID的时间戳，则说明时钟回拨了。
	if now < s.lastTimestamp {
		panic("clock moved backwards")
	}

	// 如果当前时间等于上次生成ID的时间戳，则需要生成序列号。
	if now == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & ((1 << sequenceBits) - 1)
		if s.sequence == 0 {
			// 序列号已经用完，等待下一毫秒。
			for now <= s.lastTimestamp {
				now = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		// 当前时间大于上次生成ID的时间戳，序列号从0开始重新计数。
		s.sequence = 0
	}

	// 保存当前时间戳。
	s.lastTimestamp = now

	// 组装ID并返回。
	return uint64((now-epoch)<<workerBits | int64(s.workerId)<<sequenceBits | int64(s.sequence))
}
