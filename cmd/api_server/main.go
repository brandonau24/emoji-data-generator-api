package main

import (
	"log"
	"net/http"

	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/request_handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &request_handlers.EmojisHandler{})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
