package models

import "fmt"

// Estructura que representa un producto
type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

// Metodo para mostrar la informacion del producto
func (p Producto) MostrarInfo() {
	fmt.Println("------- PRODUCTO -------")
	fmt.Println("ID:", p.ID)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)
	fmt.Println("Stock:", p.Stock)
	fmt.Println("Categoria:", p.Categoria)
}
