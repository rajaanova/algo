package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type raj struct {
	Name string `json:name`
	D    string
}

func main() {
	s := `{"name":"raj","date":1577885312}`
	customJsonUnMarshaller(s)
}

func customJsonUnMarshaller(s string) {
	var r raj
	json.Unmarshal([]byte(s), &r)
	fmt.Println(r.Name)
	fmt.Println(r.D)

}

func (r *raj) UnmarshalJSON(b []byte) error {
	type X raj

	aux := &struct {
		Dat int64 `json:"date"`
		*X
	}{}
	json.Unmarshal(b, &aux)
	fmt.Println(aux.Name)
	r.D = strconv.FormatInt(aux.Dat, 10)

	return nil
}
