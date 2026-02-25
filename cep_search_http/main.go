package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	VIACEP_URL = "https://viacep.com.br/ws/"
)

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
	http.HandleFunc("/", SearchCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// https://viacep.com.br/ws/{cep}/json
	fullUrl, err := url.Parse(VIACEP_URL + cepParam + "/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req, err := http.Get(fullUrl.String())
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data CEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
