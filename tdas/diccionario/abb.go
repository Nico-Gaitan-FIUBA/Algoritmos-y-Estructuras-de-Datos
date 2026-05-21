package diccionario

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz        *nodoAbb[K, V]
	funcion_cmp func(K, K) int
	cant        int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz: nil,
		funcion_cmp: func(clave1, clave2 K) int {
			return funcion_cmp(clave1, clave2)
		},
		cant: 0,
	}
}

// La función de comparación, recibe dos claves y devuelve:

//     Un entero menor que 0 si la primera clave es menor que la segunda. --> negativo si 1era < 2da
//     Un entero mayor que 0 si la primera clave es mayor que la segunda. --> positivo si 1era > 2da
//     0 si ambas claves son iguales.

func (a *abb[K, V]) buscarNodo(claveDelNodoABuscar K, subRaiz *nodoAbb[K, V]) *nodoAbb[K, V] {
	if subRaiz == nil {
		return nil
	}
	cmp := a.funcion_cmp(claveDelNodoABuscar, subRaiz.clave)

	switch {
	case cmp < 0: // --> clave: 1  raiz.clave: 2
		return a.buscarNodo(claveDelNodoABuscar, subRaiz.izq)
	case cmp > 0: // --> clave: 3  raiz.clave: 2
		return a.buscarNodo(claveDelNodoABuscar, subRaiz.der)
	default:
		return subRaiz
	}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {

}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.buscarNodo(clave, a.raiz) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodoAObtener := a.buscarNodo(clave, a.raiz)

	if nodoAObtener == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodoAObtener.dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}
