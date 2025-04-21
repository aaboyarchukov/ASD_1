package main

import (
	"fmt"
	"math"
)

const NUMBER byte = 75

type NativeCache[T any] struct {
	size      int
	slots     []string
	fillSlots []bool
	hits      []int
	step      int
	values    []T
	cap       int
}

func Init[T any](sz int) NativeCache[T] {
	nc := NativeCache[T]{size: sz}
	nc.slots = make([]string, sz)
	nc.hits = make([]int, sz)
	nc.values = make([]T, sz)
	nc.fillSlots = make([]bool, sz)
	nc.step = 1
	return nc
}

func (nc *NativeCache[T]) HashFun(value string) int {
	if nc.size == 0 {
		return -1
	}

	var incx int
	var sum byte
	for i, item := range value {
		sum += byte(item) * (byte(i) + NUMBER)
	}

	incx = int(sum) % nc.size

	return incx
}

func (nc *NativeCache[T]) SeekSlot(value string) int {
	if nc.size == 0 {
		return -1
	}

	hash := nc.HashFun(value)

	if !nc.fillSlots[hash] {
		return hash
	}

	if nc.cap < nc.size {

		resultIndx, indx := hash, hash
		for nc.fillSlots[resultIndx] {
			indx += nc.step
			resultIndx = indx % nc.size
		}

		return resultIndx
	}

	if nc.cap >= nc.size {
		return nc.FindIndex(value)
	}

	return -1
}

func (nc *NativeCache[T]) MinHits() int {
	var resultIndx int
	min := math.MaxInt
	for indx := range nc.hits {
		tempElement := nc.hits[indx]
		if tempElement < min {
			min = tempElement
			resultIndx = indx
		}
	}

	return resultIndx
}

func (nc *NativeCache[T]) IsKey(key string) bool {
	if nc.size == 0 {
		return false
	}

	hash := nc.SeekSlot(key)

	return nc.fillSlots[hash] && nc.slots[hash] == key
}

func (nc *NativeCache[T]) Get(key string) (T, error) {

	var result T

	if !nc.IsKey(key) {
		return result, fmt.Errorf("key is not in array")
	}

	hash := nc.SeekSlot(key)

	if hash == -1 {
		return result, fmt.Errorf("size is zero")
	}

	result = nc.values[hash]
	nc.hits[hash] += 1

	return result, nil
}

func (nc *NativeCache[T]) Put(key string, value T) {
	if nc.cap >= nc.size {
		removedIndx := nc.MinHits()
		nc.fillSlots[removedIndx] = false
		nc.hits[removedIndx] = 0
		nc.slots[removedIndx] = ""
		nc.cap -= 1
	}

	hash := nc.SeekSlot(key)

	if !nc.fillSlots[hash] {
		nc.slots[hash] = key
		nc.values[hash] = value
		nc.fillSlots[hash] = true
		nc.cap += 1
	}

}

func (nc *NativeCache[T]) FindIndex(key string) int {
	for indx := range nc.slots {
		if nc.slots[indx] == key {
			return indx
		}
	}

	return -1
}
