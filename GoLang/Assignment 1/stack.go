package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack that holds a slice of integers.
type Stack struct {
	elements []int
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(element int) {
	// TODO: Implement this method
	//by using append it will add the element to the top of the stack
	s.elements = append(s.elements, element)
}

// Pop removes the element at the top of the stack and returns it.
// It returns an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
	// TODO: Implement this method
	//check if the stack is empty or not by using the method IsEmpty()
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	//index value should be less than the length of total elements
	index := len(s.elements) - 1
	//take a variable that will store the element with provided index
	value := s.elements[index]
	//it will remove element from the top of the stack
	s.elements = s.elements[:index]
	return value, nil
}

// Peek returns the element at the top of the stack without removing it.
// It returns an error if the stack is empty.
func (s *Stack) Peek() (int, error) {
	// TODO: Implement this method
	//check the stack is empty or not
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	//it will return the top of the element in the stack
	return s.elements[len(s.elements)-1], nil

}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *Stack) IsEmpty() bool {
	// TODO: Implement this method
	//check elements in stack is empty, if empty means it return true
	if len(s.elements) == 0 {
		return true
	}
	return false

}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	// TODO: Implement this method
	//simply we can get the elements by use length of the stack elements
	return len(s.elements)
}

func main() {
	stack := Stack{}

	// Example usage (You can modify this part to test your implementation)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack size:", stack.Size())

	topElement, err := stack.Peek()
	if err == nil {
		fmt.Println("Top element:", topElement)
	} else {
		fmt.Println(err)
	}

	for !stack.IsEmpty() {
		element, err := stack.Pop()
		if err == nil {
			fmt.Println("Popped element:", element)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Stack size after popping all elements:", stack.Size())
}
