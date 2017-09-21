package utils

import "time"

// EpochNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC.
func EpochNano() int64 {
	return time.Now().UnixNano()
}

// Epoch returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func Epoch() int64 {
	return time.Now().Unix()
}
