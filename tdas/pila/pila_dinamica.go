package pila

/* Definición del struct pila proporcionado por la cátedra. */

const TAMANO_INICIAL = 10
const FACTOR_REDIMENSION = 2
const FACTOR_REDUCCION = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, TAMANO_INICIAL)}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * FACTOR_REDIMENSION)
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	valorDesapilado := p.datos[p.cantidad-1]
	p.cantidad--
	if p.cantidad <= len(p.datos)/FACTOR_REDUCCION && len(p.datos) > TAMANO_INICIAL {
		p.redimensionar(len(p.datos) / FACTOR_REDIMENSION)
	}
	return valorDesapilado
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	nuevoSlice := make([]T, tam)
	copy(nuevoSlice, p.datos[:p.cantidad])
	p.datos = nuevoSlice
}
