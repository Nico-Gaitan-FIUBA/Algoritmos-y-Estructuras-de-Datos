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
	h := &heap[T]{arr: arreglo, funcion_cmp: funcion_cmp}

	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		downHeap(h.arr, h.funcion_cmp, h.arr[i], i)
	}
	return h
}

func buscoPosHijoIzquierdo(indice int) int {
	posHijoIzq := 2*indice + 1
	return posHijoIzq
}

func buscoPosHijoDerecho(indice int) int {
	posHijoDer := 2*indice + 2
	return posHijoDer
}

func buscoPosPadre[T any](arr []T, indice int) (T, int) {
	posPadre := (indice - 1) / 2
	padre := arr[posPadre]
	return padre, posPadre
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func upHeap[T any](arr []T, cmp func(T, T) int, elem T, pos int) {
	padre, posPadre := buscoPosPadre(arr, pos)
	if cmp(elem, padre) <= 0 {
		return
	}
	swap(arr, pos, posPadre)
	upHeap(arr, cmp, elem, posPadre)
}

func downHeap[T any](arr []T, cmp func(T, T) int, elem T, pos int) { // [20, 15, 10, 13, ""]
	posHijoIzquierdo := buscoPosHijoIzquierdo(pos)
	posHijoDerecho := buscoPosHijoDerecho(pos)
	posHijoMayor := posHijoDerecho
	posDelMayor := pos
	hijoMayor := elem

	switch {
	case posHijoIzquierdo >= len(arr):
		return
	case posHijoDerecho >= len(arr):
		hijoIzquierdo := arr[posHijoIzquierdo]
		hijoMayor = hijoIzquierdo
		posHijoMayor = posHijoIzquierdo
	default:
		hijoIzquierdo := arr[posHijoIzquierdo]
		hijoDerecho := arr[posHijoDerecho]
		if cmp(hijoIzquierdo, hijoDerecho) > 0 {
			hijoMayor = hijoIzquierdo
			posHijoMayor = posHijoIzquierdo
		} else {
			hijoMayor = hijoDerecho
		}
	}
	if cmp(hijoMayor, elem) > 0 {
		posDelMayor = posHijoMayor
		swap(arr, pos, posDelMayor)
		downHeap(arr, cmp, elem, posDelMayor)
	}
}

func (h *heap[T]) EstaVacia() bool {
	return len(h.arr) == PRIMER_POS
}

func (h *heap[T]) Encolar(elem T) {
	h.arr = append(h.arr, elem)
	if len(h.arr) > UN_ELEMENTO {
		upHeap(h.arr, h.funcion_cmp, elem, len(h.arr)-1)
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

	swap(h.arr, PRIMER_POS, posUltimoELem)
	primerElem := h.arr[PRIMER_POS]

	h.arr = h.arr[:posUltimoELem]
	if !h.EstaVacia() {
		downHeap(h.arr, h.funcion_cmp, primerElem, PRIMER_POS)
	}

	return elemDesencolado
}

func (h *heap[T]) Cantidad() int {
	return len(h.arr)
}
