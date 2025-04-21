package main

import (
	"fmt"
	"testing"
)

func TestHits(t *testing.T) {
	tests := []struct {
		name     string
		nc       NativeCache[int]
		pairs    map[string]int
		key      string
		wantHits int
	}{
		{"Test1", Init[int](5), map[string]int{
			"1":   1,
			"2":   2,
			"3":   3,
			"key": 4,
			"4":   5,
		}, "key", 5},
	}

	for _, test := range tests {
		for key, value := range test.pairs {
			test.nc.Put(key, value)
		}

		for range test.wantHits {
			test.nc.Get(test.key)
		}

		indxHits := test.nc.FindIndex(test.key)

		if indxHits == -1 {
			t.Errorf("test %v: wrong indx", test.name)
			return
		}

		if test.nc.hits[indxHits] != test.wantHits {
			t.Errorf("test %v: wrong number of hits", test.name)
		}

	}
}
func TestWorkLoad(t *testing.T) {
	tests := []struct {
		name             string
		nc               NativeCache[int]
		pairs            map[string]int
		hits             []int
		key              string
		value            int
		indexWithMinHits int
	}{
		{"Test1", Init[int](5), map[string]int{
			"1": 1,
			"2": 2,
			"3": 3,
			"5": 5,
			"4": 4,
		}, []int{2, 1, 0, 3, 1}, "key", 10, 2},
	}

	for _, test := range tests {
		test.nc.hits = test.hits

		for key, value := range test.pairs {
			test.nc.Put(key, value)
		}

		test.nc.Put(test.key, test.value)
		if test.nc.slots[test.indexWithMinHits] != test.key ||
			!test.nc.fillSlots[test.indexWithMinHits] ||
			test.nc.values[test.indexWithMinHits] != test.value {
			t.Errorf("test %v: wrong eliminate key", test.name)
			fmt.Println(test.nc)
		}

	}
}
