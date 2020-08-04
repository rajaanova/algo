package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	//b := make([]byte, 25)
	bf := bytes.Buffer{}
	fs, er := os.OpenFile("newfile", os.O_RDONLY, 0666)
	fmt.Println(er)
	rea := bufio.NewReader(fs)
	scanner := bufio.NewScanner(fs)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
	for {
		rea.ReadString('\n')
	}

	fmt.Println("fs is opneind now ", fs, er)
	if er != nil {
		panic("unable to open file")
	}
	for {
		//_, err := fs.Read(b)
		_, err := fs.Read(bf.Bytes())
		//fmt.Println(b, err)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic("error reading file ")
		}
		fmt.Println("b has ", bf.String())
	}
}
