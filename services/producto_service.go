package services

import (
	"awesomeProject/models"
	"fmt"
)

// Servicio para administrar productos
type ProductoService struct {
	Productos []models.Producto
}

// Registrar un producto
func (s *ProductoService) RegistrarProducto(p models.Producto) {
	s.Productos = append(s.Productos, p)
	fmt.Println("Producto registrado correctamente")
}

// Mostrar todos los productos
func (s ProductoService) ListarProductos() {
	fmt.Println("===== LISTA DE PRODUCTOS =====")

	for _, producto := range s.Productos {
		producto.MostrarInfo()
	}
}

// Buscar un producto por su ID
func (s ProductoService) BuscarProducto(id int) {
	for _, producto := range s.Productos {
		if producto.ID == id {
			producto.MostrarInfo()
			return
		}
	}

	fmt.Println("Producto no encontrado")
}

// Eliminar un producto
func (s *ProductoService) EliminarProducto(id int) {
	for i, producto := range s.Productos {
		if producto.ID == id {
			s.Productos = append(
				s.Productos[:i],
				s.Productos[i+1:]...,
			)
			fmt.Println("Producto eliminado correctamente")
			return
		}
	}

	fmt.Println("Producto no encontrado")
}
