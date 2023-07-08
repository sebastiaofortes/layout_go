package service

import "github.com/sebastiaofortes/layout_go/internal/domain"

//go:generate mockery --name=PessoaRepository --output=./impl_mocks --outpkg=impl_mocks
type PessoaRepository interface {
	GetAllPessoas() ([]domain.Pessoa, error)
	GetPessoasPorPais(int32) ([]domain.Pessoa, error)
	GetPessoasPorIdade(int32) ([]domain.Pessoa, error)
}

//go:generate mockery --name=PaisRepository --output=./impl_mocks --outpkg=impl_mocks
type PaisRepository interface {
	GetPais(int32) (domain.Pais, error)
}
