package repository

import (
	"github.com/sebastiaofortes/layout_go/internal/domain"
	"github.com/sebastiaofortes/layout_go/internal/platform/dao"
)

// repository lida com a persistência de agregados
// para persistir esses dados ele pode usar um dao ou um ORM
type ImplementsPessoaRepository struct {
	enfermidadeDao  dao.EnfermidadeDao
	estadoFisicoDao dao.EstadoFisicoDao
	pessoaDao       dao.PessoaDao
}

func (p *ImplementsPessoaRepository) GetAllPessoas() ([]domain.Pessoa, error) {
	return []domain.Pessoa{}, nil
}

func (p *ImplementsPessoaRepository) GetPessoasPorPais(pais int32) ([]domain.Pessoa, error) {
	pes := p.pessoaDao.GetPessoasPais(pais)
	result := []domain.Pessoa{}
	for _, v := range pes {
		enfermidade := p.enfermidadeDao.GetEnfermidades(v.Id)
		r := v.ToDomain()
		enfs := []domain.Enfermidade{}
		for _, v := range enfermidade {
			enfs = append(enfs, v.ToDomain())
		}
		estFis := p.estadoFisicoDao.GetEstadoFisico(v.Id)
		r.Enfermidades = enfs
		r.EstadoFisico = estFis.ToDomain()
		result = append(result, r)
	}
	return result, nil
}

func (p *ImplementsPessoaRepository) GetPessoasPorIdade(idade int32) ([]domain.Pessoa, error) {
	return []domain.Pessoa{}, nil
}

// Recomenda-se criar métodos de build das implemtaçoes das interfaces para verificar se todos os métodos estão de fato sendo implemetados
// nosse função tem como tipo uma interface para que suas impermentações sejam aceitas como objeto de retorno.
func NewImplementsPessoaRepository() domain.PessoaRepository {
	return &ImplementsPessoaRepository{}
}
