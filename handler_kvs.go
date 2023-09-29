package main

import (
	"net/http"
	"fmt"
)
type KVResponse struct {
	Result string `json:"result"`
	Ok     bool   `json:"ok"`
}

func getKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Get(key)
	respondWithJSON(w, http.StatusOK, resp)
}

func putKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	val, ok := get_header(r.Header, "Val")
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Please provide a value for the key")
		return
	}
	fmt.Printf(val)
	kvMap.Put(key, val)
	respondWithJSON(w, http.StatusOK, KVResponse{"", true})
}

func delKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	kvMap.Del(key)
	respondWithJSON(w, http.StatusOK, KVResponse{"", true})
}
