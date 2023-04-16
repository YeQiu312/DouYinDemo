package util

import (
	"errors"
	"time"
)

const (
	// epoch是开始时间戳，可以自定义
	epoch        int64 = 1621104000000 // 2021-05-16 00:00:00 UTC+8
	workerIDBits uint8 = 10
	sequenceBits uint8 = 12
	maxWorkerID  int64 = -1 ^ (-1 << workerIDBits)
	maxSequence  int64 = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
	workerID      int64
	lastTimestamp int64
	sequence      int64
}

func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	return &Snowflake{
		workerID: workerID,
	}, nil
}

func (s *Snowflake) NextID() (int64, error) {
	timestamp := s.timeGen()
	if timestamp < s.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			timestamp = s.tilNextMillis(s.lastTimestamp)
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = timestamp
	return ((timestamp - epoch) << (workerIDBits + sequenceBits)) | (s.workerID << sequenceBits) | s.sequence, nil
}

func (s *Snowflake) timeGen() int64 {
	return time.Now().UnixNano()/int64(time.Millisecond) - epoch
}

func (s *Snowflake) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := s.timeGen()
	for timestamp <= lastTimestamp {
		timestamp = s.timeGen()
	}
	return timestamp
}
