package main

import (
	"fmt"
	_ "os"
	_ "strconv"
)

type NativeDictionary[T any] struct {
	size      int
	slots     []string
	values    []T
	fillSlots []bool
	cap       int
}

// создание экземпляра словаря
func Init[T any](sz int) NativeDictionary[T] {
	nd := NativeDictionary[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	return nd
}

func (nd *NativeDictionary[T]) HashFun(value string) int {
	// всегда возвращает корректный индекс слота
	if nd.size == 0 {
		return -1
	}

	var indx int
	var sum byte
	for i, item := range value {
		sum += byte(item) * byte(i)
	}

	indx = int(sum) % nd.size

	return indx
}

func (nd *NativeDictionary[T]) IsKey(key string) bool {
	// возвращает true если ключ имеется
	if nd.size == 0 {
		return false
	}

	for _, keyFromSlots := range nd.slots {
		if keyFromSlots == key {
			return true
		}
	}
	return false
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {

	var result T

	if !nd.IsKey(key) {
		return result, fmt.Errorf("key is not in array")
	}

	hash := nd.HashFun(key)

	if hash == -1 {
		return result, fmt.Errorf("size is zero")
	}

	result = nd.values[hash]

	return result, nil
}

func (nd *NativeDictionary[T]) Put(key string, value T) {

	if nd.cap < nd.size {
		hash := nd.HashFun(key)

		if hash == -1 {
			return
		}

		nd.slots[hash] = key
		nd.values[hash] = value
	}

}
