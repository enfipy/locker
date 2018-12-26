package locker

import (
	"log"
	"sync"
	"testing"
	"time"
)

func testLock(lkr *Locker, wg *sync.WaitGroup, num int, sleep time.Duration) {
	lkr.Lock("test")
	time.Sleep(sleep)
	log.Printf("Write lock #%d", num)
	lkr.Unlock("test")
	wg.Done()
}

func testRLock(lkr *Locker, wg *sync.WaitGroup, num int, sleep time.Duration) {
	lkr.RLock("test")
	time.Sleep(sleep)
	log.Printf("Read lock #%d", num)
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
			go testLock(lkr, &wg, i, 100*time.Millisecond)
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
			go testRLock(lkr, &wg, i, 100*time.Millisecond)
		}
		wg.Wait()
		waitForTests.Done()
	}()
	waitForTests.Wait()
}
