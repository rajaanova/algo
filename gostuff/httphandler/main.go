package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/raj", http.HandlerFunc(raj))
	http.ListenAndServe(":8080", nil)
}
func raj(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello raj")
}
