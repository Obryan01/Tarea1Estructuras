package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

type registro struct {
	contraseña [7]byte //Se declara la variable
	Usuario    int
	Contraseña string
}

func main() {

	fmt.Println("Desea")
	// La longitud será el doble de la capacidad del arreglo (64 en este caso)
	var contraseña [7]byte //Se declara la variable
	var nuevo string

	_, err := io.ReadFull(rand.Reader, contraseña[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Su contraseña aleatoria es = ", hex.EncodeToString(contraseña[:]))

	fmt.Println("Es obligatorio cambiar de contraseña: ")
	fmt.Scanln(&nuevo)

	fmt.Println("Su nueva contraseña: ", nuevo)
	fmt.Println("Su Anterior contraseña: ", contraseña)

}

func Menu() int {
	var operacion int
	fmt.Println(">>>>>>>>>>>>MENU<<<<<<<<<<<<")
	fmt.Println(">>>>>>>>>>>>Bienvenidos<<<<<<<<<<<<")
	fmt.Println("1) Inicio sesion")
	fmt.Println("2) Registrarme")
	fmt.Println("Seleccione el tipo de operacion:")
	fmt.Scanf("%d\n", &operacion)
	return operacion
}

func Registrarme(operacion int) {

}
