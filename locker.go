package locker

import "sync"

// Locker type for map of mutexes and self lock to add new mutex to map
type Locker struct {
	pool  sync.Pool
	locks sync.Map
}

// Initialize - return initialized locker struct
func Initialize() *Locker {
	return &Locker{
		pool: sync.Pool{
			New: func() interface{} {
				return new(sync.RWMutex)
			},
		},
	}
}

// Lock - locks mutex
func (lkr *Locker) Lock(key interface{}) {
	lkr.getLock(key).Lock()
}

// Unlock - unlocks mutex
func (lkr *Locker) Unlock(key interface{}) {
	lkr.getLock(key).Unlock()
}

// RLock - locks rw for reading
func (lkr *Locker) RLock(key interface{}) {
	lkr.getLock(key).RLock()
}

// RUnlock - unlocks a single RLock call
func (lkr *Locker) RUnlock(key interface{}) {
	lkr.getLock(key).RUnlock()
}

func (lkr *Locker) getLock(key interface{}) *sync.RWMutex {
	newLock := lkr.pool.Get()
	lock, stored := lkr.locks.LoadOrStore(key, newLock)
	if stored {
		lkr.pool.Put(newLock)
	}
	return lock.(*sync.RWMutex)
}
