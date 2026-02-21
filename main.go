package main

import (
	"fmt"
)

func main() {
	var valor_em_brl float64
	var moeda_destino string
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

	fmt.Println("Bem Vindo, Aventureiro! Vamos converter suas moedas!")

	value, ok := rates[moeda_destino]
	if ok {
		fmt.Println("existe essa moeda: ", value)

		ConvertMoney(valor_em_brl, value)
	} else {
		fmt.Println("Moeda nao existe")
	}
}

func ConvertMoney(real float64, exchange float64) float64 {
	return real * exchange
}
