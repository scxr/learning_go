package main

import (
	"fmt"
	"net/http"
)

func indexhandler(writer http.ResponseWriter, req *http.Request) { 
	fmt.Fprintf(writer, "Hello world")
}

func main() {
	http.HandleFunc("/", indexhandler) // index page handler
	http.ListenAndServe(":4444", nil) // serve localhost port 4444
}

