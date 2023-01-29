package dao

import "github.com/sebastiaofortes/layout_go/internal/domain"

type PaisDao struct{
	
}

func(p *PaisDao)GetPais() domain.Pais{
	return domain.Pais{}
}

func NewPaisDao() PaisDao {
	return PaisDao{}
}