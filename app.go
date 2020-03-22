package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Going to listen on port %s", os.Getenv("PORT"))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		hostName, _ := os.Hostname()
		_, _ = w.Write([]byte("Hello from " + hostName))
	})

	if err := http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
