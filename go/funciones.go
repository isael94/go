//Funciones

//Declaramos un struct para guardar informacion de un usuario
//Declarar una funcion que crea un valor y regresa apuntadores de este
// tipo y un error como valor
//Llamar esta funcion desde main y regresar el valor
//Hacer una segunda llamada a la funcion pero ignrorando el valor
// y solo probando el error como valor

package main 

import "fmt"

//el usuario representa un usuario en el sistema

type usuario struct{
	nombre string
	email string

}

//crear un nuevo usuario y regresa apuntadores de valores de tipo usuario

func nuevoUsuario() (*usuario, error) {
	return &usuario{"Javier", "javier@isael.com"}, nil
}
// main es la entrada de todos los programas
func main() {
	//crear el valor del tipo usuario
	u, err := nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Imprimimos el valor

	fmt.Println(*u)

	//Segundo llamado a la funcion y que solo cheque el error en el regreso

	_, err = nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}
}