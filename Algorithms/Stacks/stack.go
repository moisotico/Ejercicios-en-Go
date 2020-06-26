package main

import (
	"fmt"
)

type Stack struct {
	slice []int
}

//Adds the integer to the top of the stack
func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

// Returns the topm item of the stack without removing it
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

// Returns and removes the top item of the stack
func (s *Stack) Pop() int {
	returnVal := s.Peek()
	s.slice = s.slice[0 : len(s.slice)-1]
	return returnVal
}

func main() {
	var s *Stack = new(Stack)
	s.Push(123)
	s.Push(321)
	s.Push(112)
	fmt.Println("Slice: ", s)
	fmt.Println("Popping element: ", s.Pop())
	fmt.Println("Popping element: ", s.Pop())
	fmt.Println("Slice: ", s)
	fmt.Println("Peeking element: ", s.Peek())
	fmt.Println("Slice: ", s)
}
