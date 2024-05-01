package data_structures

import (
	"fmt"
	"log"
)

/*
Eventhough Go lang's slice provides allocation of extra tail, and re-allocation with twice the capacity.
But anyways I will initialize a slice capacity with twice the length and will maintain in manually.
Also slices in Go don't reduce capacity when element are deleted.
So reducing the capacity will be implemented manually too .
*/

type DynamicArray[T any] struct {
	length int
	data   []T
}

// ----------- ContainerSequenceInterface implementation -----------------
func (arr *DynamicArray[T]) Build(vals ...T) {
	/*
		DynamicArray.length - number of elements in vals parameter
		DynamicArray.data   - slice with capacity of twice the DynamicArray.length
	*/
	arr.length = len(vals)
	arr.data = make([]T, len(vals)*2, len(vals)*2)
	copy(arr.data, vals)
}

func (arr *DynamicArray[T]) Len() int {
	return arr.length
}

// ----------- StaticSequenceInterface implementation --------------------
func (arr *DynamicArray[T]) GetAt(i int) (T, error) {
	if i < 0 || i >= arr.length {
		var emptyT T
		return emptyT, fmt.Errorf("index: %v is out of bound", i)
	}
	return arr.data[i], nil
}

func (arr *DynamicArray[T]) SetAt(i int, x T) bool {
	if i < 0 || i >= arr.length {
		log.Fatalf("index: %v is out of bound", i)
		return false
	}
	arr.data[i] = x
	return true
}

// ----------- DynamicSequenceInterface implementation -------------------
func (arr *DynamicArray[T]) InsertFirst(x T) {
	/*
		Element will be inserted at position 0 (first).
		Thus reallocate new address with increase slice for the data - O(n) time
	*/
	new_data := make([]T, 2*(arr.length+1), 2*(arr.length+1))
	new_data[0] = x
	for index, v := range arr.data {
		if index >= arr.length {
			break
		}
		new_data[index+1] = v
	}
	arr.data = new_data
	arr.length++
}

func (arr *DynamicArray[T]) DeleteFirst() (T, error) {
	if arr.length < 1 {
		var emptyT T
		return emptyT, fmt.Errorf("arr has no elements to delete")
	}

	deleted := arr.data[0]
	if 2*arr.length < cap(arr.data) {
		new_data := make([]T, 2*(arr.length-1), 2*(arr.length-1))
		copy(new_data, arr.data[1:arr.length])
		arr.data = new_data
		arr.length--
		return deleted, nil
	}

	arr.data = arr.data[1:]
	arr.length--
	return deleted, nil
}

func (arr *DynamicArray[T]) InsertLast(x T) {
	// Go automatically re-allocates and increases capacity if len reached cap
	if arr.length < cap(arr.data) {
		arr.data[arr.length-1] = x
	} else {
		suffixArr := make([]T, arr.length+1, arr.length+1)
		arr.data = append(arr.data, suffixArr...)
		arr.data[arr.length-1] = x
	}
	arr.length++
}

func (arr *DynamicArray[T]) DeleteLast() (T, error) {
	if len(arr.data) < 1 {
		var emptyT T
		return emptyT, fmt.Errorf("arr has no elements to delete")
	}

	var emptyT T
	deleted := arr.data[len(arr.data)-1]
	if 2*arr.length < cap(arr.data) {
		new_data := make([]T, 2*(arr.length-1), 2*(arr.length-1))
		copy(new_data, arr.data[:arr.length-1])
		arr.data = new_data
		arr.length--
		return deleted, nil
	}

	arr.data[len(arr.data)-1] = emptyT
	arr.length--
	return deleted, nil
}

// InsertAt(int, T) bool
func (arr *DynamicArray[T]) InsertAt(i int, x T) bool {
	if i < 0 || i >= arr.length {
		log.Fatalf("index: %v is out of bound", i)
		return false
	}
	arr.data = append(arr.data[:i], arr.data[i-1:]...)
	arr.data[i] = x
	arr.length++
	return true
}

func (arr *DynamicArray[T]) DeleteAt(i int) (T, error) {
	if i < 0 || i >= arr.length {
		var emptyT T
		return emptyT, fmt.Errorf("index: %v is out of bound", i)
	}
	deleted := arr.data[i]
	new_data := make([]T, 2*(arr.length-1), 2*(arr.length-1))
	for index, v := range arr.data {
		if index < i && index < arr.length {
			new_data[index] = v
		}
		if index > i && index < arr.length {
			new_data[index-1] = v
		}
	}
	arr.data = new_data
	arr.length--
	return deleted, nil
}

func (arr *DynamicArray[T]) String() string {
	return fmt.Sprint(arr.data[:arr.length])
}

func (arr *DynamicArray[T]) Iterator() *DynamicArrayIterator[T] {
	return &DynamicArrayIterator[T]{
		arr:          arr,
		currentIndex: 0,
	}
}

func (arr *DynamicArray[T]) GetData() []T {
	data := arr.data[:]
	return data
}

type DynamicArrayIterator[T any] struct {
	arr          *DynamicArray[T]
	currentIndex int
}

func (it *DynamicArrayIterator[T]) HasNext() bool {
	return it.currentIndex < it.arr.length
}

func (it *DynamicArrayIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		return *new(T), fmt.Errorf("Iterator: No more elements")
	}
	val := it.arr.data[it.currentIndex]
	it.currentIndex++
	return val, nil
}
