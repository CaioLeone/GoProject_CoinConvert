package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rates := map[string]float64{
		"USD": 0.151,
		"EUR": 0.137,
		"JPY": 16.29,
		"GBP": 0.13,
		"CHF": 0.1402,
		"AUD": 0.2712,
		"CAD": 0.2374,
		"CNY": 1.251,
		"HKD": 1.326,
		"NZD": 0.2922,
		"SEK": 1.655,
		"NOK": 1.806,
		"DKK": 1.122,
		"SGD": 0.2249,
		"KRW": 242.97,
		"ZAR": 3.239,
		"MXN": 3.454,
		"INR": 14.71,
		"ILS": 0.63,
		"THB": 5.74,
		"IDR": 2875.0,
		"MYR": 0.754,
		"PHP": 9.74,
		"PLN": 0.644,
		"CZK": 3.77,
		"HUF": 61.59,
		"TRY": 6.49,
		"BGN": 0.293,
		"RON": 0.746,
	}

	if len(os.Args) != 3 {
		fmt.Print("./Convert [valor_em_brl] [moeda_destino]: ")
		return
	}
	//CAPTURA STRING
	valor_em_string := os.Args[1]
	moeda_destino := os.Args[2]

	//CONVERTER STRING FLOAT
	valor_em_brl, err := strconv.ParseFloat(valor_em_string, 64)
	if err != nil {
		fmt.Println("Valor em BRL deve ser um numero valido")
		return
	}

	//TRATAMENTO MOEDA. TUDO EM MAIUSCULO
	moeda_destino = strings.ToUpper(moeda_destino)

	//VERIFICACAO DA MOEDA
	rate, exists := rates[moeda_destino]
	if !exists {
		fmt.Println("Moeda nao existe")
		return
	}

	resultado := ConvertMoney(valor_em_brl, rate)
	fmt.Printf("%.2f \n", resultado)
}

func ConvertMoney(real float64, exchange float64) float64 {
	return real * exchange
}
