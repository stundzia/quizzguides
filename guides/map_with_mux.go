package guides

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// A lot of junior devs are aware that they need to use a mutex for concurrent map
// access handling (race condition prevention), but are unaware that they can use

type muxedMap interface {
	Get(count int) int
	Add(k int, v string)
}

type mapWithMux struct {
	m       map[int]string
	mux     sync.Mutex
	opCount *int32
}

type mapWithRWMux struct {
	m       map[int]string
	mux     sync.RWMutex
	opCount *int32
}

func (mm *mapWithMux) Add(k int, v string) {
	mm.mux.Lock()
	defer mm.mux.Unlock()
	defer atomic.AddInt32(mm.opCount, 1)
	mm.m[k] = v
}

func (mr *mapWithRWMux) Add(k int, v string) {
	mr.mux.Lock()
	defer mr.mux.Unlock()
	defer atomic.AddInt32(mr.opCount, 1)
	mr.m[k] = v
}

func (mm *mapWithMux) Get(count int) int {
	res := 0
	mm.mux.Lock()
	defer mm.mux.Unlock()
	defer atomic.AddInt32(mm.opCount, 1)
	for k, _ := range mm.m {
		res = res + k
		count--
		if count <= 0 {
			break
		}
	}
	return res
}

func (mr *mapWithRWMux) Get(count int) int {
	res := 0
	mr.mux.RLock()
	defer mr.mux.RUnlock()
	defer atomic.AddInt32(mr.opCount, 1)
	for k, _ := range mr.m {
		res = res + k
		count--
		if count <= 0 {
			break
		}
	}
	return res
}

func testIt(mm muxedMap) {
	for i := 1; i < 6; i++ {
		go func() {
			i := i
			for t := 1; ; t++ {
				go mm.Add(t*i, fmt.Sprintf("asdf%d", i*t))
			}
		}()
		go func() {
			for t := 1; ; t++ {
				go mm.Get(1 + t)
			}
		}()
	}
}

func MuxVsRWMux() {
	var opCount int32 = 0
	mm := mapWithMux{
		m:       map[int]string{},
		mux:     sync.Mutex{},
		opCount: &opCount,
	}
	testIt(&mm)
	time.Sleep(time.Second * 10)
	fmt.Println("Ops per second: ", *mm.opCount/10)
}
