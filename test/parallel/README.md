## 锁机制

写锁，在写的时候加锁，不允许其他操作

```
func WriteToMap(k, v string) {
	// DELETE lock and unlock code, then run go test . -race, it must fail
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}
```

读锁

```
func ReadFromMap(k string) string {
	// DELETE lock and unlock code, then run go test . -race, it must fail
	locker.RLock()
	defer locker.RUnlock()
	return data[k]
}
```

