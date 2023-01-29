package dao

import "github.com/sebastiaofortes/layout_go/internal/domain"

type EstadoFisicoDao struct{
	
}


func(p *EstadoFisicoDao)GetEstadoFisico() domain.EstadoFisico{
	return domain.EstadoFisico{}
}

func NewEstadoFisicoDao()EstadoFisicoDao {
	return EstadoFisicoDao{}
}