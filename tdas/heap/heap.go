package cola_prioridad

type heap[T any] struct {
	arr         []T
	funcion_cmp func(T, T) int
}

func CrearColaPrioridad[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{arr: []T{}, funcion_cmp: funcion_cmp}
}

func (h *heap[T]) EstaVacia() bool {
	return len(h.arr) == 0
}
