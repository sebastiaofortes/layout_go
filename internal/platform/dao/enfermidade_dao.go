package dao

import "github.com/sebastiaofortes/layout_go/internal/domain"

type EnfermidadeDao struct{
	
}

func(p *EnfermidadeDao)GetEnfermidade() domain.Enfermidade{
	return domain.Enfermidade{}
}

func NewEnfermidadeDao()EnfermidadeDao {
	return EnfermidadeDao{}
}