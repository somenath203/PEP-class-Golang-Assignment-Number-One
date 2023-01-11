package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}


// Method to create a linked list
func (l *LinkedList) CreateNode(n *Node) {

	temp := l.head

	l.head = n

	l.head.next = temp

	l.length++

}

// Method to print the created linked list
func (l *LinkedList) PrintLinkedList() {

	head := l.head

	for l.length != 0 {
		fmt.Printf("%d", head.data)
		head = head.next
		l.length--

		fmt.Printf("\n")

	}

}

func main() {

	firstLinkedList := LinkedList{}

	node_one := &Node{data: 20}
	node_two := &Node{data: 15}
	node_three := &Node{data: 45}
	node_four := &Node{data: 37}
	node_five := &Node{data: -45}
	node_six := &Node{data: 102}

	firstLinkedList.CreateNode(node_one)
	firstLinkedList.CreateNode(node_two)
	firstLinkedList.CreateNode(node_three)
	firstLinkedList.CreateNode(node_four)
	firstLinkedList.CreateNode(node_five)
	firstLinkedList.CreateNode(node_six)

	firstLinkedList.PrintLinkedList()

}
