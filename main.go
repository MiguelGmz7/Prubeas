package main
import "fmt"

func ModString(s string) string {
	rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    for k, r = range(rns) {
		fmt.Println(k, string(r))
	}
	
    return string(rns)
}

func main() {

	var nombre string

	fmt.Print("Ingresa input: ")
	fmt.Scanln(&nombre)

	rvs := ModString(nombre)

	fmt.Printf("Revertido: %s \n",rvs)
}