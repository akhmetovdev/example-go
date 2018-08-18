package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r)
	fmt.Fprintln(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Example Go Server")
	http.ListenAndServe(":3000", nil)
}
