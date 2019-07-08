package main

import (
	"dynamic-equity/pkg/api"
	"log"
	"net/http"
)

func main() {
	r := api.New()
	log.Println("Start server...")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
