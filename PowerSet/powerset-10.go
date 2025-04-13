package main

import (
	"constraints"
	"os"
	// "fmt"
	"slices"
	"strconv"
)

type PowerSet[T constraints.Ordered] struct {
	cap   int
	slots []T
}

func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{
		cap:   0,
		slots: make([]T, 0, 20000),
	}
}

func (ps *PowerSet[T]) Size() int {
	return ps.cap
}

func (ps *PowerSet[T]) Put(value T) {
	if ps.cap < len(ps.slots) && !ps.Get(value) {
		ps.slots = append(ps.slots, value)
		ps.cap++
	}

	return
}

func (ps *PowerSet[T]) Get(value T) bool {
	return slices.Contains(ps.slots, value)
}

func (ps *PowerSet[T]) Remove(value T) bool {
	if ps.Get(value) {
		indx := slices.Index(ps.slots, value)
		ps.slots = slices.Delete(ps.slots, indx, indx+1)
		ps.cap--
		return true
	}

	return false
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	for _, val := range ps.slots {
		if set2.Get(val) {
			result.Put(val)
		}
	}
	return result
}

func (ps *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	for _, val := range ps.slots {
		result.Put(val)
	}

	for _, val := range set2.slots {
		result.Put(val)
	}

	return result
}

func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	for _, val := range ps.slots {
		if !set2.Get(val) {
			result.Put(val)
		}
	}
	return result
}

func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
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
