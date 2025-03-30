package main

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		step  int
		slots []string
	}{
		{"Test1", 19, 3, make([]string, 19)},
		{"Test2", 0, 0, make([]string, 0)},
		{"Test3", 1, 1, make([]string, 1)},
	}

	for _, test := range tests {
		hashTable := Init(test.size, test.step)

		if hashTable.size != test.size {
			t.Errorf("Failed %s: wrong size", test.name)
		}

		if len(hashTable.slots) != len(test.slots) {
			t.Errorf("Failed %s: wrong slots size", test.name)
		}

		if hashTable.step != test.step {
			t.Errorf("Failed %s: wrong step size", test.name)
		}
	}
}

func TestHashFunc(t *testing.T) {
	tests := []struct {
		name       string
		inputTable HashTable
		value      string
		wantIndx   int
	}{
		{"Test1", Init(19, 3), "abc", 2},
		{"Test2", Init(19, 3), "bac", 1},
		{"Test3", Init(19, 3), "abcd", 8},
	}

	for _, test := range tests {
		hash := test.inputTable.HashFun(test.value)

		if hash != test.wantIndx {
			t.Errorf("Failed %s: wrong hashFunc work", test.name)
		}
	}
}

func TestSeekSlot(t *testing.T) {
	tests := []struct {
		name       string
		inputTable HashTable
		value      string
		expcetIndx int
	}{
		{"Test1", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"пустая строка",
				"   пробелы   ",
				"EOF",
				"I ❤️ Open Source",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
				"🎲 Рандом",
				"Last one!",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			size: 19,
			step: 3,
			cap:  19,
		}, "testValue", -1},

		{"Test2", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"",
				"   пробелы   ",
				"EOF",
				"I ❤️ Open Source",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
				"🎲 Рандом",
				"Last one!",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				false,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			size: 19,
			step: 1,
			cap:  18,
		}, "testValue", 10},

		{"Test3", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"",
				"",
				"EOF",
				"Wrong",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				false,
				false,
				true,
				true,
				true,
				true,
				true,
			},
			size: 17,
			step: 3,
			cap:  15,
		}, "testValue", 10},
		{"Test4", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"",
				"",
				"",
				"",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				false,
				false,
				false,
				false,
				true,
				true,
				true,
			},
			size: 17,
			cap:  13,
			step: 3,
		}, "testValue", 12},
		{"Test4", HashTable{
			slots: []string{
				"g",
				"e",
				"f",
				"s",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
			},
			size: 4,
			cap:  4,
			step: 3,
		}, "testValue", -1},
		{"Test5", HashTable{
			slots: []string{
				"f",
				"our",
				"four",
			},
			fillSlots: []bool{
				true,
				true,
				true,
			},
			size: 3,
			cap:  3,
			step: 1,
		}, "testValue", -1},
		{"Test6", HashTable{
			slots: []string{
				"",
				"",
				"p",
				"sdsd",
			},
			fillSlots: []bool{
				true,
				false,
				true,
				true,
			},
			size: 4,
			cap:  0,
			step: 3,
		}, "testValue", 1},
	}

	for _, test := range tests {
		slot := test.inputTable.SeekSlot(test.value)
		if slot != test.expcetIndx {
			t.Errorf("Failed %s: wrong finding slot", test.name)
			fmt.Printf("Expected: %v, get: %v\n", test.expcetIndx, slot)
		}
	}
}

func TestPut(t *testing.T) {
	// сделать проверку на то, что нет свободных слотов
	tests := []struct {
		name       string
		value      string
		inputTable HashTable
		wantIndx   int
	}{
		{"Test1", "one", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"",
				"",
				"",
				"",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				false,
				false,
				false,
				false,
				true,
				true,
				true,
			},
			size: 17,
			cap:  13,
			step: 3,
		}, 11},
		{"Test2", "one", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"abc",
				"casbf",
				"d",
				"popopop",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			size: 17,
			cap:  17,
			step: 3,
		}, -1},
		{"Test3", "two", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"",
				"plplp",
				"qqq",
				"asdadas",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				false,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			size: 17,
			cap:  16,
			step: 3,
		}, 12},
		{"Test4", "three", HashTable{
			slots: []string{
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				``,
				"",
			},
			fillSlots: []bool{
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			size: 17,
			cap:  0,
			step: 3,
		}, 15},
		{"Test5", "four", HashTable{
			slots: []string{
				"",
				"",
				"",
				"",
				"",
				"four",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				``,
				"",
			},
			fillSlots: []bool{
				false,
				false,
				false,
				false,
				false,
				true,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			size: 17,
			cap:  1,
			step: 3,
		}, 5},
		{"Test6", "four", HashTable{
			slots:     []string{},
			fillSlots: []bool{},
			size:      0,
			cap:       0,
			step:      3,
		}, -1},
		{"Test7", "four", HashTable{
			slots: []string{
				"",
			},
			fillSlots: []bool{
				false,
			},
			size: 1,
			cap:  0,
			step: 3,
		}, 0},
		{"Test8", "four", HashTable{
			slots: []string{
				"four",
			},
			fillSlots: []bool{
				true,
			},
			size: 1,
			cap:  1,
			step: 3,
		}, 0},
	}

	for _, test := range tests {
		indx := test.inputTable.Put(test.value)

		if indx != test.wantIndx {
			t.Errorf("Failed %s: wrong index", test.name)
		}

		if indx != -1 && test.inputTable.slots[indx] != test.value {
			t.Errorf("Failed %s: didnt put it to table", test.name)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name       string
		inputTable HashTable
		value      string
		expectIndx int
	}{
		{"Test1", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"abc",
				"casbf",
				"d",
				"popopop",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
				"break",
				"Line",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			size: 19,
			cap:  19,
			step: 3,
		}, "popopop", 13},
		{"Test2", HashTable{
			slots: []string{
				"Hello, 世界",
				"Привет, Go!",
				"Lorem ipsum",
				"12345",
				"🍕 Pizza",
				"Тестовая строка",
				"Random data",
				"Go is awesome",
				"こんにちは",
				"!@#$%^&*()",
				"abc",
				"casbf",
				"d",
				"popopop",
				"Строка с \"кавычками\"",
				`Backtick \n raw string`,
				"Line\nbreak",
				"break",
				"Line",
			},
			fillSlots: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			size: 19,
			cap:  19,
			step: 3,
		}, "p", -1},
		{"Test3", HashTable{
			slots:     []string{},
			fillSlots: []bool{},
			size:      0,
			cap:       0,
			step:      3,
		}, "p", -1},
	}

	for _, test := range tests {
		indx := test.inputTable.Find(test.value)

		if indx != test.expectIndx {
			t.Errorf("Failed %s: wrong index", test.name)
		}

		if indx != -1 && test.inputTable.slots[indx] != test.value {
			t.Errorf("Failed %s: is not in table", test.name)
		}
	}
}

func TestFillSeekFind(t *testing.T) {

}
