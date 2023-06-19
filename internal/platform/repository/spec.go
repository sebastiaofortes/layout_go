package repository

import "github.com/sebastiaofortes/layout_go/internal/domain"

type PessoaRepository interface {
	GetAllPessoas() ([]domain.Pessoa, error)
	GetPessoasPorPais(int32) ([]domain.Pessoa, error)
	GetPessoasPorIdade(int32) ([]domain.Pessoa, error)
}

type PaisRepository interface {
	GetPais(int32)(domain.Pais, error)
}
