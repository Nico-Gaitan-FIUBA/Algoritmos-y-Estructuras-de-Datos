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

type iteradorLista[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
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
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevoNodo := &nodo[T]{dato: valor}
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.sig = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
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
	l.largo--
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

//iteradores

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := &iteradorLista[T]{}
	iterador.lista = l
	iterador.actual = l.primero
	iterador.anterior = nil
	return iterador
}

func (i *iteradorLista[T]) VerActual() T {
	if i.actual == nil {
		panic("El iterador termino de iterar")
	}
	return i.actual.dato
}

func (i *iteradorLista[T]) HayAlgoMas() bool {
	return i.actual != nil
}

func (i *iteradorLista[T]) Avanzar() {
	if i.actual.sig == nil {
		panic("El iterador termino de iterar")
	}
	actualAux := i.actual
	i.anterior = actualAux
	i.actual = i.actual.sig
}

func (i *iteradorLista[T]) Insertar(valor T) {
	if i.anterior == nil {
		i.lista.InsertarPrimero(valor)
		i.actual = i.anterior
	} else if !i.HayAlgoMas() {
		i.lista.InsertarUltimo(valor)
		i.actual = i.actual.sig
	} else {
		nuevoNodo := &nodo[T]{dato: valor}
		aux := i.actual
		i.actual = nuevoNodo
		i.anterior.sig = nuevoNodo
		i.actual.sig = aux
		i.lista.largo++
	}
}

func (i *iteradorLista[T]) Borrar() T {
	if i.actual == nil {
		panic("El iterador termino de iterar")
	}
	if i.anterior == nil {
		i.lista.BorrarPrimero()
	} else if !i.HayAlgoMas() {
		i.lista.ultimo = i.anterior
	}
	return i.actual.dato
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	return true
}
