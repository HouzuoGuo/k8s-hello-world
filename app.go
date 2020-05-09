package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Going to listen on port %s", os.Getenv("PORT"))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Serving URL \"%s\" to visitor \"%s:%s\"", req.Host, req.URL.Port(), req.URL)
		hostName, _ := os.Hostname()
		_, _ = fmt.Fprintf(w, "Hello from host name %s, your host appears to be %s, your requested URL is:%s", hostName, req.Host, req.URL)
	})

	if err := http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
