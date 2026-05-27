package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"fmt"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.False(t, pila.EstaVacia())
}

func TestPilaDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.Equal(t, 10, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.Equal(t, 10, pila.VerTope())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestPilaUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 10, pila.VerTope())
	require.Equal(t, 10, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaPocosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 1, pila.VerTope())

	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())

	require.Equal(t, 2, pila.Desapilar())

	require.Equal(t, 1, pila.VerTope())

	require.Equal(t, 1, pila.Desapilar())

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaApilarDespuesDeVaciar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	pila.Desapilar()
	pila.Apilar(20)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 20, pila.VerTope())
	pila.Apilar(25)
	require.Equal(t, 25, pila.VerTope())
	pila.Apilar(30)
	require.Equal(t, 30, pila.VerTope())
	require.Equal(t, 30, pila.Desapilar())
	require.Equal(t, 25, pila.VerTope())
	require.Equal(t, 25, pila.Desapilar())
	require.Equal(t, 20, pila.VerTope())
	require.Equal(t, 20, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaMuchosElementosInts(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := 9999; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
		if i > 0 {
			require.Equal(t, i-1, pila.VerTope())
		}
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	require.Equal(t, "hola", pila.VerTope())
	pila.Apilar("mundo")
	require.Equal(t, "mundo", pila.VerTope())
	require.Equal(t, "mundo", pila.Desapilar())
	require.Equal(t, "hola", pila.VerTope())
	require.Equal(t, "hola", pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaMuchosElementosStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	for i := 0; i < 10000; i++ {
		pila.Apilar(fmt.Sprintf("string%d", i))
		require.Equal(t, fmt.Sprintf("string%d", i), pila.VerTope())
	}
	for i := 9999; i >= 0; i-- {
		require.Equal(t, fmt.Sprintf("string%d", i), pila.Desapilar())
		if i > 0 {
			require.Equal(t, fmt.Sprintf("string%d", i-1), pila.VerTope())
		}
	}
	require.True(t, pila.EstaVacia())
}
