package main

import (
	"net/http"
	"log"
)

type keyHandler func(w http.ResponseWriter, r *http.Request, key string)

func get_header(headers http.Header, header string) (string, bool) {
	key := headers.Get(header)
	if key == "" {
		return "", false
	}
	return key, true
}

func middleware_key(kh keyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		key, ok := get_header(r.Header, "Key")
		if !ok {
			respondWithError(w, http.StatusBadRequest, "Please provide a key")
			return
		}
		log.Println("Key:", key)
		kh(w, r, key)
	}
}
