package snow_id

import (
	"sync"
	"time"
)

//雪花ID生成
const (
	workerBit uint8 = 10
	nuberBit  uint8 = 12
	workerMax int64 = 1 << workerBit
	nuberMax  int64 = 1 << nuberBit
	timeShift       = workerBit + nuberBit
	startTime int64 = 1671702133068
)

type Snow struct {
	mu        sync.Mutex //互斥锁
	timestamp int64      //当前的时间戳
	workerId  int64      //共同的机器共有1024台
	number    int64      //同一时间内生成的数量最多4096个
}

// NewSnow 初始化雪花ID对象
func NewSnow(id int64) *Snow {
	if id < 0 || id > workerMax {
		panic("worker Id must > 0 and < 1024")
	}
	return &Snow{
		timestamp: 0,
		workerId:  id,
		number:    0,
	}
}
func (s *Snow) GetId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.number++
		if s.number > nuberMax {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.number = 0
		s.timestamp = now
	}
	return int64((now-startTime)<<timeShift | (s.workerId << nuberBit) | s.number)
}
