package main

import (
	"testing"
)

func TestAccount(t *testing.T) {
	da := &BankDynArray[any]{}
	da.Init()
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, item := range values {
		da.Append(item)
	}
	for _, item := range values {
		da.Append(item)
	}
	for _, item := range values {
		da.Append(item)
	}
	for _, item := range values {
		da.Append(item)
	}
	for _, item := range values {
		da.Append(item)
	}
	for _, item := range values {
		da.Append(item)
	}

	t.Log(da.account, da.count, da.capacity)

}

func TestMDAInit(t *testing.T) {
	tests := []struct {
		name   string
		dimens []int
		size   int
	}{
		{"Test1", []int{1, 2, 3}, 6},
		{"Test2", []int{0, 2, 3}, 0},
		{"Test2", []int{3, 3}, 9},
	}

	for _, test := range tests {
		mda := InitMDA(test.dimens...)

		if len(mda.storage) != test.size {
			t.Errorf("%s: failed init MDA", test.name)
		}

		if !EqualsDemns(test.dimens, mda.dimens) {
			t.Errorf("%s: failed init MDA, dimens are not equal", test.name)
		}
	}
}

func TestMDADemIndex(t *testing.T) {
	tests := []struct {
		name      string
		dimens    []int
		indexs    []int
		wantIndex int
	}{
		{"Test1", []int{1, 2, 3}, []int{0, 1, 0}, 3},
		{"Test2", []int{3, 3}, []int{1, 1}, 4},
		{"Test3", []int{3, 3}, []int{3, 3}, -1},
	}

	for _, test := range tests {
		mda := InitMDA(test.dimens...)

		if mda.GetOneDemIndex(test.indexs...) != test.wantIndex {
			t.Errorf("%s: failed get index", test.name)
		}
	}
}

func TestMDAGet(t *testing.T) {
	tests := []struct {
		name    string
		mda     MultiDimArray
		indexs  []int
		element int
	}{
		{"Test1", MultiDimArray{
			dimens:  []int{3, 3},
			size:    9,
			storage: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, []int{1, 1}, 5},

		{"Test2", MultiDimArray{
			dimens:  []int{3, 3},
			size:    9,
			storage: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, []int{3, 3}, -1},
	}

	for _, test := range tests {
		getValue := test.mda.Get(test.indexs...)

		if getValue != test.element {
			t.Errorf("%s: failed get element", test.name)
		}
	}

}
func TestMDAPut(t *testing.T) {
	tests := []struct {
		name    string
		mda     MultiDimArray
		indexs  []int
		element int
	}{
		{"Test1", MultiDimArray{
			dimens:  []int{3, 3},
			size:    9,
			storage: []int{1, 2, 3, 4, 0, 6, 7, 8, 9},
		}, []int{1, 1}, 5},
	}

	for _, test := range tests {
		test.mda.Put(test.element, test.indexs...)
		getValue := test.mda.Get(test.indexs...)

		if getValue != test.element {
			t.Errorf("%s: failed put element", test.name)
		}
	}
}
