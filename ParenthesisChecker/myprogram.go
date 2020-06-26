package main

import (
	"errors"
	"fmt"
)

//Stack
type Stack struct {
	slice []string
}

func (s *Stack) Push(i string) {
	s.slice = append(s.slice, i)
}

func (s *Stack) Peek() string {
	return s.slice[len(s.slice)-1]
}

func (s *Stack) Pop() string {
	returnVal := s.Peek()
	s.slice = s.slice[0 : len(s.slice)-1]
	return returnVal
}

func main() {
	var s *Stack = new(Stack)

	//String to send
	str := "f{e(d)]"
	fmt.Println("String to send: ", str)

	checkBalance := balancedString(s, str)
	if checkBalance {
		fmt.Println("The parenteses (), Brackets [], and Curly brackets are balanced")
	} else {
		err := errors.New("\nERROR: Unbalanced elements")
		fmt.Println(err)
	}

}

func balancedString(s *Stack, str string) bool {
	var checkFalse bool
	// Must check
	for i, c := range str {
		checkBalance(i, s, string(c))
		//fmt.Println("Caracter read of string: ", string(c))
		checkFalse = checkRemaining(s)
	}
	return checkFalse
}

func checkBalance(i int, s *Stack, str string) {
	switch str {
	case "(":
		s.Push(")")
	case "[":
		s.Push("]")
	case "{":
		s.Push("}")
	case ")":
		if s.Peek() == str {
			s.Pop()
			break
		}
		fmt.Printf("Error on char %d of string \nValue read: ) Value expected: %s", i+1, s.Peek())
	case "]":
		if s.Peek() == str {
			s.Pop()
			break
		}
		fmt.Printf("Error on char %d of string \nValue read: ] Value expected: %s", i+1, s.Peek())

	case "}":
		if s.Peek() == str {
			s.Pop()
			break
		}
		fmt.Printf("Error on char %d of string \nValue read: } Value expected: %s", i+1, s.Peek())
	}
}

func checkRemaining(s *Stack) bool {
	if len(s.slice) != 0 {
		return false
	}
	return true
}
