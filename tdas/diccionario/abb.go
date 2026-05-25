package diccionario

import (
	TDAPila "tdas/pila"
)

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

type iterDiccionarioAbb[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	arbol *abb[K, V]
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

func (a *abb[K, V]) buscarNodoYPadre(claveDelNodoABuscar K, raiz *nodoAbb[K, V], padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if raiz == nil {
		return nil, padre
	}
	cmp := a.funcion_cmp(claveDelNodoABuscar, raiz.clave)

	switch {
	case cmp < 0: // --> clave: 1  raiz.clave: 2
		return a.buscarNodoYPadre(claveDelNodoABuscar, raiz.izq, raiz)
	case cmp > 0: // --> clave: 3  raiz.clave: 2
		return a.buscarNodoYPadre(claveDelNodoABuscar, raiz.der, raiz)
	default:
		return raiz, padre
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

func (a *abb[K, V]) buscarReemplazo(inicio *nodoAbb[K, V]) *nodoAbb[K, V] {
	if inicio.der == nil {
		return inicio
	}
	return a.buscarReemplazo(inicio.der)
}

func (a *abb[K, V]) procesarCasoDosHijos(nodoABorrar *nodoAbb[K, V]) {
	reemplazo := a.buscarReemplazo(nodoABorrar.izq)
	claveReemplazo := reemplazo.clave
	datoReemplazo := reemplazo.dato

	a.Borrar(claveReemplazo)

	nodoABorrar.clave = claveReemplazo
	nodoABorrar.dato = datoReemplazo
}

func (a *abb[K, V]) procesarCasoCeroOUnHijo(nodoABorrar *nodoAbb[K, V], padre *nodoAbb[K, V]) {
	var huerfano *nodoAbb[K, V]
	if nodoABorrar.izq == nil {
		huerfano = nodoABorrar.der
	} else {
		huerfano = nodoABorrar.izq
	}

	if padre == nil {
		a.raiz = huerfano
	} else if a.funcion_cmp(nodoABorrar.clave, padre.clave) > 0 {
		padre.der = huerfano
	} else {
		padre.izq = huerfano
	}
	a.cant--
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodoABorrar, padre := a.buscarNodoYPadre(clave, a.raiz, nil)
	if nodoABorrar == nil {
		panic("La clave no pertenece al diccionario")
	}

	datoBorrado := nodoABorrar.dato

	if nodoABorrar.izq != nil && nodoABorrar.der != nil {
		a.procesarCasoDosHijos(nodoABorrar)
	} else {
		a.procesarCasoCeroOUnHijo(nodoABorrar, padre)
	}
	return datoBorrado
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if a.raiz != nil {
		a.raiz.iterar(visitar)
	}
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}

	nodo.izq.iterar(visitar)
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	nodo.der.iterar(visitar)
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if a.raiz != nil {
		a.raiz.iterarRango(a, desde, hasta, visitar)
	}
}

func (nodo *nodoAbb[K, V]) iterarRango(arbol *abb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}

	if arbol.funcion_cmp(nodo.clave, *hasta) > 0 {
		nodo.izq.iterarRango(arbol, desde, hasta, visitar)
	}

	if arbol.funcion_cmp(nodo.clave, *desde) > 0 && arbol.funcion_cmp(nodo.clave, *hasta) < 0 {
		nodo.izq.iterarRango(arbol, desde, hasta, visitar)

		if !visitar(nodo.clave, nodo.dato) {
			return
		}

		nodo.der.iterarRango(arbol, desde, hasta, visitar)
	}

	if arbol.funcion_cmp(nodo.clave, *desde) < 0 {
		nodo.der.iterarRango(arbol, desde, hasta, visitar)
	}
}

func (a *abb[K, V]) buscarElMasChico(iterador *iterDiccionarioAbb[K, V], inicio *nodoAbb[K, V], desde *K, hasta *K) *nodoAbb[K, V] {
	iterador.pila.Apilar(inicio) //inicio = 5

	// pila.vertope() esta en el rango?
	// si: veo inicio.izq
	// no: pila.vertope.desapilar()

	if inicio.izq == nil {
		return inicio
	}
	return a.buscarElMasChico(iterador, inicio.izq, desde, hasta)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := &iterDiccionarioAbb[K, V]{
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
		arbol: a,
	}

	primerNodo := a.buscarElMasChico(iterador, a.raiz, desde, hasta)

}

func (i *iterDiccionarioAbb[K, V]) HayAlgoMas() bool {
	return !i.pila.EstaVacia()
}

func (i *iterDiccionarioAbb[K, V]) VerActual() (K, V) {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return i.pila.VerTope().clave, i.pila.VerTope().dato
}

func (i *iterDiccionarioAbb[K, V]) Avanzar() {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

}
