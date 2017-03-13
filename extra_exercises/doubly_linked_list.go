// Implementation of a Doubly Linked List
/*
	Doubly Linked List is a variation of linked list in which navigation is both ways, either forward or backward.
	Terms used:

	Link/Node:	Each Node of a linked list can store a data called an element.
	Next:		Each Node of a linked list contains a link to the next link called Next.
	Prev:		Each Node of a linked list contains a link to the previous link called Prev.
	LinkedList:	A Linked List contains connection link to the first link called 'Head' and to the last link called 'Tail'.

*/

package main

import (
	"errors"
	"fmt"
)

type Node struct {
	val        int
	prev, next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Ahead() *Node {
	return l.head
}

func (n *Node) Next() *Node { // Returns the next node; We cannot return tail of the list cuz, tail is the last node and keeps on changing
	return n.next
}

func (l *List) Push(v int) {
	n := &Node{val: v}

	if l.head == nil {
		l.head = n // If there is no head, make the node 'head'
	} else {
		l.tail.next = n //Make the present node of the list i.e. 'l.tail', point to the new node
		n.prev = l.tail //Make the prev of new node point to previous node of the list (previously tail)
	}

	l.tail = n // The new node is now the tail of the list
	//At first entry, the first node is both head and tail of the node

}

var errEmpty error = errors.New("no remaining node in the list")

func (l *List) Pop() (v int, err error) {

	if l.head == nil {
		err = errEmpty
	} else {

		v = l.tail.val
		l.tail = l.tail.prev
		if l.tail == nil { //If there is no previous tail (i.e. only one node left was popped out) , make the head also nil
			l.head = nil
		}
	}

	return v, err

}

func main() {
	l := new(List)
	l.Push(23)
	l.Push(21)
	l.Push(4)

	for n := l.Ahead(); n != nil; n = n.Next() { // Assigns n, the head of the List and loops through the next nodes
		fmt.Println(n.val)
	}

	fmt.Println()

	for v, err := l.Pop(); err == nil; v, err = l.Pop() { //Loops through the list by popping out the last item or 'tail'
		fmt.Println(v)
	}
}
