package ds_interfaces

type StaticSequenceInterface[T any] interface {
	ContainerInterface[T]
	GetAt(int) (T, error)
	SetAt(int, T) bool
}

type DynamicSequenceInterface[T any] interface {
	StaticSequenceInterface[T]
	InsertFirst(T)
	DeleteFirst() (T, error)
	InsertLast(T)
	DeleteLast() (T, error)
	InsertAt(int, T) bool
	DeleteAt(int) (T, error)
}
