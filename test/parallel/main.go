package main

import "sync"

var (
	data   = make(map[string]string)
	locker sync.RWMutex
)

func WriteToMap(k, v string) {
	// DELETE lock and unlock code, then run go test . -race, it must fail
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}

func ReadFromMap(k string) string {
	// DELETE lock and unlock code, then run go test . -race, it must fail
	locker.RLock()
	defer locker.RUnlock()
	return data[k]
}
