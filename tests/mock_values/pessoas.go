package mock_values

import "github.com/sebastiaofortes/layout_go/internal/domain"

var Pessoas []domain.Pessoa = []domain.Pessoa{
	{
		Nome:         "João",
		Enfermidades: []domain.Enfermidade{},
		Idade:        32,
		EstadoFisico: domain.EstadoFisico{},
		Pais:         1,
	},
}