package main

import (
	"fmt"
	_ "os"
)

// task 2
// Stack with head peek
type Node[T any] struct {
	next  *Node[T]
	value T
}
type StackFront[T any] struct {
	head *Node[T]
	size int
}

// t = O(1), mem = O(1)
func (st *StackFront[T]) Size() int {
	return st.size
}

// t = O(1), mem = O(1)
func (st *StackFront[T]) Peek() (T, error) {
	var result T
	if st.size == 0 {
		return result, fmt.Errorf("this element is not in stack")
	}

	result = st.head.value

	return result, nil
}

// t = O(1), mem = O(1)
func (st *StackFront[T]) Pop() (T, error) {
	var result T
	result, err_peek := st.Peek()
	if err_peek != nil {
		return result, err_peek
	}

	nextHead := st.head.next
	st.head.next = nil
	st.head = nextHead

	st.size--
	return result, nil
}

// t = O(1)*, *with amortization time, when needed allocation, mem = O(1)
func (st *StackFront[T]) Push(itm T) {
	node := &Node[T]{
		value: itm,
		next:  nil,
	}

	if st.size > 0 {
		node.next = st.head
	}

	st.head = node
	st.size++
}

func EqualStackFront[T comparable](st1 *StackFront[T], st2 *StackFront[T]) bool {
	if st1.size != st2.size {
		return false
	}

	tempNode_1 := st1.head
	tempNode_2 := st2.head

	for tempNode_1 != nil {
		if tempNode_1.value != tempNode_2.value {
			return false
		}

		tempNode_1 = tempNode_1.next
		tempNode_2 = tempNode_2.next
	}

	return true
}

func GetStackFront[T any](values []T) *StackFront[T] {
	var result StackFront[T]

	for _, item := range values {
		result.Push(item)
	}

	return &result
}

// task 3
// 1. Когда в стэке нечетное количество элементов:
// на n - 1 вызове pop() - будет ошибка о том, что стэк пуст
// -> цикл прервется и программа завершится с ошибкой
// 2. Когда в стэке четное (не равное 0) количество элементов:
// поскольку мы каждую итерацию цикла удаляем по 2 элемента, то цикл отработает корректно и завершится без ошибок
// -> в конце размер стэка будет равен 0
// 3. Когда в стэке количество элементов равно 0:
// цикл не будет выполнен, так как количество элементов равно 0

// task 4 and 5
// t = O(n), where n = len(sequence), mem = O(m), m = stack.size
func BracketsSequence(sequence string) bool {
	// set: (){}[]
	var stack Stack[rune]

	for _, item := range sequence {
		bracket := item

		if element, err := stack.Peek(); err == nil && (bracket-element == 2 || bracket-element == 1) {
			stack.Pop()
		} else {
			stack.Push(bracket)
		}
	}

	return stack.size == 0

}

// task 6 and 7

type MinStack struct {
	base []int
	size int
	min  []int
	sum  int
}

// t = O(1), mem = O(1)
func (st *MinStack) Size() int {
	return st.size
}

// t = O(1), mem = O(1)
func (st *MinStack) Peek() (int, error) {
	var result int
	if st.size == 0 {
		return result, fmt.Errorf("this element is not in stack")
	}

	result = st.base[st.size-1]

	return result, nil
}

// t = O(1), mem = O(1)
func (st *MinStack) Pop() (int, error) {
	var result int
	result, err_peek := st.Peek()
	if err_peek != nil {
		return result, err_peek
	}

	st.base = st.base[:st.size-1]
	st.min = st.min[:st.size-1]

	st.size--
	st.sum -= result

	return result, nil
}

// t = O(1)*, *with amortization time, when needed allocation, mem = O(1)
func (st *MinStack) Push(itm int) {
	st.base = append(st.base, itm)

	if st.size > 1 {
		min := st.GetMin()

		if itm < min {
			st.min = append(st.min, itm)
		} else {
			st.min = append(st.min, min)
		}

	} else {
		st.min = append(st.min, itm)
	}

	st.size++
	st.sum += itm
}

// t = O(1), mem = O(1)
func (st *MinStack) GetMin() int {
	return st.min[st.size-1]
}

// t = O(1), mem = O(1)
func (st *MinStack) GetMean() float32 {
	return float32(st.sum) / float32(st.size)
}

func GetMinStack(values []int) *MinStack {
	var result MinStack

	for _, item := range values {
		result.Push(item)
	}

	return &result
}

// task 8
// The postfix record of an expression
// t = O(n), where n = stack.size, mem = O(m), m = stack.size
func ResultPostfixRecord[T any](expression []any) (int, error) {
	var result int

	var (
		st1 Stack[int]
		st2 Stack[any]
	)

	// good without this func
	st2 = *GetStack(expression)

	for range st2.size {
		item, err_pop := st2.Pop()
		if err_pop != nil {
			return -1, err_pop
		}

		// with switch type
		switch resultItem := item.(type) {
		case int:
			st1.Push(resultItem)
		case string:
			if resultItem == "=" {
				resultPop, err_pop_result := st1.Pop()
				if err_pop_result != nil {
					return -1, err_pop_result
				}
				result = resultPop
				break
			}
			numberPop_1, err_pop_1 := st1.Pop()
			numberPop_2, err_pop_2 := st1.Pop()

			if err_pop_1 != nil || err_pop_2 != nil {
				return -1, err_pop_1
			}

			switch resultItem {
			case "+":
				st1.Push(numberPop_1 + numberPop_2)
			case "*":
				st1.Push(numberPop_1 * numberPop_2)
			}
		}

		// another example of switch type
		// if number, ok := item.(int); ok {
		// 	st1.Push(number)
		// } else {

		// 	if item == "=" {
		// 		resultPop, err_pop_result := st1.Pop()
		// 		if err_pop_result != nil {
		// 			return -1, err_pop_result
		// 		}
		// 		result = resultPop
		// 		break
		// 	}

		// 	numberPop_1, err_pop_1 := st1.Pop()
		// 	numberPop_2, err_pop_2 := st1.Pop()

		// 	if err_pop_1 != nil || err_pop_2 != nil {
		// 		return -1, err_pop_1
		// 	}

		// 	switch item {
		// 	case "+":
		// 		st1.Push(numberPop_1 + numberPop_2)
		// 	case "*":
		// 		st1.Push(numberPop_1 * numberPop_2)

		// 	}

		// }
	}

	return result, nil
}
