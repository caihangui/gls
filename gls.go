package gls

import (
	"sync"
)

// glsMapNum the length of goroutineLocalStore.
const glsMapNum = 16

// goroutineLocalStore
var goroutineLocalStore = [glsMapNum]sync.Map{}

// Store stores the specific key/val
func Store(key string, val interface{}) {
	val2, ok := Load(key)
	if !ok {
		val2 = map[string]interface{}{}
	}
	val2.(map[string]interface{})[key] = val

	goid := GetGoroutineID()
	goroutineLocalStore[goid%glsMapNum].Store(goid, val2)
}

// Load loads the specific data
func Load(key string) (val interface{}, ok bool) {
	goid := GetGoroutineID()
	m, ok := goroutineLocalStore[goid%glsMapNum].Load(goid)
	if !ok {
		return nil, ok
	}
	val, ok = m.(map[string]interface{})[key]
	return val, ok
}

// Delete cleans the specific data
func Delete(key string) {
	goid := GetGoroutineID()
	value, ok := goroutineLocalStore[goid%glsMapNum].Load(goid)
	if !ok {
		return
	}
	delete(value.(map[string]interface{}), key)
}

// DeleteAll cleans all data under current goroutine.
// you should call this func when your goroutine exited
// to avoid anther goroutine get the dirty data next time.
func DeleteAll() {
	goid := GetGoroutineID()
	goroutineLocalStore[goid%glsMapNum].Delete(goid)
}
