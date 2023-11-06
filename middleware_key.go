package main

import (
	"log"
	"net/http"
)

type keyHandler func(w http.ResponseWriter, r *http.Request, key string)

func get_header(headers http.Header, header string) (string, bool) {
	key := headers.Get(header)
	if key == "" {
		return "", false
	}
	return key, true
}

// Extracts key from the http request and calls the handler passed in as the argument.
// Follows Chain of Responsibility Design Pattern.
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
