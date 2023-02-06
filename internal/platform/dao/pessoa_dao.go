package dao

import (
	"github.com/sebastiaofortes/layout_go/internal/platform/db"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
)

type PessoaDao struct {
}

func (p *PessoaDao) GetPessoa(id int32) dto.Pessoa {
	result := dto.Pessoa{}
	db.Storege.First(&result, id)
	return result
}

func (p *PessoaDao) GetPessoasPais(id int32) []dto.Pessoa {
	result := []dto.Pessoa{}
	db.Storege.Where("pais=?", id).Find(&result)
	return result
}

func (p *PessoaDao) GetPessoasIdade(id int32) []dto.Pessoa {
	result := []dto.Pessoa{}
	db.Storege.Where("idade=?", id).Find(&result)
	return result
}

func NewPessoaDao() PessoaDao {
	return PessoaDao{}
}
