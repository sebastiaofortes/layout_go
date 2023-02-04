package repository

import (
	"github.com/sebastiaofortes/layout_go/internal/domain"
	"github.com/sebastiaofortes/layout_go/internal/platform/dao"
)

//repository lida com a persistência de agregados
//para persistir esses dados ele pode usar um dao ou um ORM
type ImplementsPessoaRepository struct{
	enfermidadeDao dao.EnfermidadeDao
	estadoFisicoDao dao.EstadoFisicoDao

}

func (p *ImplementsPessoaRepository ) GetAllPessoas() ([]domain.Pessoa, error) {
	return []domain.Pessoa{}, nil
}

func (p *ImplementsPessoaRepository ) GetPessoasPorPais(pais int32) ([]domain.Pessoa, error) {
	return []domain.Pessoa{}, nil
}

func (p *ImplementsPessoaRepository ) GetPessoasPorIdade(idade int32) ([]domain.Pessoa, error) {
	return []domain.Pessoa{}, nil
}

//Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
//nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewImplementsPessoaRepository() domain.PessoaRepository{
	return &ImplementsPessoaRepository{}
}