package main

import (
	_ "os"
	//	"fmt" включите если используете
)

type Queue[T any] struct {
}

func (q *Queue[T]) Size() int {
	return 0
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	// ...
	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
}
