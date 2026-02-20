package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var valor_em_brl float64
	//var moeda_destino string

	fmt.Println("Bem Vindo, Aventureiro! Vamos converter suas moedas!")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("./convert [valor_em_brl] [moeda_destino]\n")
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%f", &valor_em_brl)
	fmt.Scanln(&valor_em_brl)

}
