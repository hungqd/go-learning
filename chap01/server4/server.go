package main

import (
	"log"
	"net/http"

	lissajous "github.com/hungqd/go-learning/chap01/lissajous/lib"
)

func main() {
	http.HandleFunc("GET /", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}
