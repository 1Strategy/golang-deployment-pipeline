package main

import (
	"fmt"
	"log"
	"net/http"
)

const version string = "2.0.1"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
// func versionHandler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, version)
// }

func main() {
	log.Printf("Listening on port 8000...")
	//http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n Building GO Applications in AWS is awesome!", r.URL.Path)
	})
	http.ListenAndServe(":8000", nil)
}
