package main

import (
	"os"
	"strconv"
)

type HashTable struct {
	size      int
	step      int
	slots     []string
	fillSlots []bool
	cap       int
}

func Init(sz int, stp int) HashTable {
	if stp == 0 {
		stp = 1
	}

	ht := HashTable{size: sz, step: stp, slots: nil, fillSlots: nil, cap: 0}
	ht.slots, ht.fillSlots = make([]string, sz), make([]bool, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	if ht.size == 0 {
		return -1
	}

	var indx int
	var sum byte
	for i, item := range value {
		sum += byte(item) * byte(i)
	}

	indx = int(sum) % ht.size

	return indx
}

func (ht *HashTable) SeekSlot(value string) int {
	if ht.size == 0 {
		return -1
	}

	hash := ht.HashFun(value)

	if !ht.fillSlots[hash] {
		return hash
	}

	if ht.cap < ht.size {

		resultIndx, indx := hash, hash
		for ht.fillSlots[resultIndx] {
			indx += ht.step
			resultIndx = indx % ht.size
		}

		return resultIndx
	}

	return -1
}

func (ht *HashTable) Put(value string) int {
	if ht.size == 0 {
		return -1
	}

	hash := ht.HashFun(value)
	if ht.slots[hash] == value {
		return hash
	}

	findIndx := ht.Find(value)
	if findIndx != -1 {
		return findIndx
	}

	indx := ht.SeekSlot(value)
	if indx != -1 {
		ht.slots[indx] = value
		ht.fillSlots[indx] = true
		ht.cap++
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







