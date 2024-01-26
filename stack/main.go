package main

type Stack[T any] []T

func NewStack[T any](cap int64) Stack[T] {
	return make(Stack[T], 0, cap)
}

func (s Stack[T]) Add(value T) Stack[T] {
	return append(s, value)
}

func (s Stack[T]) Pop() Stack[T] {
	return s[:len(s)-1]
}

func main() {}
