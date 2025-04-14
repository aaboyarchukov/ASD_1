package main

import (
	"constraints"
	// "fmt"
	"os"
	"strconv"
)

type PowerSet[T constraints.Ordered] struct {
	cap   int
	slots []T
}

func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{
		cap:   0,
		slots: make([]T, 0),
	}
}

func (ps *PowerSet[T]) Size() int {
	return ps.cap
}

func (ps *PowerSet[T]) Put(value T) {
	if !ps.Get(value) {
		ps.slots = append(ps.slots, value)
		ps.cap++
	}
}

func (ps *PowerSet[T]) Index(value T) int {
	for indx, item := range ps.slots {
		if item == value {
			return indx
		}
	}

	return -1
}

func Delete[T constraints.Ordered](slice []T, indx int) []T {
	slice = append(slice[:indx], slice[indx+1:]...)
	return slice
}

func (ps *PowerSet[T]) Get(value T) bool {
	return ps.Index(value) >= 0
}

func (ps *PowerSet[T]) Remove(value T) bool {
	if ps.Get(value) {
		indx := ps.Index(value)
		ps.slots = Delete(ps.slots, indx)
		ps.cap--
		return true
	}

	return false
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	var result []T
	for _, val := range ps.slots {
		if set2.Get(val) {
			result = append(result, val)
		}
	}

	return PowerSet[T]{
		slots: result,
		cap:   len(result),
	}
}

func (ps *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	if ps.Equals(set2) {
		return set2
	}

	for _, item := range set2.slots {
		if !ps.Get(item) {
			ps.Put(item)
		}
	}

	return PowerSet[T]{
		slots: ps.slots,
		cap:   ps.cap,
	}

}

func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	var result []T
	for _, val := range ps.slots {
		if !set2.Get(val) {
			result = append(result, val)
		}
	}
	return PowerSet[T]{
		slots: result,
		cap:   len(result),
	}
}

func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	if set2.cap > ps.cap {
		return false
	}

	for _, val := range set2.slots {
		if !ps.Get(val) {
			return false
		}
	}

	return true
}

func (ps *PowerSet[T]) Equals(set2 PowerSet[T]) bool {
	if ps.cap != set2.cap {
		return false
	}

	for _, val := range ps.slots {
		if !set2.Get(val) {
			return false
		}
	}

	return true
}

func GetPowerSet[T constraints.Ordered](values []T) *PowerSet[T] {
	var result *PowerSet[T] = &PowerSet[T]{
		slots: make([]T, 0),
		cap:   0,
	}

	for _, item := range values {
		result.Put(item)
	}

	return result
}







