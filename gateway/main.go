package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRouter(mux)
	http.ListenAndServe(":9000", mux)
}
