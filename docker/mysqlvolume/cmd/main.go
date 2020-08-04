package main

import (
	"encoding/base64"
)

func main() {
	base64.StdEncoding.EncodeToString([]byte("pleasure."))
	//mux := http.NewServeMux()
	//controller := student.NewController()
	//mux.HandleFunc("/put", controller.PutStudent)
	//http.ListenAndServe(":8080", mux)
}
