package main

import (
	"fmt"
	_ "os"
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
	var arr = make([]T, sz)
	//  копируем содержимое array в arr ...
	da.capacity = sz
	da.array = arr //
}

func (da *DynArray[T]) Insert(itm T, index int) error {
	return fmt.Errorf("bad index '%d'", index)
}

func (da *DynArray[T]) Remove(index int) error {
	return fmt.Errorf("bad index '%d'", index)
}

func (da *DynArray[T]) Append(itm T) {

}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T
	// ...
	return result, nil
}
