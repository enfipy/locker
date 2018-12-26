package locker

import (
	"log"
	"sync"
	"testing"
	"time"
)

func testLock(lkr *Locker, wg *sync.WaitGroup, text string, sleep time.Duration) {
	lkr.Lock("test")
	time.Sleep(sleep)
	if lk, ok := lkr.locks.Load("test"); ok {
		log.Printf("Pending count from '%s': #%d", text, lk.(*lock).pending)
	} else {
		log.Fatal("Can not load lock")
	}
	lkr.Unlock("test")
	wg.Done()
}

func testRLock(lkr *Locker, wg *sync.WaitGroup, text string, sleep time.Duration) {
	lkr.RLock("test")
	time.Sleep(sleep)
	if lk, ok := lkr.locks.Load("test"); ok {
		log.Printf("Pending count from '%s': #%d", text, lk.(*lock).pending)
	} else {
		log.Fatal("Can not load lock")
	}
	lkr.RUnlock("test")
	wg.Done()
}

func TestLocker(_ *testing.T) {
	var waitForTests sync.WaitGroup
	waitForTests.Add(2)
	// Locker
	go func() {
		lkr := Initialize()
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go testLock(lkr, &wg, "Mr. Writelock", 100*time.Millisecond)
		}
		wg.Wait()
		waitForTests.Done()
	}()
	// RLocker
	go func() {
		lkr := Initialize()
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go testRLock(lkr, &wg, "Mr. Readlock", 100*time.Millisecond)
		}
		wg.Wait()
		waitForTests.Done()
	}()
	waitForTests.Wait()
}
