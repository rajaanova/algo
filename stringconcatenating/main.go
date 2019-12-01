package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	builderBuffer()
	bytesBuffer()
}
func builderBuffer() {
	listToConcatenage := []string{"hello", "raj", "Agrawal"}
	var builder strings.Builder
	for _, val := range listToConcatenage {
		builder.WriteString(val)
	}
	fmt.Println(builder.String())
}
func bytesBuffer() {
	listToConcatenage := []string{"hello", "raj", "Agrawal"}
	var buf bytes.Buffer
	for _, val := range listToConcatenage {
		buf.WriteString(val)
	}
	fmt.Println(buf.String())
}
