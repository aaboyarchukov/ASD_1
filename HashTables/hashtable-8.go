package main

import (
	"os"
	"strconv"
)

type HashTable struct {
	size  int
	step  int
	slots []string
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	var indx int
	var sum byte
	for i, item := range value {
		sum += byte(item) * byte(i)
	}

	indx = int(sum) % ht.size

	return indx
}

func (ht *HashTable) SeekSlot(value string) int {
	hash := ht.HashFun(value)

	if ht.slots[hash] == "" {
		return hash
	}

	for i := hash + 1; i != hash; i += ht.step {
		indx := i
		if indx >= ht.size {
			i = indx % ht.size
			indx = i
		}

		if ht.slots[indx] == "" {
			return indx
		}
	}

	return -1
}

func (ht *HashTable) Put(value string) int {
	indx := ht.SeekSlot(value)
	if indx != -1 {
		ht.slots[indx] = value
	}

	return indx
}

func (ht *HashTable) Find(value string) int {
	// return SeekSlot(value) or
	for indx, item := range ht.slots {
		if item == value {
			return indx
		}
	}

	return -1
}






