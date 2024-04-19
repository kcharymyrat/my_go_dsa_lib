package ds_interfaces

type OrderSetInterface[T any] interface {
	FindMin() T
	FindMax() T
	FindPrev() T
	FindNext() T
	// IterOrd() T
}

type StaticSetInterface[T any] interface {
	ContainerInterface[T]
	Find(T) T
	OrderSetInterface[T]
}

type DynamicSetInterface[T any] interface {
	StaticSequenceInterface[T]
	Insert(T)
	Delete(T) T
}
