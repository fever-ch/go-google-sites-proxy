package common

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type LazyContainer struct {
	Get func() unsafe.Pointer
}

const (
	empty         uint32 = iota
	getting_ready
	ready
)

func NewLazyContainer(f func() unsafe.Pointer) *LazyContainer {
	lc := LazyContainer{}
	var content unsafe.Pointer
	var mtx sync.Mutex

	status := empty

	prepare := func() unsafe.Pointer {
		mtx.Lock()
		value := f()
		atomic.StorePointer(&content, value)

		// next getters doesn't need to look at any lock
		defer atomic.StoreUint32(&status, ready)
		defer mtx.Unlock()
		return value
	}

	lc.Get = func() unsafe.Pointer {
		if atomic.LoadUint32(&status) != ready {
			if atomic.CompareAndSwapUint32(&status, empty, getting_ready) {
				return prepare()
			}
			// IT IS GETTING READY
			// tries to acquire the lock and release it
			for atomic.LoadUint32(&status) != ready {
				mtx.Lock()
				mtx.Unlock()
			}
		}

		return atomic.LoadPointer(&content)
	}

	return &lc
}
