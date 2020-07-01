package handler

import "net/http"

type productHandler struct {
}

//NewProductHandler app
func NewProductHandler() IHandler {
	return productHandler{}
}

func (h productHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product"))
}
