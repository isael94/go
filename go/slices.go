//Slices
//Declarar un slice de enteros vacio
//Crear un loop que meta 10 valores

//Iterar sobre el slice y mostrar cada valor
//Declarar un slice de 5 string e inicializar dicho slice con valores string
//Impriir todos los elementos
//Tomar un slice del primer y segundo indice
//Desplegar la posicion y valor de cada uno de los elementos

package main 

import "fmt"

func main() {
	//Declarar un slice de enteros vacios
	var numeros []int

	//Meter 10 numeros al slice 
	for i:= 0; i < 10; i++{
		numeros = append(numeros, i*10)
	}

	//Mostrar cada valor
	for _, numero := range numeros{
		fmt.Println(numero)
	}

	//Declarar un slice de strings"
	frutas := []string{"manzana","naranja","pera","piña","mango"}

	//Mostrar cada indice(posicion) y nombre

	for i, fruta := range frutas{
		fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
	}

	//Tomar un slice de indice 1 y 2

	slice := frutas[0:3]

	//Mostrar el valor del nuevo slice

	for i, fruta := range slice{
		fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
	}
}