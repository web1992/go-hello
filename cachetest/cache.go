package cachetest

import "sync"

// 缓存的实现一
var (
	m       sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	m.Lock()
	v := mapping[key]
	m.Unlock()
	return v
}

// 缓存的实现二
var cache = struct {
	// 匿名
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup2(key string) string {

	cache.Lock()

	v := cache.mapping[key]

	cache.Unlock()

	return v
}
