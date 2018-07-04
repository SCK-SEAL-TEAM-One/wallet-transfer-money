package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/test", testserve)
	http.ListenAndServe(":3000", nil)
}

func testserve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}