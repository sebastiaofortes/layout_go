package repository

import (
	"errors"

	"github.com/sebastiaofortes/layout_go/internal/domain"
	"github.com/sebastiaofortes/layout_go/internal/platform/dao"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
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
	pes, err := p.pessoaDao.GetPessoasPais(pais)
	if err != nil {
		return []domain.Pessoa{}, errors.New("Nenhum resultado encontrado")
	}
	result := []domain.Pessoa{}
	for _, v := range pes {
		r := p.makePessoa(v)
		result = append(result, r)
	}
	return result, nil
}

func (p *ImplementsPessoaRepository) GetPessoasPorIdade(idade int32) ([]domain.Pessoa, error) {
	pes, err := p.pessoaDao.GetPessoasIdade(idade)
	if err != nil {
		return []domain.Pessoa{}, errors.New("Nenhum resultado encontrado")
	}
	result := []domain.Pessoa{}
	for _, v := range pes {
		r := p.makePessoa(v)
		result = append(result, r)
	}
	return result, nil
}

func (p *ImplementsPessoaRepository) makePessoa(pes dto.Pessoa) domain.Pessoa {
	r := pes.ToDomain()
	enfermidades := p.enfermidadeDao.GetEnfermidades(pes.Id)
	r.Enfermidades = p.makeEnfermidades(enfermidades)

	estFis := p.estadoFisicoDao.GetEstadoFisico(pes.Id)
	r.EstadoFisico = estFis.ToDomain()
	return r
}

func (p *ImplementsPessoaRepository) makeEnfermidades(enf []dto.Enfermidade) []domain.Enfermidade {
	enfs := []domain.Enfermidade{}
	for _, v := range enf {
		enfs = append(enfs, v.ToDomain())
	}
	return enfs
}

func NewImplementsPessoaRepository() *ImplementsPessoaRepository {
	return &ImplementsPessoaRepository{}
}
