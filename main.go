package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", fileForm)
	http.HandleFunc("/upload", fileUpload)
	log.Fatal(http.ListenAndServe(":80", nil))


	
}
