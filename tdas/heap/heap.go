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

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arrAdaptado := adaptoArregloAHeap(arreglo)
	return &heap[T]{arr: arrAdaptado, funcion_cmp: funcion_cmp}
}

func adaptoArregloAHeap[T any](arr []T) []T {
	mitad := (0 + len(arr)) / 2
	return _adaptoArregloAHeap(arr, 0, mitad)
}

func _adaptoArregloAHeap[T any](arr []T, inicio int, mitad int) []T {
	if inicio != mitad {
		downHeap(arr[mitad], mitad)
	}
	return _adaptoArregloAHeap(arr, inicio, mitad-1)
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
	if h.funcion_cmp(elem, padre) <= 0 {
		return
	}
	h.swap(pos, posPadre)
	h.upHeap(elem, posPadre)

}

func (h *heap[T]) downHeap(elem T, pos int) {
	hijoIzquierdo, posHijoIzquierdo := h.buscoPosHijoIzquierdo(pos)
	hijoDerecho, posHijoDerecho := h.buscoPosHijoDerecho(pos)
	hijoMayor := hijoDerecho
	posHijoMayor := posHijoDerecho
	posDelMayor := pos

	switch {
	case posHijoIzquierdo >= len(h.arr):
		return
	case posHijoDerecho >= len(h.arr):
		hijoMayor = hijoIzquierdo
		posHijoMayor = posHijoIzquierdo
	default:
		if h.funcion_cmp(hijoIzquierdo, hijoDerecho) > 0 {
			hijoMayor = hijoIzquierdo
			posHijoMayor = posHijoIzquierdo
		}
	}
	if h.funcion_cmp(hijoMayor, elem) > 0 {
		posDelMayor = posHijoMayor
		h.swap(pos, posDelMayor)
		h.downHeap(elem, posDelMayor)
	}
}

func (h *heap[T]) EstaVacia() bool {
	return len(h.arr) == PRIMER_POS
}

func (h *heap[T]) Encolar(elem T) {
	h.arr = append(h.arr, elem)
	if len(h.arr) > UN_ELEMENTO {
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
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	posUltimoELem := len(h.arr) - 1
	elemDesencolado := h.arr[PRIMER_POS]

	h.swap(PRIMER_POS, posUltimoELem)
	h.arr = h.arr[:posUltimoELem]

	primerElem := h.arr[PRIMER_POS]
	h.downHeap(primerElem, PRIMER_POS)

	return elemDesencolado
}

func (h *heap[T]) Cantidad() int {
	return len(h.arr)
}
