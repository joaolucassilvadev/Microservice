package main

import "net/http"

type handler struct {
	// gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRouter(mux *http.ServeMux) {
	mux.HandleFunc("/api/customers/{customerId}/orders", h.ServeHTTP)
}

func (h *handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
