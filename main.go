package main

import (
	"fmt"
	"os"
)

func main() {
	var comando int

	fmt.Println("Salve salve Aurigod")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	// & serve para acessar o endereço de memoria.
	fmt.Scan(&comando)

	switch comando {
	case 1:
		fmt.Println("\nIniciando monitoramento...")
	case 2:
		fmt.Println("\nExibindo logs...")
	case 0:
		fmt.Println("\nSaindo...")
	default:
		fmt.Println("Não reconheço este comando.")
		os.Exit(1)
	}
}
