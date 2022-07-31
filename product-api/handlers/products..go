package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	l *l.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {

}
