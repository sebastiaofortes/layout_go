package dao

import (
	"github.com/sebastiaofortes/layout_go/internal/platform/db"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
)

type PaisDao struct {
}

func (p *PaisDao) GetPais(id int32) dto.Pais {
	result := dto.Pais{}
	db.Storege.First(&result, id)
	return result
}

func NewPaisDao() PaisDao {
	return PaisDao{}
}
