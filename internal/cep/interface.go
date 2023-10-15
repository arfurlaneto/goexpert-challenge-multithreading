package cep

type CepApiResponse interface {
	GetProviderName() string
	Format() string
}
