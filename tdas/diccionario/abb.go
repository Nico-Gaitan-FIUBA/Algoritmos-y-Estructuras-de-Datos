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

func (a *abb[K, V]) buscarNodoYPadre(claveDelNodoABuscar K, subRaiz *nodoAbb[K, V], padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if subRaiz == nil {
		return nil, nil
	}
	cmp := a.funcion_cmp(claveDelNodoABuscar, subRaiz.clave)

	switch {
	case cmp < 0: // --> clave: 1  raiz.clave: 2
		return a.buscarNodoYPadre(claveDelNodoABuscar, subRaiz.izq, subRaiz)
	case cmp > 0: // --> clave: 3  raiz.clave: 2
		return a.buscarNodoYPadre(claveDelNodoABuscar, subRaiz.der, subRaiz)
	default:
		return subRaiz, padre
	}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	nodoAObtener, padre := a.buscarNodoYPadre(clave, a.raiz, nil)

	if nodoAObtener == nil {
		nuevoNodo := &nodoAbb[K, V]{clave: clave, dato: dato}
		if a.raiz == nil {
			a.raiz = nuevoNodo
		} else if a.funcion_cmp(clave, padre.clave) < 0 {
			padre.izq = nuevoNodo
		} else {
			padre.der = nuevoNodo
		}
		a.cant++
	} else {
		nodoAObtener.dato = dato
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	nodoAObtener, _ := a.buscarNodoYPadre(clave, a.raiz, nil)

	return nodoAObtener != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodoAObtener, _ := a.buscarNodoYPadre(clave, a.raiz, nil)

	if nodoAObtener == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodoAObtener.dato
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodoAObtener, padre := a.buscarNodoYPadre(clave, a.raiz, nil)

}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}
