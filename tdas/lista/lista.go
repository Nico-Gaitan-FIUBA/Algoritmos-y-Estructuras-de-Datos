package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar recorre la lista desde el primero al último elemento, aplicando la función dada a cada uno de ellos.
	// Si la función devuelve false, se detiene la iteración. Si devuelve true, se continúa con el siguiente elemento.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador que recorre la lista desde el primero al último elemento.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	VerActual() T

	HayAlgoMas() bool

	Avanzar()

	Insertar(T)

	Borrar() T
}
