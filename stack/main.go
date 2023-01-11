package main

import "fmt"

type Stack struct {
	names []string
}

// push function
func (s *Stack) Push(newName string) {

	s.names = append(s.names, newName)

}

// pop function
func (s *Stack) Pop() string {

	removedName := s.names[len(s.names)-1]

	s.names = s.names[:len(s.names)-1]

	return removedName

}

func main() {

	// creating a new stack
	myNamesStack := Stack{}

	fmt.Printf("Initial Stack %v:\n", myNamesStack)

	myNamesStack.Push("NameOne")
	myNamesStack.Push("NameTwo")
	myNamesStack.Push("NameThree")
	myNamesStack.Push("NameFour")

	fmt.Printf("Stack after pushing 4 new names:%v\n", myNamesStack)

	myNamesStack.Pop()

	fmt.Printf("Stack after popping out the last element:%v\n", myNamesStack)

	myNamesStack.Pop()

	fmt.Printf("Stack after popping out the second last element:%v\n", myNamesStack)

}
