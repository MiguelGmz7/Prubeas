package main
import "fmt"

func ModString(s string) string {
	rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }

	for l = 0; l <= len(rns)-1; l++ {
		if rns[l] = "3" {
			fmt.Println("hola")
		}
	}

    return string(rns)
}

func main() {

	var nombre string

	fmt.Print("Ingresa isnput: ")
	fmt.Scanln(&nombre)

	rvs := ModString(nombre)

	fmt.Println(len(rvs))
	fmt.Printf("Revertido: %s \n",rvs)
}