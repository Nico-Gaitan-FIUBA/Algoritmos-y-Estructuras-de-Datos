package lista

type nodo[T any] struct {
	dato T
	sig  *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}
