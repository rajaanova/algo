package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	_, port, _ := SplitHostPort("localhost:9090")
	fmt.Println(strconv.Itoa(port))
	fmt.Println(string(byte(49)))
	fmt.Println(strconv.Atoi("48"))
	fmt.Println(strconv.Itoa(48))
}

func SplitHostPort(addr string) (host string, port int, err error) {
	h, p, err := net.SplitHostPort(addr)
	if err != nil {
		return "", -1, err
	}
	if p == "" {
		return "", -1, &net.AddrError{Err: "missing port in address", Addr: addr}
	}

	pi, err := strconv.Atoi(p)
	if err != nil {
		return "", -1, err
	}
	return h, pi, nil
}
