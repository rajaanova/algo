package main

import "fmt"

type BurgerDecorator interface {
	getCost() int
}
type somflavor struct {
	b BurgerDecorator
}
type anothersomeflavor struct {
	b BurgerDecorator
}
type basiccost struct {
}

func NewBasiccost() BurgerDecorator {
	return &basiccost{}
}
func NewSomFlavor(a BurgerDecorator) BurgerDecorator {
	return &somflavor{b: a}
}
func (a *somflavor) getCost() int {
	return a.b.getCost() + 3
}
func NewAnotherSomeflavor(a BurgerDecorator) BurgerDecorator {
	return &anothersomeflavor{b: a}
}
func (a *anothersomeflavor) getCost() int {
	return a.b.getCost() + 2
}

func (b *basiccost) getCost() int {
	return 9
}
func main() {
	c := NewAnotherSomeflavor(NewSomFlavor(NewBasiccost())).getCost()
	fmt.Println(c)
}
