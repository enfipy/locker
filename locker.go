package locker

import (
	"fmt"
	"sync"
)

// Locker type for map of mutexes and self lock to add new mutex to map
type Locker struct {
	locks    map[string]*sync.RWMutex
	selfLock *sync.RWMutex
}

// Initialize - return initialized locker struct
func Initialize() *Locker {
	return &Locker{
		locks:    make(map[string]*sync.RWMutex),
		selfLock: new(sync.RWMutex),
	}
}

// Lock - locks mutex
func (lkr *Locker) Lock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		lk = lkr.newLock(key)
	}
	lk.Lock()
}

// Unlock - unlocks mutex
func (lkr *Locker) Unlock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		panic(fmt.Errorf("Lock for key '%s' is not initialized", key))
	}
	lk.Unlock()
}

// RLock - locks rw for reading
func (lkr *Locker) RLock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		lk = lkr.newLock(key)
	}
	lk.RLock()
}

// RUnlock - unlocks a single RLock call
func (lkr *Locker) RUnlock(key string) {
	lk, ok := lkr.getLock(key)
	if !ok {
		panic(fmt.Errorf("Lock for key '%s' is not initialized", key))
	}
	lk.RUnlock()
}

func (lkr *Locker) newLock(key string) *sync.RWMutex {
	lkr.selfLock.Lock()
	defer lkr.selfLock.Unlock()

	if lk, ok := lkr.locks[key]; ok {
		return lk
	}
	lk := new(sync.RWMutex)
	lkr.locks[key] = lk
	return lk
}

func (lkr *Locker) getLock(key string) (*sync.RWMutex, bool) {
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

// Todo: Add remLock
