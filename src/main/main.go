package main

import (
	"net/http"
	"api"
)

func main() {
	http.HandleFunc("/transfer", api.TransferHandler)
	http.ListenAndServe(":3000", nil)
}