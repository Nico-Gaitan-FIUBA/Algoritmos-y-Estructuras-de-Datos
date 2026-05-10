package diccionario

import (
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
