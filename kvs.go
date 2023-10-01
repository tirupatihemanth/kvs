package main

import "sync"

type KVMap struct {
	data map[string]string
	mu   sync.RWMutex
}

func (kvs *KVMap) Put(key string, val string) (string, bool) {
	kvs.mu.Lock()
	kvs.data[key] = val
	val, ok := kvs.data[key]
	kvs.mu.Unlock()
	return val, ok
}

func (kvs *KVMap) Get(key string) (string, bool) {
	kvs.mu.RLock()
	val, ok := kvs.data[key]
	kvs.mu.RUnlock()
	return val, ok
}

func (kvs *KVMap) Del(key string) (string, bool) {
	kvs.mu.Lock()
	val, ok := kvs.data[key]
	delete(kvs.data, key)
	kvs.mu.Unlock()
	return val, ok
}
