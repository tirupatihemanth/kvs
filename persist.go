package main

import (
	"encoding/json"
	"os"
)

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
