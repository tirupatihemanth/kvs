package main

import "net/http"

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
	val := r.Header.Get("Val")
	kvMap.Put(key, val)
	respondWithJSON(w, http.StatusOK, KVResponse{"", true})
}

func delKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	kvMap.Del(key)
	respondWithJSON(w, http.StatusOK, KVResponse{"", true})
}
