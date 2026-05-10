package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	TAMANO_INICIAL   = 11
	CANTIDAD_INICIAL = 0
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

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		tabla:    make([]TDALista.Lista[parClaveValor[K, V]], TAMANO_INICIAL),
		tam:      TAMANO_INICIAL,
		cantidad: CANTIDAD_INICIAL,
	}
}

func JenkinsHash(key string) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(key); i++ {
		hash += uint32(key[i])
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash % TAMANO_INICIAL
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	claveHash := JenkinsHash(string(convertirABytes(clave)))
	lista := h.tabla[claveHash]
	if h.Pertenece(clave) {

	} else {

	}

}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	claveHash := JenkinsHash(string(convertirABytes(clave)))
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
