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

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	valor := l.primero.dato
	l.primero = l.primero.sig
	if l.primero == nil {
		l.ultimo = l.primero
	}
	return valor
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	valor := l.primero.dato
	return valor
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	valor := l.ultimo.dato
	return valor
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}
