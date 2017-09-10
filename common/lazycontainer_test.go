// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package common

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"strconv"
	"sync"
	"time"
	"sync/atomic"
	"unsafe"
)

func TestLazyContainer(t *testing.T) {

	const entries = 100000
	const iterations = 20
	const extra = 100

	var wg sync.WaitGroup

	var containers [entries + extra]*LazyContainer

	var computations int32 = 0

	for i := 0; i < entries+extra; i++ {
		z := i
		containers[i] = NewLazyContainer(
			func() unsafe.Pointer { time.Sleep(5 * time.Millisecond); pt := strconv.Itoa(z); atomic.AddInt32(&computations, 1); return unsafe.Pointer(&pt) })
	}

	for i := 0; i < entries; i+=2 {
		for j := 0; j < iterations; j++ {
			wg.Add(1)
			go func(idx int) {
				assert.Equal(t, strconv.Itoa(idx), *(*string)(containers[idx].Get()))
				wg.Done()
			}(i)
		}
	}

	wg.Wait()

	assert.Equal(t, entries/2, int(computations))
}
