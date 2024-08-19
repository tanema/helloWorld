package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handler)
	log.Printf("listening on %v\n", port)
	log.Fatal(http.ListenAndServe(port, logRequest(http.DefaultServeMux)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello world!\n"); err != nil {
		log.Fatal(err)
	}
}

func logRequest(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(rw, req)
		duration := time.Since(start)
		log.Printf("%v %v %v \n", req.Method, req.RequestURI, duration)
	}
	return http.HandlerFunc(loggingFn)
}
