package main

import (
	"dictionnaire/httphandler"
	"net/http"
	// "time"
)

func main() {

	http.HandleFunc("/hello", httphandler.Hello)
	http.HandleFunc("/list", httphandler.List)
	http.HandleFunc("/get", httphandler.Get)
	http.HandleFunc("/add", httphandler.Add)
	http.HandleFunc("/delete", httphandler.Remove)

	http.ListenAndServe(":8089", nil)

}
