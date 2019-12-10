package main

import "fmt"

func main() {
	doubleval := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	moveSpiral(doubleval)
}
func moveSpiral(a [][]int) {
	rowlength := len(a) - 1
	columnLength := len(a[0]) - 1
	colstart := 0
	rowstart := 0
	//j := rowlength
	for rowstart <= rowlength && colstart <= columnLength {
		for i := 0; i <= columnLength; i++ {
			fmt.Println(a[rowstart][i])
		}
		for row := rowstart + 1; row <= rowlength; row++ {
			fmt.Println(a[row][columnLength])
		}
		for k := columnLength - 1; k >= colstart; k-- {
			fmt.Println(a[rowlength][k])
		}
		for k := rowlength - 1; k > rowstart; k-- {
			fmt.Println(a[k][colstart])
		}
		rowstart++
		colstart++
		rowlength--
		columnLength--
	}

}
