package main

import "fmt"


type Queue struct {
	allMarks []int 
} 

// push function
func (q *Queue) Push(newMarks int) {

	q.allMarks = append(q.allMarks, newMarks)

}

// pop function
func (q *Queue) Pop() int {

	removedItem := q.allMarks[0];

	q.allMarks = q.allMarks[1:] 

	return removedItem

}


func main() {

	studentMarks := Queue{}

	fmt.Printf("Marks of students before adding any marks%v:\n", studentMarks)
	
	studentMarks.Push(90)
	studentMarks.Push(70)
	studentMarks.Push(80)
	studentMarks.Push(65)

	fmt.Printf("Marks of students before after adding four marks%v:\n", studentMarks)

	studentMarks.Pop()

	fmt.Printf("Popping out the marks present at first position %v:\n", studentMarks)

	studentMarks.Pop()

	fmt.Printf("Popping out the marks present at second position %v:\n", studentMarks)
	

}