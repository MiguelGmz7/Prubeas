package main
import "fmt"

func main() {

	var nombre string

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Printf("Nombre: %s \n",nombre)
}