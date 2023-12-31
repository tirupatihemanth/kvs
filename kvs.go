package main

import "sync"

type KVMap struct {
	data map[string]string
	mu   sync.RWMutex
}

// Functions to access Dictionary underlying KV Store

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
	if ok {
		delete(kvs.data, key)
	}
	kvs.mu.Unlock()
	return val, ok
}
