package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/F4llenPotAto/golang-microservice-testing/product-api/data"
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
