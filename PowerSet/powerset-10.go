package main

import (
	"constraints"
	"fmt"
	"os"
	"strconv"
)

type PowerSet[T constraints.Ordered] struct {
	cap       int
	slots     []T
	fillSlots []bool
	step      int
}

func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{
		cap:       0,
		slots:     make([]T, 20000),
		fillSlots: make([]bool, 20000),
		step:      3,
	}
}

func (ps *PowerSet[T]) SeekSlot(value T) int {

	hash := ps.HashFunc(value)

	if !ps.fillSlots[hash] {
		return hash
	}

	if ps.cap < len(ps.slots) {

		resultIndx, indx := hash, hash
		for ps.fillSlots[resultIndx] {
			indx += ps.step
			resultIndx = indx % len(ps.slots)
		}

		return resultIndx
	}

	return -1
}

func (ps *PowerSet[T]) HashFunc(value T) int {
	valueForHash := fmt.Sprintf("%v", value)

	var indx int
	var sum byte
	for i, item := range valueForHash {
		sum += byte(item) * byte(i)
	}

	indx = int(sum) % len(ps.slots)

	return indx
}

func (ps *PowerSet[T]) Size() int {
	return ps.cap
}

func (ps *PowerSet[T]) Put(value T) {
	if ps.cap < len(ps.slots) && !ps.Get(value) {
		hash := ps.SeekSlot(value)
		ps.slots[hash] = value
		ps.fillSlots[hash] = true
		ps.cap++
	}

	return
}

func (ps *PowerSet[T]) Get(value T) bool {
	hash := ps.SeekSlot(value)
	return ps.fillSlots[hash]
}

func (ps *PowerSet[T]) Remove(value T) bool {
	if ps.Get(value) {
		hash := ps.SeekSlot(value)
		ps.fillSlots[hash] = false
		return true
	}

	return false
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	return result
}

func (*PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	return result
}

func (*PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T]
	return result
}

func (*PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	return false
}





func (*PowerSet[T]) Equals(set2 PowerSet[T]) bool {
	return false
}
