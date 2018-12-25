package locker

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type lock struct {
	*sync.RWMutex
	pending int32
}

func initLock() *lock {
	return &lock{
		RWMutex: new(sync.RWMutex),
		pending: 0,
	}
}

// Locker type for map of mutexes and self lock to add new mutex to map
type Locker struct {
	locks    map[string]*lock
	selfLock *sync.RWMutex
}

// Initialize - return initialized locker struct
func Initialize() *Locker {
	return &Locker{
		locks:    make(map[string]*lock),
		selfLock: new(sync.RWMutex),
	}
}

// Lock - locks mutex
func (lkr *Locker) Lock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		lk = lkr.newLock(key)
	}
	atomic.AddInt32(&lk.pending, 1)
	lk.Lock()
}

// Unlock - unlocks mutex
func (lkr *Locker) Unlock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		panic(fmt.Errorf("Lock for key '%s' is not initialized", key))
	}
	atomic.AddInt32(&lk.pending, -1)
	lk.Unlock()
	if lk.pending < 1 {
		lkr.remLock(key)
	}
}

// RLock - locks rw for reading
func (lkr *Locker) RLock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		lk = lkr.newLock(key)
	}
	atomic.AddInt32(&lk.pending, 1)
	lk.RLock()
}

// RUnlock - unlocks a single RLock call
func (lkr *Locker) RUnlock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		panic(fmt.Errorf("Lock for key '%s' is not initialized", key))
	}
	atomic.AddInt32(&lk.pending, -1)
	lk.RUnlock()
	if lk.pending < 1 {
		lkr.remLock(key)
	}
}

func (lkr *Locker) newLock(key string) *lock {
	lkr.selfLock.Lock()
	defer lkr.selfLock.Unlock()

	if lk, ok := lkr.locks[key]; ok {
		return lk
	}
	lk := initLock()
	lkr.locks[key] = lk
	return lk
}

func (lkr *Locker) getLock(key string) (*lock, bool) {
	lkr.selfLock.RLock()
	defer lkr.selfLock.RUnlock()

	lock, ok := lkr.locks[key]
	return lock, ok
}

func (lkr *Locker) remLock(key string) {
	lkr.selfLock.Lock()
	defer lkr.selfLock.Unlock()

	if _, ok := lkr.locks[key]; ok {
		delete(lkr.locks, key)
	}
}
