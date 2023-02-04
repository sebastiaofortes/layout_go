package domain

type PessoaRepository interface {
	GetAllPessoas() ([]Pessoa, error)
	GetPessoasPorPais(int32) ([]Pessoa, error)
	GetPessoasPorIdade(int32) ([]Pessoa, error)
}

type PaisRepository interface {
	GetPais(int32)(Pais, error)
}
