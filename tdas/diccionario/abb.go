package diccionario

type nodoABB[T any] struct {
	izq  *nodoABB
	der  *nodoABB
	dato T
}

type ABB[T any] struct {
	raiz *nodoABB[T]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &ABB[K, V]{raiz: nil}
}
