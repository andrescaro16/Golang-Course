package main

import "fmt"

func main() {

	valores := []interface{}{1, "Hola", 3.14}

	for _, valor := range valores {
		switch v := valor.(type) {
		case int:
			fmt.Printf("Entero: %d\n", v)
		case string:
			fmt.Printf("Cadena: %s\n", v)
		case float64:
			fmt.Printf("Flotante: %f\n", v)
		default:
			fmt.Println("Tipo desconocido")
		}
	}
	
}