package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"fmt"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaEncolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.False(t, cola.EstaVacia())
}

func TestColaDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.Equal(t, 10, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.Equal(t, 10, cola.VerPrimero())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestColaUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 10, cola.VerPrimero())
	require.Equal(t, 10, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaPocosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())

	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())

	require.Equal(t, 1, cola.Desencolar())

	require.Equal(t, 2, cola.VerPrimero())

	require.Equal(t, 2, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaEncolarDespuesDeVaciar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	cola.Desencolar()
	cola.Encolar(20)
	require.False(t, cola.EstaVacia())
	cola.Encolar(25)
	cola.Encolar(30)
	require.Equal(t, 20, cola.VerPrimero())
	require.Equal(t, 20, cola.Desencolar())
	require.Equal(t, 25, cola.VerPrimero())
	require.Equal(t, 25, cola.Desencolar())
	require.Equal(t, 30, cola.VerPrimero())
	require.Equal(t, 30, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaMuchosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for i := 0; i < 10000; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("primero")
	cola.Encolar("segundo")
	require.Equal(t, "primero", cola.VerPrimero())
	require.Equal(t, "primero", cola.Desencolar())
	require.Equal(t, "segundo", cola.VerPrimero())
	require.Equal(t, "segundo", cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaMuchasStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	for i := 0; i < 10000; i++ {
		cola.Encolar(fmt.Sprintf("string%d", i))
		require.Equal(t, "string0", cola.VerPrimero())
	}
	for i := 0; i < 10000; i++ {
		require.Equal(t, fmt.Sprintf("string%d", i), cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}
