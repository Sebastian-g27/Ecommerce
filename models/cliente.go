package models

import "fmt"

// Estructura que representa un cliente
type Cliente struct {
	ID        int
	Nombre    string
	Correo    string
	Telefono  string
	Direccion string
}

// Metodo para mostrar informacion del cliente
func (c Cliente) MostrarInfo() {
	fmt.Println("------- CLIENTE -------")
	fmt.Println("ID:", c.ID)
	fmt.Println("Nombre:", c.Nombre)
	fmt.Println("Correo:", c.Correo)
	fmt.Println("Telefono:", c.Telefono)
	fmt.Println("Direccion:", c.Direccion)
}
