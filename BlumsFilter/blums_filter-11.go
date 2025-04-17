package main

import (
	_ "os"
)

type BloomFilter struct {
	filter_len int
	filter     int64
}

const CONST_17 int = 17

func (bf *BloomFilter) Hash1(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum += code * CONST_17
	}
	sum %= bf.filter_len
	return sum
}

const CONST_223 int = 223

func (bf *BloomFilter) Hash2(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum += code * CONST_223
	}
	sum %= bf.filter_len
	return sum
}

func (bf *BloomFilter) Add(s string) {
	hash1 := bf.Hash1(s)
	hash2 := bf.Hash2(s)
	bf.filter |= 1 << hash1
	bf.filter |= 1 << hash2
}

func (bf *BloomFilter) IsValue(s string) bool {
	var mask int64
	mask |= 1 << bf.Hash1(s)
	mask |= 1 << bf.Hash2(s)

	if mask == bf.filter&mask {
		return true
	}

	return false
}
