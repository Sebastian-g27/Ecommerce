package services

import (
	"awesomeProject/models"
	"fmt"
)

// Servicio para administrar clientes
type ClienteService struct {
	Clientes []models.Cliente
}

// Registrar un cliente
func (s *ClienteService) RegistrarCliente(c models.Cliente) {
	s.Clientes = append(s.Clientes, c)
	fmt.Println("Cliente registrado correctamente")
}

// Mostrar todos los clientes
func (s ClienteService) ListarClientes() {
	fmt.Println("===== LISTA DE CLIENTES =====")

	for _, cliente := range s.Clientes {
		cliente.MostrarInfo()
	}
}

// Buscar cliente por ID
func (s ClienteService) BuscarCliente(id int) {
	for _, cliente := range s.Clientes {
		if cliente.ID == id {
			cliente.MostrarInfo()
			return
		}
	}

	fmt.Println("Cliente no encontrado")
}

// Eliminar cliente
func (s *ClienteService) EliminarCliente(id int) {
	for i, cliente := range s.Clientes {
		if cliente.ID == id {
			s.Clientes = append(
				s.Clientes[:i],
				s.Clientes[i+1:]...,
			)

			fmt.Println("Cliente eliminado correctamente")
			return
		}
	}

	fmt.Println("Cliente no encontrado")
}
