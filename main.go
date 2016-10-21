package main

import (
	"fmt"
    "net/http"
)

func main() {
	fmt.Println("Starting local server on port :3000")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Hello world..."))
}










