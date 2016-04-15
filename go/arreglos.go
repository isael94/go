//Arreglos 

// Declarar un array de 5 strings cada uno
// lo iniciamos en cero, se declara un segundo arreglo
//de 5 strings

package main

import "fmt"

func main() {
	var nombres [5]string 

	//Declarar un arreglo pre-inicializado
	amigos:= [5]string{"Carmen", "Caro","Arturo gay","Carlos","Gerardo"}
	//Asignar el arreglo de los amigos
	nombres = amigos

	
		for i, nombre := range nombres{
			fmt.Println(nombre, &nombres[i])
		}
	
}