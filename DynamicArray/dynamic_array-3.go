package main

import (
	"fmt"
	"os"
)

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *DynArray[T]) MakeArray(sz int) {
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
	copy(arr, da.array)
	da.capacity = sz
	da.array = arr
}

// t = O(n), mem = O(l), where l is memory will have to reached our arrays
func (da *DynArray[T]) Insert(itm T, index int) error {
	if index > da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}

	if da.count == da.capacity {
		da.MakeArray(da.capacity * 2)
	}

	if index == da.count {
		da.Append(itm)
		return nil
	}

	for i := da.count - 1; i >= index; i-- {
		da.array[i+1] = da.array[i]
	}
	da.array[index+1] = da.array[index]
	da.array[index] = itm
	da.count++
	return nil

}

// t = O(n + m), where m is count of memory allocation, mem = O(l),
// where l is memory will have to slice our array
func (da *DynArray[T]) Remove(index int) error {
	if da.count == 0 {
		return nil
	}

	if index >= da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}

	for i := index + 1; i <= da.count; i++ {
		indx := i
		da.array[indx-1] = da.array[indx]
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

	if changed {
		da.MakeArray(int(newCap))
	}

	return nil

}

func (da *DynArray[T]) Append(itm T) {
	if da.count == da.capacity {
		da.MakeArray(da.capacity * 2)
	}

	da.array[da.count] = itm
	da.count++
}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T
	if index >= da.count || index < 0 {
		return result, fmt.Errorf("index out of range")
	}

	result = da.array[index]
	return result, nil
}

func EqualArrays[T comparable](da1 *DynArray[T], da2 *DynArray[T]) bool {
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

// func GenerateArray(count, cap int) []int {

// 	return []int{}
// }

