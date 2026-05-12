package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	_CARGA_MAX          = 2
	_TAMANO_INICIAL     = 11
	_CANTIDAD_INICIAL   = 0
	_FACTOR_REDIMENSION = 2
	_POS_INICIAL        = 0
)

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	tabla       *hashAbierto[K, V]
	actualHash  int
	actualLista TDALista.IteradorLista[K]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		tabla:    make([]TDALista.Lista[parClaveValor[K, V]], _TAMANO_INICIAL),
		tam:      _TAMANO_INICIAL,
		cantidad: _CANTIDAD_INICIAL,
	}
}

func JenkinsHash(clave string, tam int) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(clave); i++ {
		hash += uint32(clave[i])
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash % uint32(tam)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashAbierto[K, V]) redimensionar(tam int) {
	nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	for _, lista := range h.tabla {
		if lista != nil {
			lista.Iterar(func(par parClaveValor[K, V]) bool {
				claveHash := JenkinsHash(string(convertirABytes(par.clave)), tam)
				if nuevaTabla[claveHash] == nil {
					nuevaTabla[claveHash] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
				}
				nuevaTabla[claveHash].InsertarUltimo(parClaveValor[K, V]{clave: par.clave, valor: par.valor})
				return true
			})
		}
	}
	h.tabla = nuevaTabla
	h.tam = tam
}

func (h *hashAbierto[K, V]) actualizarValor(lista TDALista.Lista[parClaveValor[K, V]], clave K, dato V) bool {
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			iter.Borrar()
			iter.Insertar(parClaveValor[K, V]{clave: clave, valor: dato})
			return true
		}
	}
	return false
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {

	claveHash := JenkinsHash(string(convertirABytes(clave)), h.tam)
	lista := h.tabla[claveHash]

	if lista == nil {
		lista = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
		h.tabla[claveHash] = lista
	}
	if h.actualizarValor(lista, clave, dato) {
		return
	}

	lista.InsertarUltimo(parClaveValor[K, V]{clave: clave, valor: dato})
	h.cantidad++

	if h.cantidad/h.tam == _CARGA_MAX {
		h.redimensionar(h.tam * _FACTOR_REDIMENSION)

	}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	claveHash := JenkinsHash(string(convertirABytes(clave)), h.tam)
	lista := h.tabla[claveHash]
	if lista == nil {
		return false
	}
	var pertenece bool
	lista.Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			pertenece = true
			return false
		}
		return true
	})
	return pertenece
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	claveHash := JenkinsHash(string(convertirABytes(clave)), h.tam)
	lista := h.tabla[claveHash]

	if lista == nil {
		panic("La clave no pertenece al diccionario")
	}

	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			return iter.VerActual().valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	claveHash := JenkinsHash(string(convertirABytes(clave)), h.tam)
	lista := h.tabla[claveHash]

	if lista == nil {
		panic("La clave no pertenece al diccionario")
	}

	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			elemento := iter.Borrar()
			valorBorrado := elemento.valor
			h.cantidad--
			if h.cantidad == h.tam/2 {
				h.redimensionar(h.tam / _FACTOR_REDIMENSION)
			}
			return valorBorrado
		}
	}

	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	debeContinuar := true
	for i := _POS_INICIAL; i < h.tam && debeContinuar; i++ {
		lista := h.tabla[i]
		if lista != nil {
			lista.Iterar(func(par parClaveValor[K, V]) bool {
				if !visitar(par.clave, par.valor) {
					debeContinuar = false
				}
				return debeContinuar
			})
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterDiccionario[K, V]{}
	iterador.actualHash = 0

	for _, lista := range h.tabla {
		if lista == nil {
			iterador.actualHash++
			if iterador.actualHash >= h.tam {
				return
			}
		}

	}
}
