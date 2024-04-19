package ds_interfaces

type StaticSequenceInterface[T any] interface {
	ContainerInterface[T]
	GetAt(int) T
	SetAt(int, T)
}

type DynamicSequenceInterface[T any] interface {
	StaticSequenceInterface[T]
	InsertFirst(T)
	DeleteFirst() T
	InsertLast(T)
	DeleteLast() T
	InsertAt(int, T)
	DeleteAt(int) T
}
