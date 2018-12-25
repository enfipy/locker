package locker

import (
	"log"
	"sync"
	"testing"
	"time"
)

func run(lk *Locker, wg *sync.WaitGroup) {
	lk.Lock("test")
	time.Sleep(500 * time.Millisecond)
	log.Printf("Pending count: #%d", lk.locks["test"].pending)
	lk.Unlock("test")

	wg.Done()
}

func TestLocker(_ *testing.T) {
	lk := Initialize()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go run(lk, &wg)
	}
	wg.Wait()
}
