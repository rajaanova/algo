package main

import "fmt"

func main() {
	his := []int{2, 6, 1}

	fmt.Println(getLargest(his))
}

type stack struct {
	somestack []int
	top       int
}

func NewStack(len int) *stack {
	return &stack{somestack: make([]int, len), top: -1}
}
func (a *stack) peek() int {
	return a.somestack[a.top]
}

func (a *stack) pop() int {
	top := a.top
	a.top--
	return a.somestack[top]
}
func (a *stack) push(i int) {
	a.top++
	a.somestack[a.top] = i
}
func (a *stack) isEmpty() bool {
	if a.top == -1 {
		return true
	}
	return false
}
func getLargest(his []int) int {
	sta := NewStack(len(his))
	i := 0
	maxArea := 0
	area := 0
	for i < len(his) {
		if sta.isEmpty() || his[sta.peek()] <= his[i] {
			sta.push(i)
			i++
		} else {
			currentMaxIndex := sta.pop()
			if sta.isEmpty() {
				area = his[currentMaxIndex] * (i)
			} else {
				area = his[currentMaxIndex] * (i - sta.peek() - 1)
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}
	for !sta.isEmpty() {
		currentMaxIndex := sta.pop()
		if sta.isEmpty() {
			area = his[currentMaxIndex] * (i)
		} else {
			b := i - sta.peek() - 1
			area = his[currentMaxIndex] * (b)
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
