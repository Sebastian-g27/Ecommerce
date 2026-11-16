package models

import "fmt"

// Representa un producto dentro de un pedido
type DetallePedido struct {
	Producto Producto
	Cantidad int
}

// Estructura que representa un pedido
type Pedido struct {
	ID       int
	Cliente  Cliente
	Detalles []DetallePedido
	Total    float64
	Estado   string
}

// Metodo para calcular el total del pedido
func (p *Pedido) CalcularTotal() {
	p.Total = 0

	for _, detalle := range p.Detalles {
		p.Total += detalle.Producto.Precio *
			float64(detalle.Cantidad)
	}
}

// Metodo para mostrar la informacion del pedido
func (p Pedido) MostrarInfo() {
	fmt.Println("------- PEDIDO -------")
	fmt.Println("ID:", p.ID)
	fmt.Println("Cliente:", p.Cliente.Nombre)

	fmt.Println("Productos:")

	for _, detalle := range p.Detalles {
		fmt.Println(
			detalle.Producto.Nombre,
			"Cantidad:", detalle.Cantidad,
		)
	}

	fmt.Println("Total:", p.Total)
	fmt.Println("Estado:", p.Estado)
}
