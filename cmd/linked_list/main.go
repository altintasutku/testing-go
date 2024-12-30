package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length uint
}

func (ll *LinkedList) appendValue(value int) {
	newNode := Node{
		value: value,
		next:  nil,
	}

	ll.length++

	if ll.head == nil {
		ll.head = &newNode
		ll.tail = &newNode
		return
	}

	ll.tail.next = &newNode
	ll.tail = &newNode
}

func printLinkedList(list *LinkedList) {
	current := list.head

	for current != nil {
		fmt.Print(current.value, " ")

		current = current.next
	}
}

func main() {
	list := new(LinkedList)

	list.appendValue(1)
	list.appendValue(2)
	list.appendValue(3)
	list.appendValue(4)
	list.appendValue(10)
	list.appendValue(6)

	printLinkedList(list)
}
