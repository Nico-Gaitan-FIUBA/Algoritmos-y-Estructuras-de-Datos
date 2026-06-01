package cola_prioridad

const (
	PRIMER_POS  = 0
	UN_ELEMENTO = 1
)

type heap[T any] struct {
	arr         []T
	funcion_cmp func(T, T) int
}

// La función de comparación, recibe dos claves y devuelve:

//     Un entero menor que 0 si la primera clave es menor que la segunda. --> negativo si 1era < 2da
//     Un entero mayor que 0 si la primera clave es mayor que la segunda. --> positivo si 1era > 2da
//     0 si ambas claves son iguales.

func CrearColaPrioridad[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{arr: []T{}, funcion_cmp: funcion_cmp}
}

func (h *heap[T]) buscoPosHijoIzquierdo(indice int) (T, int) {
	posHijoIzq := 2*indice + 1
	hijoIzq := h.arr[posHijoIzq]
	return hijoIzq, posHijoIzq
}

func (h *heap[T]) buscoPosHijoDerecho(indice int) (T, int) {
	posHijoDer := 2*indice + 2
	hijoDer := h.arr[posHijoDer]
	return hijoDer, posHijoDer
}

func (h *heap[T]) buscoPosPadre(indice int) (T, int) {
	posPadre := (indice - 1) / 2
	padre := h.arr[posPadre]
	return padre, posPadre
}

func (h *heap[T]) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *heap[T]) upHeap(elem T, pos int) {
	padre, posPadre := h.buscoPosPadre(pos)
	if posPadre == PRIMER_POS {
		return
	}
	if h.funcion_cmp(elem, padre) > 0 {
		h.swap(pos, posPadre)
		h.upHeap(elem, posPadre)
	}
}

func (h *heap[T]) downHeap(elem T, pos int) {
	hijoIzquierdo, posHijoIzquierdo := h.buscoPosHijoIzquierdo(pos)
	hijoDerecho, posHijoDerecho := h.buscoPosHijoDerecho(pos)
	hijoMayor := hijoDerecho
	posHijoMayor := posHijoDerecho
	mayor := elem

	if h.funcion_cmp(hijoIzquierdo, hijoDerecho) > 0 {
		hijoMayor = hijoIzquierdo
		posHijoMayor = posHijoIzquierdo
	}
	if h.funcion_cmp(hijoMayor, elem) > 0 {
		mayor = hijoMayor
	}
	h.swap(pos, mayor)

}

func (h *heap[T]) EstaVacia() bool {
	return len(h.arr) == PRIMER_POS
}

func (h *heap[T]) Encolar(elem T) {
	h.arr = append(h.arr, elem)
	if len(h.arr)-1 > UN_ELEMENTO {
		h.upHeap(elem, len(h.arr)-1)
	}
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.arr[PRIMER_POS]
}

func (h *heap[T]) Desencolar() T {
	posUltimoELem := len(h.arr) - 1    // posicion = 6
	ultimoElem := h.arr[posUltimoELem] //ultimo elemento = 20

	h.swap(PRIMER_POS, posUltimoELem) //ultimo elemento = 12 y primer elemento = 20
	h.arr = h.arr[:posUltimoELem-1]

	primerElem := h.arr[PRIMER_POS] // primer elemento = 20
	h.downHeap(primerElem, PRIMER_POS)

	return ultimoElem
}

func (h *heap[T]) Cantidad() int {
	return len(h.arr)
}
