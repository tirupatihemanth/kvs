package main

import "net/http"
type KVResponse struct {
	Result string `json:"result"`
	Ok     bool   `json:"ok"`
}

func getKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Get(key)
	if !resp.Ok{
		respondWithError(w, http.StatusBadRequest, "Key does not exist")
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

func putKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	val, ok := get_header(r.Header, "Val")
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Please provide a value for the key")
		return
	}
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Put(key, val)
	if !resp.Ok{
		respondWithError(w, http.StatusBadRequest, "Put failed")
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

func delKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Del(key)
	if !resp.Ok{
		respondWithError(w, http.StatusBadRequest, "Key does not exist")
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}
