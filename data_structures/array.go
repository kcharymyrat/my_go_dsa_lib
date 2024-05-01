package data_structures

import (
	"fmt"
	"log"
)

type Array[T any] struct {
	data   []T
	length int
}

// Sequence Interface Implemetaions
func (arr *Array[T]) Build(vals ...T) {
	arr.length = len(vals)
	arr.data = make([]T, len(vals))
	copy(arr.data, vals)
}

func (arr *Array[T]) Len() int {
	return arr.length
}

func (arr *Array[T]) GetAt(i int) (T, error) {
	if i < 0 || i >= len(arr.data) {
		var emptyT T
		return emptyT, fmt.Errorf("index: %v is out of range", i)
	}
	return arr.data[i], nil
}

func (arr *Array[T]) SetAt(i int, x T) bool {
	if i < 0 || i >= len(arr.data) {
		log.Fatalf("index: %v is out of range", i)
		return false
	}
	arr.data[i] = x
	return true
}
