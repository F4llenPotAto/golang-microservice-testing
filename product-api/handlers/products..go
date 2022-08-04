package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"product-api/data"
)

type Products struct {
	l *l.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	json.Marshal()
}
