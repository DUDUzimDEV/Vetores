package main 
 
 import "fmt"

 func main () {
	nomes := []string{"Eduardo", "Bruno", "Pedro", "Yan", "Vini"}
	fmt.Println(nomes)
	nomesOne := nomes[:2]
	fmt.Println(nomesOne)
	nomesTwo := nomes[3:5]
	fmt.Println(nomesTwo)
	rangeOne := nomes[2]
	fmt.Println(rangeOne)
 }