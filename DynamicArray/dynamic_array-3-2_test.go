package main

import "testing"

func TestAccount(t *testing.T) {
	da := &BankDynArray[int]{}
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
