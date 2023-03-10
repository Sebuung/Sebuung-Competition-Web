package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	 }
	return_string := string(data)
	fmt.Fprintf(w, return_string)
	
}
