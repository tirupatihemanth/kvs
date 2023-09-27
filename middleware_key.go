package main

import "net/http"

type keyHandler func(w http.ResponseWriter, r *http.Request, key string)

func getKey(headers http.Header) (string, bool) {
	key := headers.Get("Key")
	if key == "" {
		return "", false
	}
	return key, true
}

func middleware_key(kh keyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, ok := getKey(r.Header)
		if !ok {
			respondWithError(w, http.StatusBadRequest, `Provide a non-empty key as the header "Key"`)
			return
		}
		kh(w, r, key)
	}
}
