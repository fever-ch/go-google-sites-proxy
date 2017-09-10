// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.
//
package common

import (
	"sync"
	"unsafe"
	"sync/atomic"
)

type LazyContainer struct {
	Get func() unsafe.Pointer
}

const (
	EMPTY uint32 = iota
	GETTING_READY
	READY
)

func NewLazyContainer(f func() unsafe.Pointer) *LazyContainer {
	lc := LazyContainer{}
	var content unsafe.Pointer
	var mtx sync.Mutex

	status := EMPTY

	prepare := func() {
		mtx.Lock()
		atomic.StorePointer(&content, f())

		// next getters doesn't need to look at any lock
		atomic.StoreUint32(&status, READY)
		mtx.Unlock()
	}

	lc.Get = func() unsafe.Pointer {
		if atomic.LoadUint32(&status) != READY {
			if atomic.CompareAndSwapUint32(&status, EMPTY, GETTING_READY) {
				prepare()
			} else {
				// IT IS GETTING READY
				// tries to acquire the lock and release it
				for atomic.LoadUint32(&status) != READY {
					mtx.Lock()
					mtx.Unlock()
				}
			}
		}

		return atomic.LoadPointer(&content)
	}

	return &lc
}
