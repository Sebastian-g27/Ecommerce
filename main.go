package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"awesomeProject/models"
	"awesomeProject/services"
)

var lector = bufio.NewReader(os.Stdin)

func leerTexto(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerEntero(mensaje string) int {
	valor, err := strconv.Atoi(leerTexto(mensaje))
	if err != nil {
		fmt.Println("Ingrese un numero valido")
		return 0
	}
	return valor
}

func leerDecimal(mensaje string) float64 {
	texto := strings.ReplaceAll(leerTexto(mensaje), ",", ".")
	valor, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		fmt.Println("Ingrese un precio valido")
		return 0
	}
	return valor
}

func gestionarProductos(s *services.ProductoService) {
	for {
		fmt.Println("\n===== GESTION DE PRODUCTOS =====")
		fmt.Println("1. Registrar producto")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Buscar producto")
		fmt.Println("4. Eliminar producto")
		fmt.Println("5. Volver al menu")

		opcion := leerTexto("Seleccione: ")

		switch opcion {
		case "1":
			producto := models.Producto{
				ID:        leerEntero("ID: "),
				Nombre:    leerTexto("Nombre: "),
				Precio:    leerDecimal("Precio: "),
				Stock:     leerEntero("Stock: "),
				Categoria: leerTexto("Categoria: "),
			}

			if producto.ID <= 0 || producto.Nombre == "" ||
				producto.Precio <= 0 || producto.Stock < 0 ||
				producto.Categoria == "" {
				fmt.Println("Datos invalidos. Intente nuevamente.")
				continue
			}

			s.RegistrarProducto(producto)

		case "2":
			mostrarTablaProductos(s.Productos)

		case "3":
			id := leerEntero("Ingrese el ID del producto: ")
			s.BuscarProducto(id)

		case "4":
			id := leerEntero("Ingrese el ID a eliminar: ")
			s.EliminarProducto(id)

		case "5":
			return

		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func main() {
	productoService := services.ProductoService{}
	clienteService := services.ClienteService{}
	pedidoService := services.PedidoService{}

	for {
		fmt.Println("\n================================")
		fmt.Println("    SISTEMA DE GESTION E-COMMERCE")
		fmt.Println("================================")
		fmt.Println("1. Gestion de productos")
		fmt.Println("2. Gestion de clientes")
		fmt.Println("3. Gestion de pedidos")
		fmt.Println("4. Mostrar reportes")
		fmt.Println("5. Salir")
		fmt.Println("================================")

		opcion := leerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			gestionarProductos(&productoService)

		case "2":
			fmt.Println("\n===== LISTA DE CLIENTES =====")
			clienteService.ListarClientes()

		case "3":
			fmt.Println("\n===== LISTA DE PEDIDOS =====")
			pedidoService.ListarPedidos()

		case "4":
			fmt.Println("\n========== REPORTES ==========")
			fmt.Println("Total de productos:", len(productoService.Productos))
			fmt.Println("Total de clientes:", len(clienteService.Clientes))
			fmt.Println("Total de pedidos:", len(pedidoService.Pedidos))
			fmt.Println("==============================")

		case "5":
			fmt.Println("Gracias por utilizar el sistema E-Commerce")
			return

		default:
			fmt.Println("Opcion no valida")
		}
	}
}
func mostrarTablaProductos(productos []models.Producto) {
	fmt.Println("\n================ LISTA DE PRODUCTOS ================")
	fmt.Printf("%-8s %-20s %-12s %-10s %-15s\n",
		"ID", "NOMBRE", "PRECIO", "STOCK", "CATEGORIA")

	fmt.Println(strings.Repeat("-", 70))

	if len(productos) == 0 {
		fmt.Println("No hay productos registrados.")
		return
	}

	for _, p := range productos {
		fmt.Printf("%-8d %-20s $%-11.2f %-10d %-15s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}

	fmt.Println(strings.Repeat("-", 70))
}
