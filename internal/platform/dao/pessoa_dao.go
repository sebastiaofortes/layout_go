package dao

import "github.com/sebastiaofortes/layout_go/internal/domain"

type PessoaDao struct{
	
}

func(p *PessoaDao)GetPessoa() domain.Pessoa{
	return domain.Pessoa{}
}

func NewPessoaDao() PessoaDao {
	return PessoaDao{}
}