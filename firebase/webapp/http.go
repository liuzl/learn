package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("port", ":9090", "port to be used by the server")
)

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, api_key, Authorization")
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	http.Handle("/", corsMiddleware(http.FileServer(http.Dir(""))))
	log.Printf("Running at http://127.0.0.1%s", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
