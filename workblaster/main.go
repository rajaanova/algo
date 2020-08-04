package main

import (
	"fmt"
	"time"
)

func worker(job, res chan int) {

	for channelValue := range job {
		res <- channelValue * channelValue
		time.Sleep(2 * time.Second)
	}
	time.After(2 * time.Second)
}
func channelForWorkers(work chan int, numOfWorkder int, res chan int) {
	for i := 1; i < numOfWorkder; i++ {
		go worker(work, res)
	}
}

func responseCollector(res, term chan int) {
	a := make([]int, 0)
	for response := range res {
		a = append(a, response)
	}
	fmt.Println(a)
	<-term
}
func main() {
	//sliceOfReq := rand.Perm(100)
	//fmt.Println(sliceOfReq)
	//ch := make(chan int, 20)
	//res := make(chan int)
	//responseCollector(res, term)
	//go channelForWorkers(ch, 20, res)
	//for _, val := range sliceOfReq {
	//	ch <- val
	//}
	//close(ch)
	//<-term
	i := 13
	v := "outer"
	fmt.Println(v)
	{
		v := "inner"
		fmt.Println(v)
		{
			fmt.Println(v)
		}
	}
	{
		fmt.Println(v)
	}
	fmt.Println(v)
	if i%10 == 2 {
		return
	}
	defer fmt.Println("hello raj")
	x := time.NewTicker(2 * time.Second)
	for c := range x.C {
		fmt.Println(c)
	}
}
