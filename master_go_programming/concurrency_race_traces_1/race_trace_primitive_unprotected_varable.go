package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	wd := Watchdog{
		last: 12,
	}

	wd.KeepAliveWithoutRace()
	wd.StartWithoutRace()
}

type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

// Solution: A typical fix for this race is to use a channel or a mutex. To preserve the lock-free behavior, one can also use the sync/atomic package.

func (w *Watchdog) KeepAliveWithoutRace() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) StartWithoutRace() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
