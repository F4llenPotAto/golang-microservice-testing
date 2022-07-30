package main

import (
	"log"
	"net/http"
	"os"

	"github.com/F4llenPotAto/golang-microservice-testing/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe("localhost:9090", sm)

}
