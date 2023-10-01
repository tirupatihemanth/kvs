package main

import (
	"encoding/json"
	"os"
	"time"
)

const (
	SCRAPING_INTERVAL = time.Second * 5
	PERSIST_FILE_NAME = "persist.json"
)

func scheduleSaving() {

	ticker := time.NewTicker(SCRAPING_INTERVAL)
	for ; ; <-ticker.C {
		kvMap.mu.RLock()
		kvMap.SaveToFile(PERSIST_FILE_NAME)
		kvMap.mu.RUnlock()
	}
}

func (kvMap *KVMap) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(kvMap.data); err != nil {
		return err
	}

	return nil
}

func (kvMap *KVMap) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&kvMap.data); err != nil {
		return err
	}

	return nil
}
