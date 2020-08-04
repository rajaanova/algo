package main

import "fmt"

func main() {
	fmt.Println(taumBday(62990, 61330, 310144, 312251, 93023))
}
func taumBday(b int64, w int64, bc int64, wc int64, z int64) int64 {
	// Write your code here
	min := int64(b*bc + w*wc)
	bconvert := int64((b+w)*bc + w*z)
	if min > bconvert {
		min = bconvert
	} else {
		wconvert := int64((b+w)*wc + b*z)
		if wconvert < min {
			min = wconvert
		}
	}
	return min

}
