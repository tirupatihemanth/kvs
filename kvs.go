package main

import "sync"

type KVMap struct {
	data map[string]string
	mu   sync.RWMutex
}

func (kvs *KVMap) Put(key string, val string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	kvs.data[key] = val
}

func (kvs *KVMap) Get(key string) (string, bool) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()
	val, ok := kvs.data[key]
	return val, ok
}

func (kvs *KVMap) Del(key string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	delete(kvs.data, key)
}
