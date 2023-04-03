package main

import "sync"

type MemoryCache struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)
type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *MemoryCache {
	return &MemoryCache{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *MemoryCache) GetData(key int) (interface{}, error) {
	m.lock.Lock()
	result, ok := m.cache[key]
	m.lock.Unlock()

	if !ok {
		m.lock.Lock()
		result.value = Fibonacci(key)
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}
