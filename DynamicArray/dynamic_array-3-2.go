package main

import "fmt"

// task6
// Dynamic array with bank method

const ADD_TO_BANK = 3
const REMOVE_TO_BANK = 2
const PAYMENT = 1

// earn with operations: insert, copy values without reallocation
type BankDynArray[T any] struct {
	count    int
	capacity int
	account  int
	array    []T
}

func (da *BankDynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *BankDynArray[T]) MakeArray(sz int) {
	if sz < 16 {
		sz = 16
	}

	if sz == da.capacity {
		return
	}

	var arr = make([]T, sz)
	// copy values
	// for i := 0; i < len(da.array); i++ {
	// 	indx := i
	// 	arr[indx] = da.array[indx]
	// }
	// or
	count := copy(arr, da.array)
	da.account -= count * PAYMENT

	da.capacity = sz
	da.array = arr
}

// t = O(n), mem = O(l), where l is memory will have to reached our arrays
func (da *BankDynArray[T]) Insert(itm T, index int) error {
	if index > da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}

	if da.count == da.capacity && da.account >= da.count {
		da.MakeArray(da.capacity * 2)
	}

	if index == da.count {
		da.Append(itm)
		return nil
	}

	for i := da.count - 1; i >= index; i-- {
		da.array[i+1] = da.array[i]
		da.account += ADD_TO_BANK
		da.account -= PAYMENT
	}
	da.array[index+1] = da.array[index]
	da.array[index] = itm
	da.count++
	da.account += ADD_TO_BANK
	da.account -= PAYMENT

	return nil

}

// t = O(n + m), where m is count of memory allocation, mem = O(l),
// where l is memory will have to slice our array
func (da *BankDynArray[T]) Remove(index int) error {
	if da.count == 0 {
		return fmt.Errorf("bad index '%d'", index)
	}

	if index >= da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}

	for i := index + 1; i <= da.count; i++ {
		indx := i
		da.array[indx-1] = da.array[indx]
		da.account += REMOVE_TO_BANK
		da.account -= PAYMENT
	}
	da.count--

	difference := float64(da.count) / float64(da.capacity)
	var newCap float64
	changed := false

	if difference < 0.5 && da.capacity > 16 {
		changed = true
		newCap = float64(da.capacity) / 1.5
		da.capacity = int(newCap)
	}
	// for difference < 0.5 && da.capacity > 16 {
	// 	changed = true
	// 	newCap = float64(da.capacity) / 1.5
	// 	da.capacity = int(newCap)
	// 	difference = float64(da.count) / float64(da.capacity)
	// }

	if changed && da.account >= da.count {
		da.MakeArray(int(newCap))
	}

	return nil

}

func (da *BankDynArray[T]) Append(itm T) {
	if da.count == da.capacity && da.account >= da.count {
		da.MakeArray(da.capacity * 2)
	}

	da.array[da.count] = itm
	da.account += ADD_TO_BANK
	da.account -= PAYMENT
	da.count++
}

func (da *BankDynArray[T]) GetItem(index int) (T, error) {
	var result T
	if index >= da.count || index < 0 {
		return result, fmt.Errorf("index out of range")
	}

	result = da.array[index]
	return result, nil
}

func EqualArraysWithBank[T comparable](da1 *BankDynArray[T], da2 *BankDynArray[T]) bool {
	if da1.count != da2.count || da1.capacity != da2.capacity {
		return false
	}

	for i := 0; i < da1.count; i++ {
		idx := i
		if da1.array[idx] != da2.array[idx] {
			return false
		}
	}

	return true
}

// task 7
// Multidimensional dynamic array

type MultiDimArray struct {
	dimens  []int
	size    int
	storage []int
}

func InitMDA(dimens ...int) MultiDimArray {
	size := 1

	for _, item := range dimens {
		size *= item
	}

	return MultiDimArray{
		dimens:  dimens,
		size:    size,
		storage: make([]int, size),
	}
}

// t = O(n), where n = len(indxs), mem = O(1)
func (mda *MultiDimArray) GetOneDemIndex(indxs ...int) int {
	if len(indxs) > len(mda.dimens) {
		return -1
	}

	for indx, item := range indxs {
		if item > mda.dimens[indx]-1 {
			return -1
		}
	}

	result := 0
	for i, indx := range indxs {
		temp := 1
		for _, dem := range mda.dimens[i+1:] {
			temp *= dem
		}
		result += indx * temp
	}

	return result
}

// t = O(n), where n = len(indxs), mem = O(1)
func (mda *MultiDimArray) Get(indxs ...int) int {
	indx := mda.GetOneDemIndex(indxs...)
	if indx == -1 {
		return -1
	}

	return mda.storage[indx]
}

// t = O(n), where n = len(indxs), mem = O(1)
func (mda *MultiDimArray) Put(value int, indxs ...int) {
	indx := mda.GetOneDemIndex(indxs...)
	if indx == -1 {
		return
	}

	mda.storage[indx] = value
}
