package cola_prioridad

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaPrioridadVacia(t *testing.T) {
	cola := CrearColaPrioridad(func(a, b int) int {
		return a - b
	})
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestColaPrioridadEncolarUnElemento(t *testing.T) {
	cola := CrearColaPrioridad(func(a, b int) int {
		return a - b
	})
	cola.Encolar(2026)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Cantidad())
	require.Equal(t, 2026, cola.VerMax())
}

func TestColaPrioridadEncolarDosElementos(t *testing.T) {
	cola := CrearColaPrioridad(func(a, b int) int {
		return a - b
	})
	cola.Encolar(10)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Cantidad())
	require.Equal(t, 10, cola.VerMax())

	cola.Encolar(20)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.Cantidad())
	require.Equal(t, 20, cola.VerMax())

}

func TestColaPrioridadEncolarVariosElementosOrdenados(t *testing.T) {
	cola := CrearColaPrioridad(func(a, b int) int {
		return a - b
	})
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 5, cola.Cantidad())
	require.Equal(t, 5, cola.VerMax())

}

func TestColaPrioridadEncolarVariosElementosDesordenados(t *testing.T) {
	cola := CrearColaPrioridad(func(a, b int) int {
		return a - b
	})
	cola.Encolar(66)
	cola.Encolar(10)
	cola.Encolar(90)
	cola.Encolar(14)
	cola.Encolar(35)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 5, cola.Cantidad())
	require.Equal(t, 90, cola.VerMax())

}
