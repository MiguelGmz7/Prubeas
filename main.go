package main
import "fmt"

func ModString(s string) string {
	rns := []rune(s) 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
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

	fmt.Print("Ingresa input: ")
	fmt.Scanln(&nombre)

	rvs := ModString(nombre)
	fmt.Printf("Desencriptado: %s \n",rvs)
}