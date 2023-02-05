package dao

import (
	"github.com/sebastiaofortes/layout_go/internal/platform/db"
	"github.com/sebastiaofortes/layout_go/internal/platform/dto"
)

type EstadoFisicoDao struct {
}

func (p *EstadoFisicoDao) GetEstadoFisico(id int32) dto.EstadoFisico {
	result := dto.EstadoFisico{}
	db.Storege.First(&result, id)
	return result
}

func NewEstadoFisicoDao() EstadoFisicoDao {
	return EstadoFisicoDao{}
}
