package main

import "fmt"

func main() {
	stockprices := []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	fmt.Println(buyingandseleling(stockprices))
}

func buyingandseleling(stockPrices []int) int {

	minsofar := 1289128
	maxsofar := 0
	maxprofitbyselling := 0
	maxprofitbysellingByIndices := make([]int, len(stockPrices))
	for i := 0; i < len(stockPrices); i++ {
		if stockPrices[i] < minsofar {
			minsofar = stockPrices[i]
		}
		if maxprofitbyselling < stockPrices[i]-minsofar {
			maxprofitbyselling = stockPrices[i] - minsofar

		}
		maxprofitbysellingByIndices[i] = maxprofitbyselling
	}
	maxsofar = stockPrices[len(stockPrices)-1]
	max := 0
	for i := len(stockPrices) - 2; i > 0; i-- {
		if maxsofar < stockPrices[i] {
			maxsofar = stockPrices[i]
		}
		ss := maxprofitbysellingByIndices[i-1] + maxsofar - stockPrices[i]
		if max < ss {
			max = ss
		}

	}
	return max
}
