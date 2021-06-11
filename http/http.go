package http1

import "net/http"

type Hanler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func ListenAndServe(address string, h Hanler) error
