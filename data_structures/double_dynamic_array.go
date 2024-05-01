package data_structures

import (
	"fmt"
	"log"
)

type DoubleDynamicArray[T any] struct {
	firstIndex int
	length     int
	data       []T
}

// ----------- ContainerSequenceInterface implementation -----------------
func (arr *DoubleDynamicArray[T]) Build(vals ...T) {
	arr.firstIndex = len(vals)
	arr.length = len(vals)
	arr.data = make([]T, arr.length*3, arr.length*3)
	for i, v := range vals {
		arr.data[arr.firstIndex+i] = v
	}
}

func (arr *DoubleDynamicArray[T]) Len() int {
	return arr.length
}

// ----------- StaticSequenceInterface implementation --------------------
func (arr *DoubleDynamicArray[T]) GetAt(i int) (T, error) {
	if i < 0 || arr.firstIndex+i > arr.firstIndex+arr.length-1 {
		var emptyT T
		return emptyT, fmt.Errorf("index: %v is out of bound", i)
	}
	return arr.data[arr.firstIndex+i], nil
}

func (arr *DoubleDynamicArray[T]) SetAt(i int, value T) bool {
	if i < 0 || arr.firstIndex+i > arr.firstIndex+arr.length-1 {
		log.Fatalf("index: %v is out of bound", i)
		return false
	}
	arr.data[arr.firstIndex+i] = value
	return true
}

// ----------- DynamicSequenceInterface implementation -------------------
func (arr *DoubleDynamicArray[T]) InsertFirst(value T) {
	if arr.firstIndex < 1 {
		newArr := make([]T, arr.length+1)
		arr.firstIndex = arr.length + 1
		arr.data = append(newArr, arr.data...)
	}
	arr.firstIndex--
	arr.length++
	arr.data[arr.firstIndex] = value
}

func (arr *DoubleDynamicArray[T]) DeleteFirst() (T, error) {
	var emptyT T
	if arr.length < 1 {
		return emptyT, fmt.Errorf("arr has no elements to delete")
	}

	deleted := arr.data[arr.firstIndex]
	arr.data[arr.firstIndex] = emptyT
	arr.length--
	arr.firstIndex++

	if arr.firstIndex/2 > arr.length {
		oldArray := arr.data[arr.firstIndex:]
		prefixArray := make([]T, arr.length, arr.length)
		arr.data = append(prefixArray, oldArray...)
		arr.firstIndex = arr.length
	}

	return deleted, nil
}

func (arr *DoubleDynamicArray[T]) InsertLast(value T) {
	if arr.firstIndex+arr.length >= cap(arr.data) {
		suffixArr := make([]T, arr.length, arr.length)
		arr.data = append(arr.data, suffixArr...)
	}
	arr.data[arr.firstIndex+arr.length] = value
	arr.length++
}

func (arr *DoubleDynamicArray[T]) DeleteLast() (T, error) {
	var emptyT T
	if arr.length < 1 {
		return emptyT, fmt.Errorf("arr has no elements to delete")
	}

	deleted := arr.data[arr.firstIndex+arr.length-1]
	arr.data[arr.firstIndex+arr.length-1] = emptyT
	arr.length--

	suffixLength := cap(arr.data) - (arr.firstIndex + arr.length)
	if suffixLength/2 > arr.length {
		suffixArr := make([]T, arr.length, arr.length)
		arr.data = append(arr.data[:arr.firstIndex+arr.length], suffixArr...)
	}
	return deleted, nil
}

func (arr *DoubleDynamicArray[T]) InsertAt(i int, value T) bool {
	if i < 0 || arr.firstIndex+i > arr.firstIndex+arr.length-1 {
		log.Printf("index: %v is out of bound", i)
		return false
	}
	leftSide := arr.data[:arr.firstIndex+i]
	rightSide := arr.data[arr.firstIndex+i-1:]
	arr.data = append(leftSide, rightSide...)
	arr.data[arr.firstIndex+i] = value
	arr.length++
	return true
}

func (arr *DoubleDynamicArray[T]) DeleteAt(i int) (T, error) {
	if i < 0 || arr.firstIndex+i > arr.firstIndex+arr.length-1 {
		var emptyT T
		return emptyT, fmt.Errorf("index: %v is out of bound", i)
	}

	deleted := arr.data[arr.firstIndex+i]
	if arr.firstIndex+i+1 < cap(arr.data) {
		arr.data = append(arr.data[:arr.firstIndex+i], arr.data[arr.firstIndex+i+1:]...)
	} else {
		arr.data = arr.data[:arr.firstIndex+i]
	}
	arr.length--

	if arr.firstIndex/2 > arr.length {
		oldArray := arr.data[arr.firstIndex:]
		prefixArray := make([]T, arr.length, arr.length)
		arr.data = append(prefixArray, oldArray...)
		arr.firstIndex = arr.length
	}

	suffixLength := cap(arr.data) - (arr.firstIndex + arr.length)
	if suffixLength/2 > arr.length {
		suffixArr := make([]T, arr.length, arr.length)
		arr.data = append(arr.data[:arr.firstIndex+arr.length], suffixArr...)
	}

	return deleted, nil
}

func (arr *DoubleDynamicArray[T]) String() string {
	return fmt.Sprint(arr.data[arr.firstIndex : arr.firstIndex+arr.length])
}

func (arr *DoubleDynamicArray[T]) Iterator() *DoubleDynamicArrayIterator[T] {
	return &DoubleDynamicArrayIterator[T]{
		arr:          arr,
		currentIndex: arr.firstIndex,
	}
}

func (arr *DoubleDynamicArray[T]) GetData() []T {
	data := arr.data[:]
	return data
}

type DoubleDynamicArrayIterator[T any] struct {
	arr          *DoubleDynamicArray[T]
	currentIndex int
}

func (it *DoubleDynamicArrayIterator[T]) HasNext() bool {
	return it.currentIndex < it.arr.firstIndex+it.arr.length
}

func (it *DoubleDynamicArrayIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		return *new(T), fmt.Errorf("Iterator: No more elements")
	}
	val := it.arr.data[it.currentIndex]
	it.currentIndex++
	return val, nil
}
