package utils

import "time"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int64(MaxUint >> 1)
const MinInt = -MaxInt - 1

const MaxUint16 = ^uint16(0)
const MinUint16 = 0
const MaxInt16 = int64(MaxUint16 >> 1)
const MinInt16 = -MaxInt16 - 1

const MaxUint32 = ^uint32(0)
const MinUint32 = 0
const MaxInt32 = int64(MaxUint32 >> 1)
const MinInt32 = -MaxInt32 - 1

const MaxUint64 = ^uint64(0)
const MinUint64 = 0
const MaxInt64 = int64(MaxUint64 >> 1)
const MinInt64 = -MaxInt64 - 1

func EpochNano() int64 {
	return time.Now().UnixNano()
}

func Epoch() int64 {
	return time.Now().Unix()
}
