package cola

/* Definición del struct cola proporcionado por la cátedra. */

type nodo[T any] struct {
	dato T
	sig  *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := &nodo[T]{dato: valor}
	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.sig = nuevoNodo
		c.ultimo = nuevoNodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	valor := c.primero.dato
	c.primero = c.primero.sig
	if c.EstaVacia() {
		c.ultimo = c.primero
	}
	return valor
}
