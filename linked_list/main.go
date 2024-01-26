package main

import (
	"log"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Length int64
	Head   *Node[T]
	Tail   *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
	}
}

func (ll *LinkedList[T]) PushBack(value T) {
	defer func() {
		ll.Length++
	}()

	if ll.Head == nil {
		ll.Head = &Node[T]{
			Value: value,
		}
		ll.Tail = ll.Head
		return
	}

	ll.Head.Next = &Node[T]{
		Value: value,
	}
	ll.Head = ll.Head.Next
	return
}

func (ll *LinkedList[T]) PopBack() T {
	res := ll.Head.Value
	ll.Head = &Node[T]{}
	ll.Length--
	return res
}

func (ll *LinkedList[T]) PrintList() {
	start := ll.Tail
	for start != nil {
		log.Println(start.Value)
		start = start.Next
	}
}

func main() {
	list := NewLinkedList[any]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack("huy v rot")
	list.PrintList()
	log.Println(list.PopBack())
	list.PrintList()
}
