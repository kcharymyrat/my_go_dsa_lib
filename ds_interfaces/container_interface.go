package ds_interfaces

type ContainerInterface[T any] interface {
	Build(...T)
	Len() int
}
