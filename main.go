package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		_, err := fmt.Fprintf(w, "I'm a teapot\n")
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
