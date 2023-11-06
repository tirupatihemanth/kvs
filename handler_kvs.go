package main

import "net/http"

type KVResponse struct {
	Result string `json:"result"`
	Ok     bool   `json:"ok"`
}

// Handles Get Key requests
func getKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Get(key)
	if !resp.Ok {
		resp.Result = "Key does not exist"
		respondWithJSON(w, http.StatusOK, resp)
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

// Handles Put Key requests
func putKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	val, ok := get_header(r.Header, "Val")
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Please provide a value for the key")
		return
	}
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Put(key, val)
	if !resp.Ok {
		resp.Result = "Put failed"
		respondWithJSON(w, http.StatusInternalServerError, resp)
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}

// Handles Del Key requets.
func delKeyHandler(w http.ResponseWriter, r *http.Request, key string) {
	var resp KVResponse
	resp.Result, resp.Ok = kvMap.Del(key)
	if !resp.Ok {
		resp.Result = "Key does not exist"
		respondWithJSON(w, http.StatusOK, resp)
		return
	}
	respondWithJSON(w, http.StatusOK, resp)
}
