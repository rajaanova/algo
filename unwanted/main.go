package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"reflect"
	"time"
)

type stack []interface{}
type Dog struct {
	Breed    string
	WeightKg someint `json:",omitempty"`
}
type someint int

func (a *someint) UnmarshalJSON(value []byte) error {
	var ab float64
	json.Unmarshal(value, &ab)
	*a = someint(ab) + 4
	fmt.Println(ab)

	return nil
}

func main() {
	var m Dog
	st := `{"Breed":"some good kind ","WeightKg":89.23}`
	data := []byte(st)
	json.Unmarshal(data, &m)
	fmt.Println(m)
	var a []string
	fmt.Println(reflect.TypeOf(a))
	sta := initStack(5)
	sta.push(10)
	sta.push(11)
	x := sta.pop()
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(sta)
	d := exec.Command("xdg-open", "www.google.com").Start()
	time.Sleep(10 * time.Second)
	fmt.Println(d)
	//dd := m["WeightKg"].(float64)
	//x, _ := json.Marshal(dd)

	//fmt.Println(dd)
}

func initStack(i int) stack {
	var a stack = make(stack, 0, i)
	return a
}
func (a *stack) pop() interface{} {
	x := (*a)[0]
	*a = (*a)[1:]
	return x
}
func (a *stack) push(i interface{}) {
	var j stack
	j = append(j, i)
	d := append(j, *a...)
	*a = d

}
