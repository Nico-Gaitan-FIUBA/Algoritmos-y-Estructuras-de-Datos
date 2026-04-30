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

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevoNodo := &nodo[T]{dato: valor}
	if l.EstaVacia() {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		aux := l.primero
		l.primero = nuevoNodo
		l.primero.sig = aux
	}
}

func (l *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevoNodo := &nodo[T]{dato: valor}
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.sig = nuevoNodo
	}
	l.ultimo = nuevoNodo
}
