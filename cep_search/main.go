package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Website used: https://viacep.com.br/ws/
// Useful site: https://mholt.github.io/json-to-go/
type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// From the first to the end. Need to pass arguments while running the code
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Something went wrong")
		}
		var data CEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
		}
		fmt.Println(data)
	}
}
