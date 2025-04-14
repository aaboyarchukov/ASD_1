package main

import (
	"testing"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		name     string
		set1     PowerSet[string]
		set2     PowerSet[string]
		isEquals bool
	}{
		{"Test1", PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, true},
		{"Test2", PowerSet[string]{
			slots: []string{"3", "2", "4"},
			cap:   3,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, false},
		{"Test3", PowerSet[string]{
			slots: []string{"a", "b", "c", "d"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, false},
		{"Test4", PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, true},
		{"Test5", PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, true},
	}

	for _, test := range tests {
		isEquals := test.set1.Equals(test.set2)

		if isEquals != test.isEquals {
			t.Errorf("%v: wrong equals sets", test.name)
		}
	}
}
func TestIsSubset(t *testing.T) {
	tests := []struct {
		name     string
		set1     PowerSet[string]
		set2     PowerSet[string]
		isSubSet bool
	}{
		{"Test1", PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, true},
		{"Test2", PowerSet[string]{
			slots: []string{"3", "2", "4"},
			cap:   3,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, false},
		{"Test3", PowerSet[string]{
			slots: []string{"a", "b", "c", "d"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, false},
		{"Test4", PowerSet[string]{
			slots: []string{"5", "6", "7", "8"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, false},
		{"Test5", PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, true},
	}

	for _, test := range tests {
		isSubSet := test.set1.IsSubset(test.set2)

		if isSubSet != test.isSubSet {
			t.Errorf("%v: wrong subsets sets", test.name)
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name      string
		set1      PowerSet[string]
		set2      PowerSet[string]
		neededSet PowerSet[string]
	}{
		{"Test1", PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}},
		{"Test2", PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "4"},
			cap:   2,
		}},
		{"Test3", PowerSet[string]{
			slots: []string{"a", "b", "c", "d"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"a", "b", "c", "d"},
			cap:   4,
		}},
		{"Test4", PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}},
	}

	for _, test := range tests {
		resultSet := test.set1.Difference(test.set2)

		if !resultSet.Equals(test.neededSet) {
			t.Errorf("%v: wrong difference sets", test.name)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name      string
		set1      *PowerSet[string]
		set2      PowerSet[string]
		neededSet PowerSet[string]
	}{
		{"Test1", &PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}},
		{"Test2", &PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"1", "3", "2", "4", "5", "6"},
			cap:   6,
		}},
		{"Test3", &PowerSet[string]{
			slots: []string{"a", "b", "c", "d"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"a", "b", "c", "d", "5", "1", "2", "6"},
			cap:   8,
		}},
		{"Test4", &PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}},
		{"Test5", &PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}},
	}

	for _, test := range tests {
		resultSet := test.set1.Union(test.set2)
		if !resultSet.Equals(test.neededSet) {
			t.Errorf("%v: wrong union sets", test.name)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name      string
		set1      *PowerSet[string]
		set2      PowerSet[string]
		neededSet PowerSet[string]
	}{
		{"Test1", &PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"3", "1", "2", "4"},
			cap:   4,
		}},
		{"Test2", &PowerSet[string]{
			slots: []string{"1", "3", "2", "4"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"1", "2"},
			cap:   2,
		}},
		{"Test3", &PowerSet[string]{
			slots: []string{"a", "b", "c", "d"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}},
		{"Test4", &PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6"},
			cap:   4,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}},
		{"Test5", &PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, PowerSet[string]{
			slots: []string{},
			cap:   0,
		}},
		{"Test6", &PowerSet[string]{
			slots: []string{"1", "3", "2", "4", "g", "7"},
			cap:   6,
		}, PowerSet[string]{
			slots: []string{"5", "1", "2", "6", "7", "a", "g"},
			cap:   7,
		}, PowerSet[string]{
			slots: []string{"1", "2", "g", "7"},
			cap:   4,
		}},
	}

	for _, test := range tests {
		resultSet := test.set1.Intersection(test.set2)
		if !resultSet.Equals(test.neededSet) {
			t.Errorf("%v: wrong intersection sets", test.name)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		set     *PowerSet[string]
		value   string
		isExist bool
	}{
		{"Test1", &PowerSet[string]{
			slots: []string{"1", "2"},
			cap:   2,
		}, "3", false},
		{"Test2", &PowerSet[string]{
			slots: []string{"1", "2", "3"},
			cap:   3,
		}, "3", true},
		{"Test3", &PowerSet[string]{
			slots: []string{},
			cap:   3,
		}, "3", false},
	}

	for _, test := range tests {
		isExist := test.set.Get(test.value)

		if isExist != test.isExist {
			t.Errorf("%v: wrong find value", test.name)
		}
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		name      string
		set       *PowerSet[string]
		value     string
		neededSet PowerSet[string]
	}{
		{"Test1", &PowerSet[string]{
			slots: []string{"1", "2"},
			cap:   2,
		}, "3", PowerSet[string]{
			slots: []string{"1", "2", "3"},
			cap:   3,
		}},
		{"Test2", &PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, "3", PowerSet[string]{
			slots: []string{"3"},
			cap:   1,
		}},
		{"Test3", &PowerSet[string]{
			slots: []string{"2", "1", "4", "3"},
			cap:   4,
		}, "3", PowerSet[string]{
			slots: []string{"2", "1", "4", "3"},
			cap:   4,
		}},
	}

	for _, test := range tests {
		test.set.Put(test.value)

		if !test.set.Equals(test.neededSet) {
			t.Errorf("%v: wrong put value", test.name)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		name       string
		set        *PowerSet[string]
		neededSize int
	}{
		{"Test1", GetPowerSet([]string{"1", "2", "3", "4", "5"}), 5},
		{"Test2", GetPowerSet([]string{}), 0},
		{"Test3", GetPowerSet([]string{"1"}), 1},
	}

	for _, test := range tests {
		if test.neededSize != test.set.cap {
			t.Errorf("%v: wrong sizing set", test.name)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		set       *PowerSet[string]
		value     string
		neededSet PowerSet[string]
		isRemove  bool
	}{
		{"Test1", &PowerSet[string]{
			slots: []string{"1", "2", "3"},
			cap:   3,
		}, "3", PowerSet[string]{
			slots: []string{"1", "2"},
			cap:   2,
		}, true},
		{"Test2", &PowerSet[string]{
			slots: []string{"1", "2", "3"},
			cap:   3,
		}, "4", PowerSet[string]{
			slots: []string{"1", "2", "3"},
			cap:   3,
		}, false},
		{"Test3", &PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, "4", PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, false},
		{"Test4", &PowerSet[string]{
			slots: []string{"4"},
			cap:   1,
		}, "4", PowerSet[string]{
			slots: []string{},
			cap:   0,
		}, true},
	}

	for _, test := range tests {
		isRemove := test.set.Remove(test.value)

		if isRemove != test.isRemove {
			t.Errorf("%v: wrong remove value", test.name)
		}

		if !test.set.Equals(test.neededSet) {
			t.Errorf("%v: wrong remove value", test.name)
		}
	}
}
