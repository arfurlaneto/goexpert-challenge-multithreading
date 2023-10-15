package cep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetCepFromViaCep(ctx context.Context, cep string) (*ViaCepResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("ViaCep: invalid response from provider")
	}

	var apiResponse ViaCepResponse
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (r *ViaCepResponse) GetProviderName() string {
	return "ViaCEP (https://viacep.com.br/)"
}

func (r *ViaCepResponse) Format() string {
	return fmt.Sprintf("%s - %s - %s, %s - %s", r.Logradouro, r.Bairro, r.Localidade, r.Uf, r.Cep)
}
