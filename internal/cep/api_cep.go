package cep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetCepFromCepApi(ctx context.Context, cep string) (*ApiCepResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://cdn.apicep.com/file/apicep/"+cep+".json", nil)
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
		return nil, errors.New("ApiCep: invalid response from provider")
	}

	var apiResponse ApiCepResponse
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

type ApiCepResponse struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int64  `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (r *ApiCepResponse) GetProviderName() string {
	return "apiCEP (https://apicep.com/)"
}

func (r *ApiCepResponse) Format() string {
	return fmt.Sprintf("%s - %s - %s, %s - %s", r.Address, r.District, r.City, r.State, r.Code)
}
