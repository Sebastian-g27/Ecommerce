package services

import (
	"awesomeProject/models"
	"fmt"
)

// Servicio para administrar pedidos
type PedidoService struct {
	Pedidos []models.Pedido
}

// Registrar un pedido
func (s *PedidoService) RegistrarPedido(p models.Pedido) {
	p.CalcularTotal()
	s.Pedidos = append(s.Pedidos, p)

	fmt.Println("Pedido registrado correctamente")
}

// Mostrar todos los pedidos
func (s PedidoService) ListarPedidos() {
	fmt.Println("===== LISTA DE PEDIDOS =====")

	for _, pedido := range s.Pedidos {
		pedido.MostrarInfo()
	}
}

// Buscar pedido por ID
func (s PedidoService) BuscarPedido(id int) {
	for _, pedido := range s.Pedidos {
		if pedido.ID == id {
			pedido.MostrarInfo()
			return
		}
	}

	fmt.Println("Pedido no encontrado")
}

// Eliminar pedido
func (s *PedidoService) EliminarPedido(id int) {
	for i, pedido := range s.Pedidos {
		if pedido.ID == id {
			s.Pedidos = append(
				s.Pedidos[:i],
				s.Pedidos[i+1:]...,
			)

			fmt.Println("Pedido eliminado correctamente")
			return
		}
	}

	fmt.Println("Pedido no encontrado")
}
