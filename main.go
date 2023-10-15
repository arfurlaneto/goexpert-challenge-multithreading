package main

import (
	"context"
	"fmt"
	cepPkg "goexpert-challenge-multithreading/internal/cep"
	"os"
	"regexp"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cep := getCepFromArgs()

	c := make(chan cepPkg.CepApiResponse)

	go func() {
		apiResponse, err := cepPkg.GetCepFromCepApi(ctx, cep)
		if err == nil {
			c <- apiResponse
		}
	}()

	go func() {
		apiResponse, err := cepPkg.GetCepFromViaCep(ctx, cep)
		if err == nil {
			c <- apiResponse
		}
	}()

	select {
	case apiResponse := <-c:
		fmt.Printf("Got CEP data from %s.\n", apiResponse.GetProviderName())
		fmt.Printf("%s\n", apiResponse.Format())
	case <-time.After(time.Second * 1):
		println("Timeout error. Could not get CEP data.")
	}
}

func getCepFromArgs() string {
	cep := "04010-970"

	if len(os.Args) > 1 {
		matched, err := regexp.MatchString("\\d{5}-\\d{3}", os.Args[1])
		if matched && err == nil {
			cep = os.Args[1]
		}
	}

	return cep
}
