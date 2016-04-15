//Mapas

//Declarar y hacer un mapa de valores enteros con un string como llave
//Llenar el mapa con 5 valores e iterar sobre el mapa 

package main 

import "fmt"

func main() {
	//Declarar y hacer el mapa

	departamentos := make(map[string]int)

	//Llenar el mapa con 5 valores
	departamentos["Devs"] = 25
	departamentos["Marketing"] = 50
	departamentos["Ejecutivos"] = 4
	departamentos["Ventas"] = 60
	departamentos["Mantenimiento"] = 8

	//Desplegar el valor de cada par llave/valor

	for key, value := range departamentos {
		fmt.Printf("Departamentos: %s Personas: %d\n",key, value)
	}
} 
