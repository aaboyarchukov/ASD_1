package main

import (
	"fmt"
	"os"
)

type Stack[T any] struct {
	base []T
	size int
}

func (st *Stack[T]) Size() int {
	return st.size
}

func (st *Stack[T]) Peek() (T, error) {
	var result T
	if st.size == 0 {
		return result, fmt.Errorf("this element is not in stack")
	}

	result = st.base[st.size-1]

	return result, nil
}

// t = O(1)
func (st *Stack[T]) Pop() (T, error) {
	var result T
	result, err_peek := st.Peek()
	if err_peek != nil {
		return result, err_peek
	}

	st.base = st.base[:st.size-1]
	st.size--

	return result, nil
}

// t = O(1)*, *with amortization time, when needed allocation
func (st *Stack[T]) Push(itm T) {
	st.base = append(st.base, itm)
	st.size++
}

func EqualStack[T comparable](st1 *Stack[T], st2 *Stack[T]) bool {
	if st1.size != st2.size {
		return false
	}

	for i := 0; i < st1.size; i++ {
		if st1.base[i] != st2.base[i] {
			return false
		}
	}

	return true
}

func GetStack[T any](values []T) *Stack[T] {
	var result Stack[T]

	for _, item := range values {
		result.Push(item)
	}

	return &result
}

