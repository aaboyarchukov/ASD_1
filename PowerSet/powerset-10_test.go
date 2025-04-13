package main

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name  string
		set   PowerSet[int64]
		value int64
		want  PowerSet[int64]
	}{
		{"Test1", PowerSet[int64]{
			slots: []int64{1, 4, 2, 3},
			cap:   4,
		}, 4, PowerSet[int64]{
			slots: []int64{1, 2, 3},
			cap:   3,
		}},
	}

	for _, test := range tests {
		isRemove := test.set.Remove(test.value)
		if isRemove {
			fmt.Println(test.set)
			fmt.Println(test.want)
		}
	}
}
