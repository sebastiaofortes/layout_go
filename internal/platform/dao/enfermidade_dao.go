package dao

import (
	"github.com/sebastiaofortes/layout_go/internal/platform/db"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
)

type EnfermidadeDao struct {
}

func (p *EnfermidadeDao) GetEnfermidades(id int32) []dto.Enfermidade {
	result := []dto.Enfermidade{}
	db.Storege.First(&result, id)
	return result
}

func NewEnfermidadeDao() EnfermidadeDao {
	return EnfermidadeDao{}
}
