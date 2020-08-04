package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Roll string `json:"roll"`
}

var bufferpool = sync.Pool{
	New: func() interface{} {
		return new(Person)
	},
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/raj", raj)
	http.ListenAndServe(":8080", r)
}
func raj(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		time.Sleep(2 * time.Second)
		http.Error(w, "No handler found", 404)
		return
	}
	//sl := bufferpool.Get().(*Person)
	sl := Person{}
	fmt.Println(sl)
	json.NewDecoder(r.Body).Decode(sl)
	fmt.Println(sl)
	//sl.Roll = "4"
	//	bufferpool.Put(sl)

}

func Hello() {
	for {
		time.AfterFunc(2*time.Second, func() {
			fmt.Println("hello called")
		})

		fmt.Println("called")

	}
}
func ThisIsToCheck() {
	fmt.Println("hello raj")
}
