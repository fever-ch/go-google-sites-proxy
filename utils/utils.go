package utils

import "time"

// Maximum unsigned int
const MaxUint = ^uint(0)
// Minimum unsigned int
const MinUint = 0

// Maximum signed int
const MaxInt = int(MaxUint >> 1)
// Minimum signed int
const MinInt = -MaxInt - 1

// Maximum unsigned int16
const MaxUint16 = ^uint16(0)
// Minimum unsigned int16
const MinUint16 = 0

// Maximum signed int16
const MaxInt16 = int16(MaxUint16 >> 1)
// Minimum signed int16
const MinInt16 = -MaxInt16 - 1

// Maximum unsigned int32
const MaxUint32 = ^uint32(0)
// Minimum unsigned int32
const MinUint32 = 0

// Maximum signed int32
const MaxInt32 = int32(MaxUint32 >> 1)
// Minimum signed int32
const MinInt32 = -MaxInt32 - 1

// Maximum unsigned int64
const MaxUint64 = ^uint64(0)
// Minimum unsigned int64
const MinUint64 = 0

// Maximum signed int64
const MaxInt64 = int64(MaxUint64 >> 1)
// Minimum signed int64
const MinInt64 = -MaxInt64 - 1

// Epoch in nanoseconds
func EpochNano() int64 {
	return time.Now().UnixNano()
}

// Epoch in seconds
func Epoch() int64 {
	return time.Now().Unix()
}
