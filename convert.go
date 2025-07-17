package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const apiURL = "https://api.freecurrencyapi.com/v1/latest?"

type Coin struct {
	Data map[string]float64 `json:"data"`
}

func main() {
	var base_value float64
	var final_value float64
	var base_currency string
	var final_currency string

	fmt.Println("Enter a value to convert:")
	fmt.Scan(&base_value)

	fmt.Println("Enter a base currency (EUR,USD,JPN...): ")
	fmt.Scan(&base_currency)

	for len(base_currency) != 3 {
		fmt.Println("Insert a valid currency")
		fmt.Println("Enter a base currency (EUR,USD,JPN...): ")
		fmt.Scan(&base_currency)
	}

	fmt.Println("Enter a money to convert (EUR,USD,JPN...): ")
	fmt.Scan(&final_currency)

	for len(final_currency) != 3 {
		fmt.Println("Insert a valid currency")
		fmt.Println("Enter a base currency (EUR,USD,JPN...): ")
		fmt.Scan(&final_currency)
	}

	fmt.Println("Currency choosed:", final_currency)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar .env")
	}

	accessKey := os.Getenv("ACCESS_KEY")

	urlKey := fmt.Sprintf("%sapikey=%s&currencies=%s&base_currency=%s", apiURL, accessKey, final_currency, base_currency)
	resp, err := http.Get(urlKey)

	if err != nil {
		fmt.Println("Error in the request: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	var coin Coin

	erro := json.Unmarshal(body, &coin)
	if erro != nil {
		fmt.Println("erro: ", erro)
		return
	}

	for moeda, valor := range coin.Data {
		fmt.Printf("Conversion: %s to %s | 1 %s is equal to %.2f %s\n", base_currency, moeda, base_currency, valor, moeda)
		final_value = base_value * valor
		fmt.Printf("Converted the value: %.2f %s", final_value, moeda)
	}
}
