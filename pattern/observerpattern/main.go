package main

import "fmt"

func main() {
	var s Subject = NewWeatherStation()
	NewWeatherObserver(s)
	s.notify(2, 3, 4)
}

type Subject interface {
	notify(int, int, int)
	addObserver(observer)
}
type observer interface {
	update(int, int, int)
}
type weatherObservers struct {
	xyz Subject
}
type weatherStation struct {
	observers []observer
}

func (b *weatherStation) addObserver(a observer) {
	b.observers = append(b.observers, a)
}

func (b *weatherStation) notify(i, j, k int) {
	for _, obs := range b.observers {
		obs.update(i, j, k)
	}
}
func NewWeatherStation() *weatherStation {
	return &weatherStation{}
}

func NewWeatherObserver(sub Subject) *weatherObservers {
	obs := &weatherObservers{xyz: sub}
	sub.addObserver(obs)
	return obs
}

func (d *weatherObservers) update(a, b, c int) {
	//fmt.Println()
	fmt.Println(a, b, c)
}
