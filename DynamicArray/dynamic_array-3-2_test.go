package main

import (
	"fmt"
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

func TestMDA(t *testing.T) {
	mda := InitMDA(1, 2, 3)
	fmt.Println(mda.GetOneDemIndex(1, 1, 1))
}
