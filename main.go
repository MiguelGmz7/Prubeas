package main
import "fmt"

func ModString(s string) string {
	rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }

	for k := 0; k <= len(rns)-1; k++ {
		if (rns[k] >= 48 && rns[k] <= 57) {
			rns[k] = 0
		} 
	}

    return string(rns)
}

func main() {

	var nombre string

	fmt.Print("Ingresa isnput: ")
	fmt.Scanln(&nombre)

	rvs := ModString(nombre)
	fmt.Printf("Revertido: %s \n",rvs)
}