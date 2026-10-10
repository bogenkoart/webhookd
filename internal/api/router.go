package api

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthzPing)
	return mux
}
