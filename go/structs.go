//Structs


package main 

import "fmt"

type usuario struct{
	nombre string
	direccion string
	edad int
	puesto string
}

 func main() {
		//Declarar la variable ususario y la iniciamos usando struct
 	javi := usuario{
 		nombre: "Javier",
 		direccion: "Bosques de Colon",
 		edad: 21,
 		puesto: "Programmer",
 	}
 	//Mostrar los valores de cada campo

 	fmt.Println("Nombre: ", javi.nombre)
 	fmt.Println("Direccion: ",javi.direccion)
 	fmt.Println("Edad: ",javi.edad)
 	fmt.Println("Puesto: ",javi.puesto)
 	//Declarar otro struct anonimo

 	nicole := struct{
 		nombre string
 		direccion string
 		edad int
 	} {
 		nombre: "Nicole",
 		direccion: "Calle 143",
 		edad: 22,
 	}
 	fmt.Println(nicole.nombre)
 	fmt.Println(nicole.direccion)
	}