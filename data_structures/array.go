package data_structures

type Array[T any] struct {
	Data   []T
	Length int
}

// Sequence Interface Implemetaions
func (arr *Array[T]) Build(data ...T) {
	arr.Length = len(data)
	arr.Data = make([]T, len(data))
	for i, v := range data {
		arr.Data[i] = v
	}
}

func (arr *Array[T]) Len() int {
	return arr.Length
}

func (arr *Array[T]) GetAt(i int) T {
	return arr.Data[i]
}

func (arr *Array[T]) SetAt(i int, x T) {
	arr.Data[i] = x
}
