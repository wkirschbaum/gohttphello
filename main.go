package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	envPort := os.Getenv("PORT")
	location := ":3000"
	if len(envPort) > 0 {
		location = ":" + envPort
	}

	http.ListenAndServe(location, nil)
}
