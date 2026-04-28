package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const (
	RUTA_ARCHIVO1 = "archivo1.in"
	RUTA_ARCHIVO2 = "archivo2.in"
)

func leer_archivo(ruta string) []int {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	var numeros []int
	for scanner.Scan() {
		linea := scanner.Text()
		num, _ := strconv.Atoi(linea)
		numeros = append(numeros, num)
	}
	return numeros
}

func imprimir_arreglo(arreglo []int) {
	for _, num := range arreglo {
		fmt.Printf("%d\n", num)
	}
}

func main() {
	numeros1 := leer_archivo(RUTA_ARCHIVO1)
	numeros2 := leer_archivo(RUTA_ARCHIVO2)

	var arr_a_ordenar []int
	if ejercicios.Comparar(numeros1, numeros2) == -1 {
		arr_a_ordenar = numeros2
	} else {
		arr_a_ordenar = numeros1
	}

	ejercicios.Seleccion(arr_a_ordenar)
	imprimir_arreglo(arr_a_ordenar)

}
