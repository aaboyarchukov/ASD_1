package main

import (
	_ "os"
	//	"fmt" включите если используете
)

type Stack[T any] struct {
}

func (st *Stack[T]) Size() int {
	return 0
}

func (st *Stack[T]) Peek() (T, error) {
	var result T
	// ...
	return result, nil
}

func (st *Stack[T]) Pop() (T, error) {
	var result T
	// ...
	return result, nil
}

func (st *Stack[T]) Push(itm T) {
}
